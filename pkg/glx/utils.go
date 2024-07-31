package glx

func sliceHas[T comparable](slice []T, o T) bool {
	for _, i := range slice {
		if i == o {
			return true
		}
	}
	return false
}

func sliceHasAll[T comparable](slice []T, item ...T) bool {
	for _, i := range item {
		if !sliceHas[T](slice, i) {
			return false
		}
	}
	return true
}

func sliceHasExact[T comparable](slice []T, item ...T) bool {
	if len(slice) != len(item) {
		return false
	}
	for _, i := range item {
		if !sliceHas[T](slice, i) {
			return false
		}
	}
	return true
}
