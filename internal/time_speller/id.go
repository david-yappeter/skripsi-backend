package time_speller

import (
	"fmt"
	"strings"
)

var (
	idTimeSpeller Interface = newTimeSpeller("id", integerToIDID)
)

func integerToIDID(input int) (string, error) {
	var units = [4]string{"Hari", "Jam", "Menit", "Detik"}

	if input < 0 {
		return "", ErrNegativeValue
	}

	words := []string{}
	results := integerDivideByDurations(input)

	for idx, val := range results {
		if val == 0 {
			continue
		}

		words = append(words, fmt.Sprintf("%d %s", val, units[idx]))
	}

	return strings.Join(words, " "), nil
}
