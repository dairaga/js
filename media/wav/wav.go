package wav

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// RIFF riff chunk
type RIFF struct {
	ID     string
	Size   uint32
	Format string
}

func (r *RIFF) String() string {
	return fmt.Sprintf(`{"id": %q, "size": %d, "format": %q}`, r.ID, r.Size, r.Format)
}

// Format format chunk
type Format struct {
	ID           string
	Size         uint32
	AudioFormat  uint16
	Channels     uint16
	SampleRate   uint32
	ByteRate     uint32
	BlockAlign   uint16
	BitPerSample uint16
	Extra        []byte
}

func (f *Format) String() string {
	return fmt.Sprintf(`{"id": %q, "size": %d, "audio_format": %d, "channels": %d, "sample_rate": %d, "byte_rate": %d, "block_align": %d, "bit_per_sample": %d, "extra": %v}`,
		f.ID, f.Size, f.AudioFormat, f.Channels, f.SampleRate, f.ByteRate, f.BlockAlign, f.BitPerSample, f.Extra)
}

// Data data chunk
type Data struct {
	ID   string
	Size uint32
}

func (d *Data) String() string {
	return fmt.Sprintf(`{"id": %q, "size": %d}`, d.ID, d.Size)
}

// Info wave information.
type Info struct {
	FileSize uint32
	Duration int64
	RIFF     *RIFF
	Format   *Format
	Data     *Data
	Extra    map[string]uint32
}

func (i *Info) String() string {
	return fmt.Sprintf(`"file_size": %d, "duration": %d, "riff": %v, "format": %v, "data": %v, "extra": %v`, i.FileSize, i.Duration, i.RIFF, i.Format, i.Data, i.Extra)
}

// Read ...
func Read(raw []byte) (*Info, error) {
	info := &Info{}
	info.RIFF = new(RIFF)
	info.Format = new(Format)
	info.Data = new(Data)
	info.Extra = make(map[string]uint32)

	reader := bytes.NewReader(raw)

	id := make([]byte, 4)
	var size uint32

	// RIFF
	if err := binary.Read(reader, binary.BigEndian, id); err != nil {
		return nil, err
	}
	info.RIFF.ID = string(id)

	// size
	if err := binary.Read(reader, binary.LittleEndian, &info.RIFF.Size); err != nil {
		return nil, err
	}
	info.FileSize = info.RIFF.Size + 8

	// WAVE
	if err := binary.Read(reader, binary.BigEndian, id); err != nil {
		return nil, err
	}
	info.RIFF.Format = string(id)

	if err := binary.Read(reader, binary.BigEndian, id); err != nil {
		return nil, err
	}

	// fmt
	for !(id[0] == 'f' && id[1] == 'm' && id[2] == 't') {
		if err := binary.Read(reader, binary.LittleEndian, &size); err != nil {
			return nil, err
		}
		info.Extra[string(id)] = size

		// skip
		if _, err := reader.Seek(int64(size), io.SeekCurrent); err != nil {
			return nil, err
		}
		if err := binary.Read(reader, binary.BigEndian, id); err != nil {
			return nil, err
		}
	}

	if err := binary.Read(reader, binary.LittleEndian, &size); err != nil {
		return nil, err
	}
	info.Format.ID = string(id)
	info.Format.Size = size

	if err := binary.Read(reader, binary.LittleEndian, &info.Format.AudioFormat); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &info.Format.Channels); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &info.Format.SampleRate); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &info.Format.ByteRate); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &info.Format.BlockAlign); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &info.Format.BitPerSample); err != nil {
		return nil, err
	}

	if size-16 > 0 {
		info.Format.Extra = make([]byte, size-16)
		if err := binary.Read(reader, binary.BigEndian, info.Format.Extra); err != nil {
			return nil, err
		}
	}

	// fllr
	if err := binary.Read(reader, binary.BigEndian, id); err != nil {
		return nil, err
	}

	for !(id[0] == 'd' && id[1] == 'a' && id[2] == 't' && id[3] == 'a') {
		if err := binary.Read(reader, binary.LittleEndian, &size); err != nil {
			return nil, err
		}
		info.Extra[string(id)] = size
		if _, err := reader.Seek(int64(size), io.SeekCurrent); err != nil {
			return nil, err
		}

		binary.Read(reader, binary.BigEndian, id)
	}

	// data
	info.Data.ID = string(id)
	if err := binary.Read(reader, binary.LittleEndian, &info.Data.Size); err != nil {
		return nil, err
	}

	total := float64(info.Data.Size)
	d := total / float64((info.Format.SampleRate * uint32((info.Format.BitPerSample / 8)) * uint32(info.Format.Channels)))
	info.Duration = int64(d * 1000)
	return info, nil
}
