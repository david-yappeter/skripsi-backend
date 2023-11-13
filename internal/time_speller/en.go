package time_speller

import (
	"fmt"
	"strings"
)

var (
	enTimeSpeller Interface = newTimeSpeller("en", integerToEnUs)
)

func integerToEnUs(input int) (string, error) {
	var units = [4]string{"Day", "Hour", "Minute", "Second"}

	if input < 0 {
		return "", ErrNegativeValue
	}

	words := []string{}
	results := integerDivideByDurations(input)

	for idx, val := range results {
		if val == 0 {
			continue
		}

		// plural form for english
		unit := units[idx]
		if val > 1 {
			unit += "s"
		}
		words = append(words, fmt.Sprintf("%d %s", val, unit))
	}

	return strings.Join(words, " "), nil
}
