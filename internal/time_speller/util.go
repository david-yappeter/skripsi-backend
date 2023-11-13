package time_speller

var durations = [4]int{86400, 3600, 60, 1}

func integerDivideByDurations(input int) []int {
	results := []int{}
	idx := 0

	for {
		if input == 0 {
			break
		}

		results = append(results, input/durations[idx])
		input %= durations[idx]

		idx++
	}

	return results
}
