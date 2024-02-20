package util

func RemoveDuplicate(s []string) []string {
	result := []string{}
	seen := map[string]bool{}
	for _, val := range s {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = true
		}
	}
	return result
}

// function to check if everything in array B exist in array A
func ContainsAll(a []string, b []string) bool {
	for _, v := range b {
		if !StringInSlice(v, a) {
			return false
		}
	}
	return true
}

func SliceValueToSlicePointer[T any](sliceValue []T) []*T {
	slicePointer := []*T{}
	for i := range sliceValue {
		slicePointer = append(slicePointer, &sliceValue[i])
	}

	return slicePointer
}

func SlicePointerToSliceValue[T any](sliceValue []*T) []T {
	slicePointer := []T{}
	for i := range sliceValue {
		slicePointer = append(slicePointer, *sliceValue[i])
	}

	return slicePointer
}
