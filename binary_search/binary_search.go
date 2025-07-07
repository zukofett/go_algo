package binarysearch

func BinarySearch[S ~[]T, T comparable](needle T, haystack S, comp func(T, T) int) (int, bool) {
	start := 0
	end := len(haystack) - 1

	for start <= end {
		mid := (end-start)/2 + start
		res := comp(haystack[mid], needle)
		if res > 0 {
			end = mid - 1
		} else if res < 0 {
			start = mid + 1
		} else {
			return mid, true
		}
	}
	return 0, false
}
