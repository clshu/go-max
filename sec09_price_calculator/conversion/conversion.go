package conversion

import "strconv"

// StringsToFloats converts a slice of strings to a slice of float64.
func StringsToFloats(strings []string) ([]float64, error) {
	var numbers = make([]float64, len(strings))
	for index, str := range strings {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		numbers[index] = num
	}
	return numbers, nil
}
