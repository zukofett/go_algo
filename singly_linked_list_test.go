package main

import (
	"testing"
)

func TestNewSLL(t *testing.T) {
	sll := NewSLL[int]()

	if sll == nil {
		t.Fatal("failed to initialize list")
	}

	if sll.length != 0 {
		t.Fatal("list should be zero length")
	}

	if sll.tail != sll.head {
		t.Fatal("begin should equal end on an empty list")
	}
}

func TestLen(t *testing.T) {
	cases := []struct {
		name string
		init []int
		want int
	}{
		{
			name: "empty list",
			init: []int{},
			want: 0,
		}, {
			name: "one element",
			init: []int{1},
			want: 1,
		}, {
			name: "few elements",
			init: []int{1, 2, 3, 4, 5},
			want: 5,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			sll := createListFromSlice(tt.init)

			if got := sll.Len(); got != tt.want {
				t.Errorf("got: %d; want: %d", got, tt.want)
			}
		})
	}
}

func createListFromSlice[T any](s []T) *SinglyLinkedList[T] {
	sll := NewSLL[T]()

	for _, el := range s {
		sll.Insert(sll.End(), &el)
	}
	return sll
}
