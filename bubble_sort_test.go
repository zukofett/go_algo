package main

import (
	"cmp"
	"slices"
	"strings"
	"testing"
)

func TestBubbleSortInts(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
	}{
		{
			name: "standard arr",
			arr:  []int{1, 3, 7, 4, 2},
		}, {
			name: "sorted arr",
			arr:  []int{1, 2, 3, 4, 7},
		}, {
			name: "reversed arr",
			arr:  []int{7, 4, 3, 2, 1},
		}, {
			name: "only 1 element to sort arr",
			arr:  []int{0, 0, 0, 1, 0},
		}, {
			name: "empty arr",
			arr:  []int{},
		}, {
			name: "single element arr",
			arr:  []int{1},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.arr, cmp.Compare)
			if !slices.IsSorted(tt.arr) {
				t.Errorf("expected arr to be sorted but got: %+v", tt.arr)
			}
		})
	}
}

func TestBubbleSortStrings(t *testing.T) {
	cases := []struct {
		name string
		arr  []string
	}{
		{
			name: "standard arr",
			arr:  []string{"1", "3", "7", "4", "2"},
		}, {
			name: "sorted arr",
			arr:  []string{"1", "2", "3", "4", "7"},
		}, {
			name: "reversed arr",
			arr:  []string{"7", "4", "3", "2", "1"},
		}, {
			name: "only 1 element to sort arr",
			arr:  []string{"0", "0", "0", "1", "0"},
		}, {
			name: "empty arr",
			arr:  []string{},
		}, {
			name: "single element arr",
			arr:  []string{"1"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.arr, strings.Compare)
			if !slices.IsSorted(tt.arr) {
				t.Errorf("expected arr to be sorted but got: %+v", tt.arr)
			}
		})
	}
}
