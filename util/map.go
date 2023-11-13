package util

import "sort"

func GetSortedFloatMapWithDateStringKeys(m map[string]float64) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}

	// sort ascending
	sort.Slice(keys, func(i, j int) bool {
		iTime := ParseDate(keys[i])
		jTime := ParseDate(keys[j])

		return iTime.IsLessThan(jTime)
	})

	return keys
}
