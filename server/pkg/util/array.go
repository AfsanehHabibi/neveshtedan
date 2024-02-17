package util

func RemoveNilElements[T any](arr []*T) (result []T) {
	for _, v := range arr {
		if v != nil {
			result = append(result, *v)
		}
	}
	return
}
