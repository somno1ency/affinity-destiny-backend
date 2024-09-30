package util

func Contains[T comparable](slice []T, ele T) bool {
	for _, v := range slice {
		if v == ele {
			return true
		}
	}

	return false
}
