package singlylinkedlist

import (
	"cmp"
	"slices"
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

func TestLenEdge(t *testing.T) {
	var l *SinglyLinkedList[int]

	got := l.Len()
	if got != 0 {
		t.Errorf("expected 0 but got %d", got)
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		atFunc  func(*SinglyLinkedList[int]) *Node[int]
		val     int
		want    []int
		wantLen int
	}{
		{
			name:    "insert at end of empty list",
			initial: []int{},
			atFunc: func(l *SinglyLinkedList[int]) *Node[int] {
				return l.End()
			},
			val:     1,
			want:    []int{1},
			wantLen: 1,
		},
		{
			name:    "insert at end of non-empty list",
			initial: []int{1, 2},
			atFunc: func(l *SinglyLinkedList[int]) *Node[int] {
				return l.End()
			},
			val:     3,
			want:    []int{1, 2, 3},
			wantLen: 3,
		},
		{
			name:    "insert at beginning",
			initial: []int{2, 3},
			atFunc: func(l *SinglyLinkedList[int]) *Node[int] {
				return l.Begin()
			},
			val:     1,
			want:    []int{1, 2, 3},
			wantLen: 3,
		},
		{
			name:    "insert in middle",
			initial: []int{1, 3},
			atFunc: func(l *SinglyLinkedList[int]) *Node[int] {
				target := 3
				return l.Find(l.Begin(), l.End(), &target, compInts)
			},
			val:     2,
			want:    []int{1, 2, 3},
			wantLen: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := createListFromSlice(tt.initial)

			result := list.Insert(tt.atFunc(list), &tt.val)
			if result == nil {
				t.Error("want node; got nil")
			}

			got := list.ToSlice()
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}

			if list.Len() != tt.wantLen {
				t.Errorf("got len of %d; want %d", list.Len(), tt.wantLen)
			}
		})
	}
}

func TestInsertEdge(t *testing.T) {
	tests := []struct {
		name string
		list *SinglyLinkedList[int]
		at   *Node[int]
		val  int
		want *Node[int]
	}{
		{
			name: "nil list",
			list: nil,
			at:   &Node[int]{},
			val:  10,
		}, {
			name: "nil at",
			list: NewSLL[int](),
			at:   nil,
			val:  10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.Insert(tt.at, &tt.val)
			if got != nil {
				t.Errorf("got list = %v, want nil", got)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		atFunc  func(*SinglyLinkedList[int]) *Node[int]
		val     int
		want    []int
		wantLen int
	}{
		{
			name:    "remove at beginning",
			initial: []int{1, 2, 3},
			atFunc: func(l *SinglyLinkedList[int]) *Node[int] {
				return l.Begin()
			},
			want:    []int{2, 3},
			wantLen: 2,
		},
		{
			name:    "remove at middle",
			initial: []int{1, 2, 3},
			atFunc: func(l *SinglyLinkedList[int]) *Node[int] {
				target := 2
				return l.Find(l.Begin(), l.End(), &target, compInts)
			},
			want:    []int{1, 3},
			wantLen: 2,
		},
		{
			name:    "remove from single item list",
			initial: []int{2},
			atFunc: func(l *SinglyLinkedList[int]) *Node[int] {
				return l.Begin()
			},
			want:    []int{},
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := createListFromSlice(tt.initial)

			result := list.Remove(tt.atFunc(list))
			if result == nil {
				t.Error("want node; got nil")
			}

			got := list.ToSlice()
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}

			if list.Len() != tt.wantLen {
				t.Errorf("got len of %d; want %d", list.Len(), tt.wantLen)
			}
		})
	}
}

func TestRemoveEdge(t *testing.T) {
	t.Run("nil list", func(t *testing.T) {
		var list *SinglyLinkedList[int]
		at := &Node[int]{}
		got := list.Remove(at)
		if got != nil {
			t.Errorf("got node = %v, want nil", got)
		}
	})
	t.Run("nil at", func(t *testing.T) {
		list := NewSLL[int]()
		var at *Node[int]
		got := list.Remove(at)
		if got != nil {
			t.Errorf("got node = %v, want nil", got)
		}
	})
	t.Run("remove at tail", func(t *testing.T) {
		list := NewSLL[int]()
		at := list.End()
		got := list.Remove(at)
		if got != nil {
			t.Errorf("got node = %v, want nil", got)
		}
	})
}

