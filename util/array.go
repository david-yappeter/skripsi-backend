package util

type ArrayStringDiagramVenn struct {
	Intersection []string
	NotExistInA  []string
	ExistInA     []string
	ExistInB     []string
}

func NewArrayStringDiagramVenn(arr1 []string, arr2 []string) ArrayStringDiagramVenn {
	diagram := ArrayStringDiagramVenn{
		Intersection: []string{},
		NotExistInA:  []string{},
		ExistInA:     []string{},
		ExistInB:     []string{},
	}
	uniqueMap := map[string]int{}

	for _, v := range arr1 {
		uniqueMap[v] = 1
		diagram.ExistInA = append(diagram.ExistInA, v)
	}

	for _, v := range arr2 {
		if uniqueMap[v] == 1 {
			diagram.Intersection = append(diagram.Intersection, v)
		} else {
			diagram.NotExistInA = append(diagram.NotExistInA, v)
		}
		diagram.ExistInB = append(diagram.ExistInB, v)
	}

	return diagram
}

func RemoveStringFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func ConvertArray[K any, T any](arr []K, callback func(K) T) []T {
	nodes := []T{}

	for _, v := range arr {
		nodes = append(nodes, callback(v))
	}

	return nodes
}

func AppendIfNotNil[T any](arr []T, v *T) []T {
	if v != nil {
		return append(arr, *v)
	}

	return arr
}

func ExistInArray[T comparable](arr []T, v T) bool {
	isExist := false

	for i := range arr {
		if arr[i] == v {
			isExist = true
			break
		}
	}

	return isExist
}
