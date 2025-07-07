package bubblesort

func BubbleSort[S ~[]T, T comparable](s S, comp func(T, T) int) {
	for i := range s {
		changed := false
		for j := 0; j < len(s)-1-i; j++ {
			if res := comp(s[j], s[j+1]); res > 0 {
				s[j+1], s[j] = s[j], s[j+1]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}