func TestFind(t *testing.T) {
	cases := []struct {
		name       string
		initial    []int
		toFind     int
		shouldFind bool
	}{
		{
			name:       "find at beginning",
			initial:    []int{1, 2, 3, 4, 5},
			toFind:     1,
			shouldFind: true,
		}, {
			name:       "find at middle",
			initial:    []int{1, 2, 3, 4, 5},
			toFind:     3,
			shouldFind: true,
		}, {
			name:       "find at end",
			initial:    []int{1, 2, 3, 4, 5},
			toFind:     5,
			shouldFind: true,
		}, {
			name:       "find non existing",
			initial:    []int{1, 2, 3, 4, 5},
			toFind:     6,
			shouldFind: false,
		}, {
			name:       "find in empty list",
			initial:    []int{},
			toFind:     1,
			shouldFind: false,
		}, {
			name:       "find in single element list",
			initial:    []int{101},
			toFind:     101,
			shouldFind: true,
		}, {
			name:       "find non existent in single element list",
			initial:    []int{101},
			toFind:     100,
			shouldFind: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			l := createListFromSlice(tt.initial)

			result := l.Find(l.Begin(), l.End(), &tt.toFind, compInts)
			if tt.shouldFind {
				if result == nil {
					t.Fatal("expected to find node but didn't")
				}
				if result.Data == nil || tt.toFind != *result.Data {
					t.Errorf("expected to find %d but found %d", tt.toFind, *result.Data)

				}
			} else {
				if result == nil {
					t.Fatal("expected to find node but didn't")
				}
				if result != l.End() {
					t.Fatal("expected to return end but didn't")

				}
			}
		})
	}
}

func TestForEach(t *testing.T) {
	cases := []struct {
		name     string
		initial  []int
		do       func(*int) bool
		want     []int
		wantStop int
	}{
		{
			name:    "iterate over all",
			initial: []int{1, 2, 3, 4, 5},
			do: func(i *int) bool {
				*i = *i * *i
				return true
			},
			want:     []int{1, 4, 9, 16, 25},
			wantStop: 0,
		}, {
			name:    "stop immediately",
			initial: []int{1, 2, 3, 4, 5},
			do: func(i *int) bool {
				return false
			},
			want:     []int{1, 2, 3, 4, 5},
			wantStop: 1,
		}, {
			name:    "stop at middle",
			initial: []int{1, 2, 3, 4, 5},
			do: func(i *int) bool {
				if *i == 3 {
					return false
				}
				*i = *i * *i
				return true
			},
			want:     []int{1, 4, 3, 4, 5},
			wantStop: 3,
		}, {
			name:    "empty list",
			initial: []int{},
			do: func(i *int) bool {
				*i = *i * *i
				return true
			},
			want:     []int{},
			wantStop: 0,
		}, {
			name:    "single item",
			initial: []int{9},
			do: func(i *int) bool {
				*i = *i * *i
				return true
			},
			want:     []int{81},
			wantStop: 0,
		}, {
			name:    "single item stop immediately",
			initial: []int{9},
			do: func(i *int) bool {
				return false
			},
			want:     []int{9},
			wantStop: 9,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			l := createListFromSlice(tt.initial)
			res := l.ForEach(l.Begin(), l.End(), tt.do)

			if !slices.Equal(l.ToSlice(), tt.want) {
				t.Errorf("expected list to to equal %v but got %v", tt.want, l.ToSlice())
			}

			if res == nil {
				t.Fatal("got nil result")
			}

			stopNode := l.Find(l.Begin(), l.End(), &tt.wantStop, compInts)
			if stopNode != res {
				t.Errorf("expected node to equal %v but got %v", stopNode, res)
			}
		})
	}
}

/******************************************************************************
                            Helpers
******************************************************************************/

func createListFromSlice[T any](s []T) *SinglyLinkedList[T] {
	sll := NewSLL[T]()

	for _, el := range s {
		sll.Insert(sll.End(), &el)
	}
	return sll
}

func compInts(a, b *int) int {
	return cmp.Compare(*a, *b)
}
