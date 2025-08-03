package stack

import "testing"

func TestNewStack(t *testing.T) {
	baseCap := 10
	stk := NewStack[int](baseCap)

	if stk == nil {
		t.Fatal("failed to initialize stack")
	}

	if len(stk.stack) != 0 {
		t.Fatal("stack should be zero length")
	}

	if cap(stk.stack) != baseCap {
		t.Fatalf("stack initial capacity should be %d but got %d", baseCap, cap(stk.stack))
	}
}

func TestNewStackEdge(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("expected to panic but didn't")
		}
	}()

	baseCap := -10
	stk := NewStack[int](baseCap)

	if stk != nil {
		t.Fatal("failed to initialize stack")
	}
}

func TestStackPush(t *testing.T) {
	cases := []struct {
		name       string
		init       []int
		insertVals []int
		wantVals   []int
	}{
		{
			name:       "insert single element to empty list",
			init:       []int{},
			insertVals: []int{1},
			wantVals:   []int{1},
		}, {
			name:       "insert single element to non-empty list",
			init:       []int{1, 2, 3},
			insertVals: []int{4},
			wantVals:   []int{1, 2, 3, 4},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			stk := fromSlice(t, tt.init)

			if len(stk.stack) != len(tt.init) {
				t.Fatalf("got %d; want %d", len(stk.stack), len(tt.init))
			}

			for _, val := range tt.insertVals {
				stk.Push(val)
			}

			if len(stk.stack) != len(tt.wantVals) {
				t.Fatalf("got %d; want %d", len(stk.stack), len(tt.wantVals))
			}

			for i, el := range stk.stack {
				if el != tt.wantVals[i] {
					t.Fatalf("got element %d at index %d; want %d", el, i, tt.wantVals[i])
				}
			}
		})

	}
}

func TestStackPushEdge(t *testing.T) {
	var stk *Stack[int]

	defer func() {
		if err := recover(); err != nil {
			t.Error("expected not to panic but did")
		}
	}()

	stk.Push(1)
}

func TestStackPop(t *testing.T) {
	cases := []struct {
		name        string
		init        []int
		amount      int
		wantVals    []int
		shouldPanic bool
	}{
		{
			name:        "pop single element from empty stack",
			init:        []int{},
			amount:      1,
			wantVals:    []int{},
			shouldPanic: true,
		}, {
			name:     "pop single element from non-empty stack",
			init:     []int{1, 2, 3},
			amount:   1,
			wantVals: []int{1, 2},
		}, {
			name:        "pop multi element from empty stack",
			init:        []int{},
			amount:      4,
			wantVals:    []int{},
			shouldPanic: true,
		}, {
			name:        "pop multi element from non-empty stack",
			init:        []int{1, 2, 3},
			amount:      4,
			wantVals:    []int{},
			shouldPanic: true,
		}, {
			name:     "pop multi element from non-empty stack with leftovers",
			init:     []int{1, 2, 3, 4, 5},
			amount:   4,
			wantVals: []int{1},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if err := recover(); err == nil {
						t.Error("expected to panic but didn't")
					}
				}()
			}

			stk := fromSlice(t, tt.init)

			if len(stk.stack) != len(tt.init) {
				t.Fatalf("got %d; want %d", len(stk.stack), len(tt.init))
			}

			for range tt.amount {
				stk.Pop()
			}

			if len(stk.stack) != len(tt.wantVals) {
				t.Fatalf("got %d; want %d", len(stk.stack), len(tt.wantVals))
			}

			for i, el := range stk.stack {
				if el != tt.wantVals[i] {
					t.Fatalf("got element %d at index %d; want %d", el, i, tt.wantVals[i])
				}
			}
		})

	}
}

func TestStackPopEdge(t *testing.T) {
	var stk *Stack[int]

	defer func() {
		if err := recover(); err != nil {
			t.Error("expected not to panic but did")
		}
	}()

	stk.Pop()
}

func TestStackIsEmpty(t *testing.T) {
	cases := []struct {
		name   string
		init   []int
		insert []int
		remove int
		want   bool
	}{
		{
			name:   "empty stack",
			init:   []int{},
			insert: []int{},
			remove: 0,
			want:   true,
		}, {
			name:   "non-empty stack",
			init:   []int{1, 2, 3},
			insert: []int{},
			remove: 0,
			want:   false,
		}, {
			name:   "non-empty stack after inserts",
			init:   []int{},
			insert: []int{1, 2, 3},
			remove: 0,
			want:   false,
		}, {
			name:   "empty stack after inserts",
			init:   []int{},
			insert: []int{1, 2, 3},
			remove: 3,
			want:   true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			stk := fromSlice(t, tt.init)

			if len(stk.stack) != len(tt.init) {
				t.Fatalf("got %d; want %d", len(stk.stack), len(tt.init))
			}

			for _, el := range tt.insert {
				stk.Push(el)
			}

			for range tt.remove {
				stk.Pop()
			}

			if stk.IsEmpty() != tt.want {
				t.Fatalf("got %t; want %t", stk.IsEmpty(), tt.want)
			}
		})

	}
}

func TestStackIsEmptyEdge(t *testing.T) {
	var stk *Stack[int]

	if stk.IsEmpty() != true {
		t.Fatalf("got %t; want true", stk.IsEmpty())
	}
}

/*
	func (s *Stack[T]) IsEmpty() bool {
		return s.Len() == 0
	}

	func (s *Stack[T]) Peek() T {
		if s == nil {
			var noop T
			return noop
		}
		return s.stack[len(s.stack)-1]
	}

	func (s *Stack[T]) Len() int {
		if s == nil {
			return 0
		}
		return len(s.stack)
	}

	func (s *Stack[T]) Cap() int {
		if s == nil {
			return 0
		}
		return cap(s.stack)
	}
*/
func fromSlice(t *testing.T, s []int) *Stack[int] {
	t.Helper()

	stk := NewStack[int](len(s))

	for _, i := range s {
		stk.Push(i)
	}
	return stk
}
