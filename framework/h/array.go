package h

func Unique[T any](slice []T, key func(item T) string) []T {
	var result []T
	seen := make(map[string]bool)
	for _, v := range slice {
		k := key(v)
		if _, ok := seen[k]; !ok {
			seen[k] = true
			result = append(result, v)
		}
	}
	return result
}

func Filter[T any](slice []T, predicate func(item T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T, U any](slice []T, mapper func(item T) U) []U {
	var result []U
	for _, v := range slice {
		result = append(result, mapper(v))
	}
	return result
}
