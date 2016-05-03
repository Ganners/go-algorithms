package atoi

import (
	"errors"
	"math"
)

var (
	StringNotValid = errors.New("String is not a valid float")
)

// Parses a float base 10
func parseFloat(s string) (float64, error) {

	negative := s[0] == '-'
	if negative {
		s = s[1:]
	}

	total := float64(0)

	decimalFound := false
	j := 0

	for i := len(s) - 1; i >= 0; i-- {

		if !decimalFound && s[i] == '.' {
			decimalFound = true
			total *= math.Pow(10, -float64(j))
			j = 0
			continue
		}

		digit := float64(s[i] - '0')

		if digit > 9 {
			return 0, StringNotValid
		}

		total += digit * math.Pow(10, float64(j))
		j++
	}

	if negative {
		total = -total
	}

	return total, nil
}
