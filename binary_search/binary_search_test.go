package binarysearch

import (
	"cmp"
	"strings"
	"testing"
)

func TestBinarySearchInts(t *testing.T) {
	arr := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	cases := []struct {
		name    string
		toFind  int
		want    bool
		wantIdx int
	}{
		{
			name:    "69 exists",
			toFind:  69,
			want:    true,
			wantIdx: 3,
		}, {
			name:    "1336 doesn't exists",
			toFind:  1336,
			want:    false,
			wantIdx: 0,
		}, {
			name:    "69420 exists",
			toFind:  69420,
			want:    true,
			wantIdx: 10,
		}, {
			name:    "69421 doesn't exists",
			toFind:  69421,
			want:    false,
			wantIdx: 0,
		}, {
			name:    "1 exists",
			toFind:  1,
			want:    true,
			wantIdx: 0,
		}, {
			name:    "0 doesn't exists",
			toFind:  0,
			want:    false,
			wantIdx: 0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			gotIdx, got := BinarySearch(tt.toFind, arr, cmp.Compare[int])
			if got != tt.want {
				t.Errorf("got %t; want %t", got, tt.want)
			}
			if gotIdx != tt.wantIdx {
				t.Errorf("got index %d; want index %d", gotIdx, tt.wantIdx)
			}
		})
	}
}

func TestBinarySearchStrings(t *testing.T) {
	arr := []string{"1", "1337", "3", "4", "420", "69", "69420", "71", "81", "90", "99"}
	cases := []struct {
		name    string
		toFind  string
		want    bool
		wantIdx int
	}{
		{
			name:    "69 exists",
			toFind:  "69",
			want:    true,
			wantIdx: 5,
		}, {
			name:    "1336 doesn't exists",
			toFind:  "1336",
			want:    false,
			wantIdx: 0,
		}, {
			name:    "69420 exists",
			toFind:  "69420",
			want:    true,
			wantIdx: 6,
		}, {
			name:    "69421 doesn't exists",
			toFind:  "69421",
			want:    false,
			wantIdx: 0,
		}, {
			name:    "1 exists",
			toFind:  "1",
			want:    true,
			wantIdx: 0,
		}, {
			name:    "empty doesn't exists",
			toFind:  "",
			want:    false,
			wantIdx: 0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			gotIdx, got := BinarySearch(tt.toFind, arr, strings.Compare)
			if got != tt.want {
				t.Errorf("got %t; want %t", got, tt.want)
			}
			if gotIdx != tt.wantIdx {
				t.Errorf("got index %d; want index %d", gotIdx, tt.wantIdx)
			}
		})
	}
}

func TestBinarySearchEmptyArr(t *testing.T) {
	arr := []int{}
	vals := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420, 0}

	for _, val := range vals {
		if _, exists := BinarySearch(val, arr, cmp.Compare[int]); exists {
			t.Errorf("expected arr to be empty but found %d", val)
		}
	}
}
