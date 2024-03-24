package util

func RemoveNilElements[T any](arr []*T) (result []T) {
	for _, v := range arr {
		if v != nil {
			result = append(result, *v)
		}
	}
	return
}

func ConvertToPointerArray[T any](arr []T) (result []*T) {
	for _, v := range arr {
        temp := v  // Create a new variable to capture the value of v
        result = append(result, &temp)
    }
	return
}
