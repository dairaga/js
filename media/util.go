package media

// convert float32 to int16 for 16Bit PCM.
func float32To16BitPCM(s float32) int16 {
	if s < -1.0 {
		s = -1.0
	}

	if s > 1.0 {
		s = 1.0
	}

	if s < 0 {
		s = s * 32768
	}

	if s > 0 {
		s = s * 32767
	}

	return int16(s)
}

// Flaot32To16BitPCM converts float32 data to 16Bit PCM.
func Flaot32To16BitPCM(raw []float32) []int16 {
	result := make([]int16, len(raw))

	for i, x := range raw {
		result[i] = float32To16BitPCM(x)
	}

	return result
}
