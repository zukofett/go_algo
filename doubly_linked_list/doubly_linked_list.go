package doublylinkedlist

type DLLNode[T any] struct {
	Data       *T
	next, prev *DLLNode[T]
}

func (n *DLLNode[T]) Next() *DLLNode[T] {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *DLLNode[T]) Prev() *DLLNode[T] {
	if n == nil {
		return nil
	}
	return n.prev
}

type DoublyLinkedList[T any] struct {
	head, tail DLLNode[T]
	length     int
}

func NewDLL[T any]() *DoublyLinkedList[T] {
	newDLL := DoublyLinkedList[T]{}
	newDLL.head.next = &newDLL.tail
	newDLL.tail.prev = &newDLL.head

	return &newDLL
}

func (l *DoublyLinkedList[T]) Len() int {
	if l == nil {
		return 0
	}
	return l.length
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l == nil || l.length == 0
}

func (l *DoublyLinkedList[T]) Begin() *DLLNode[T] {
	if l == nil {
		return nil
	}
	return l.head.next
}

func (l *DoublyLinkedList[T]) End() *DLLNode[T] {
	if l == nil {
		return nil
	}
	return &l.tail
}

func (l *DoublyLinkedList[T]) Insert(at *DLLNode[T], val *T) *DLLNode[T] {
	if l == nil || at == nil {
		return nil
	}

	newNode := &DLLNode[T]{
		Data: val,
		next: at,
		prev: at.prev,
	}

	at.prev.next = newNode
	at.prev = newNode
	l.length++

	return newNode
}

func (l *DoublyLinkedList[T]) Remove(at *DLLNode[T]) *DLLNode[T] {
	if l == nil || at == nil || at == &l.head || at == &l.tail {
		return nil
	}

	next := at.next

	at.prev.next = next
	next.prev = at.prev
	l.length--

	at.next = nil
	at.prev = nil
	at.Data = nil

	return next
}

func (l *DoublyLinkedList[T]) Find(from, to *DLLNode[T], data *T, comp func(*T, *T) int) *DLLNode[T] {
	if l == nil || from == nil || to == nil {
		return nil
	}

	for from != to {
		if from.Data != nil && comp(from.Data, data) == 0 {
			return from
		}
		from = from.Next()
	}
	return from
}

func (l *DoublyLinkedList[T]) ForEach(from, to *DLLNode[T], do func(*T) bool) *DLLNode[T] {
	if l == nil || from == nil || to == nil {
		return nil
	}

	for from != to {
		if from.Data != nil && !do(from.Data) {
			return from
		}
		from = from.Next()
	}
	return from
}

func (l *DoublyLinkedList[T]) PushFront(val *T) *DLLNode[T] {
	if l == nil {
		return nil
	}

	return l.Insert(l.Begin(), val)
}

func (l *DoublyLinkedList[T]) PushBack(val *T) *DLLNode[T] {
	if l == nil {
		return nil
	}

	return l.Insert(l.End(), val)
}

func (l *DoublyLinkedList[T]) PopFront() *T {
	if l == nil || l.length == 0 {
		return nil
	}

	val := l.Begin().Data

	if ret := l.Remove(l.Begin()); ret == nil {
		return nil
	}
	return val
}

func (l *DoublyLinkedList[T]) PopBack() *T {
	if l == nil || l.length == 0 {
		return nil
	}

	val := l.End().Prev().Data

	if ret := l.Remove(l.End().Prev()); ret == nil {
		return nil
	}
	return val
}

func (l *DoublyLinkedList[T]) Splice(at, begin, end *DLLNode[T]) *DLLNode[T] {
	if l == nil || at == nil || begin == nil || end == nil {
		return nil
	}

	last := end.Prev()

	begin.prev.next = end
	end.prev = begin.prev

	at.prev.next = begin
	begin.prev = at.prev

	last.next = at
	at.prev = last

	return last
}

func (l *DoublyLinkedList[T]) MultiFind(from, to *DLLNode[T], data *T, comp func(*T, *T) int, ret *DoublyLinkedList[T]) {
	if l == nil || from == nil || to == nil || ret == nil {
		return
	}

	for from != to {
		if from.Data != nil && comp(from.Data, data) == 0 {
			ret.PushBack(from.Data)
		}
		from = from.Next()
	}
}

func (l *DoublyLinkedList[T]) ToSlice() []T {
	if l == nil || l.length == 0 {
		return []T{}
	}

	s := make([]T, 0, l.length)
	l.ForEach(l.Begin(), l.End(), func(val *T) bool {
		if val != nil {
			s = append(s, *val)
		}
		return true
	})
	return s
}
