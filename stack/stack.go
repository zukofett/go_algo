package stack

type Stack[T any] struct {
	stack []T
}

func NewStack[T any](base int) *Stack[T] {
	return &Stack[T]{
		stack: make([]T, 0, base),
	}
}

func (s *Stack[T]) Push(val T) {
	if s == nil {
		return
	}
	s.stack = append(s.stack, val)
}

func (s *Stack[T]) Pop() T {
	if s == nil {
		var noop T
		return noop
	}

	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return val
}

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

func (s *Stack[T]) ToSlice() []T {
	ret := make([]T, 0, s.Len())
	for _, el := range s.stack {
		ret = append(ret, el)
	}
	return ret
}
