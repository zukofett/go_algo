package main

func LinearSearch[S ~[]T, T comparable](needle T, haystack S) (int, bool) {
	for i, ele := range haystack {
		if ele == needle {
			return i, true
		}
	}
	return 0, false
}
