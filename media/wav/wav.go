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
	if err := binary.Read(reader, binary.LittleEndian, id); err != nil {
		return nil, err
	}
	info.RIFF.ID = string(id)

	// size
	if err := binary.Read(reader, binary.LittleEndian, &info.RIFF.Size); err != nil {
		return nil, err
	}
	info.FileSize = info.RIFF.Size + 8

	// WAVE
	if err := binary.Read(reader, binary.LittleEndian, id); err != nil {
		return nil, err
	}
	info.RIFF.Format = string(id)

	if err := binary.Read(reader, binary.LittleEndian, id); err != nil {
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
		if err := binary.Read(reader, binary.LittleEndian, id); err != nil {
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
		if err := binary.Read(reader, binary.LittleEndian, info.Format.Extra); err != nil {
			return nil, err
		}
	}

	// fllr
	if err := binary.Read(reader, binary.LittleEndian, id); err != nil {
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

		binary.Read(reader, binary.LittleEndian, id)
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

// Write ...
func Write(sampleRate uint32, bitPerSample, numOfChannels uint16, raw []byte) ([]byte, error) {
	size := uint32(len(raw))
	if size <= 0 {
		return nil, fmt.Errorf("raw data is empty")
	}
	writer := bytes.NewBuffer(make([]byte, 0, size+36+8))

	// RIFF
	if err := binary.Write(writer, binary.LittleEndian, []byte("RIFF")); err != nil {
		return nil, err
	}

	// Size
	if err := binary.Write(writer, binary.LittleEndian, size+36); err != nil {
		return nil, err
	}

	// WAVE
	if err := binary.Write(writer, binary.LittleEndian, []byte{'W', 'A', 'V', 'E'}); err != nil {
		return nil, err
	}

	// fmt
	if err := binary.Write(writer, binary.LittleEndian, []byte{'f', 'm', 't', ' '}); err != nil {
		return nil, err
	}

	// fmt size
	if err := binary.Write(writer, binary.LittleEndian, uint32(16)); err != nil {
		return nil, err
	}

	//  fmt format (PCM)
	if err := binary.Write(writer, binary.LittleEndian, uint16(1)); err != nil {
		return nil, err
	}

	// number of channels
	if err := binary.Write(writer, binary.LittleEndian, numOfChannels); err != nil {
		return nil, err
	}

	// sample rate
	if err := binary.Write(writer, binary.LittleEndian, sampleRate); err != nil {
		return nil, err
	}

	// ByteRate
	byteRate := sampleRate * uint32(bitPerSample) / 8
	if err := binary.Write(writer, binary.LittleEndian, byteRate); err != nil {
		return nil, err
	}

	// block alignment
	blockAlign := numOfChannels * bitPerSample / 8
	if err := binary.Write(writer, binary.LittleEndian, blockAlign); err != nil {
		return nil, err
	}

	if err := binary.Write(writer, binary.LittleEndian, bitPerSample); err != nil {
		return nil, err
	}

	if err := binary.Write(writer, binary.LittleEndian, []byte{'d', 'a', 't', 'a'}); err != nil {
		return nil, err
	}
	if err := binary.Write(writer, binary.LittleEndian, size); err != nil {
		return nil, err
	}

	if _, err := writer.Write(raw); err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}
