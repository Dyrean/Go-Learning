package utils

import (
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, value := range strings {
		float, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return nil, err
		}

		floats = append(floats, float)
	}

	return floats, nil
}
