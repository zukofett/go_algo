package main

type Node[T any] struct {
	Data *T
	next *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	if n == nil {
		return nil
	}
	return n.next
}

type SinglyLinkedList[T any] struct {
	head, tail *Node[T]
	length     int
}

func NewSLL[T any]() *SinglyLinkedList[T] {
	dummy := &Node[T]{}
	return &SinglyLinkedList[T]{
		head: dummy,
		tail: dummy,
	}
}

func (l *SinglyLinkedList[T]) Len() int {
	if l == nil {
		return 0
	}
	return l.length
}

func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l == nil || l.length == 0
}

func (l *SinglyLinkedList[T]) Begin() *Node[T] {
	if l == nil {
		return nil
	}
	return l.head
}

func (l *SinglyLinkedList[T]) End() *Node[T] {
	if l == nil {
		return nil
	}
	return l.tail
}

// insert inserts a new node before 'at'
func (l *SinglyLinkedList[T]) Insert(at *Node[T], val *T) *Node[T] {
	if l == nil || at == nil {
		return nil
	}

	newNode := &Node[T]{
		Data: at.Data,
		next: at.next,
	}

	at.Data = val
	at.next = newNode
	l.length++

	if at == l.tail {
		l.tail = newNode
	}
	return at
}

func (l *SinglyLinkedList[T]) Remove(at *Node[T]) *Node[T] {
	if l == nil || at == nil || at.next == nil {
		return nil
	}

	at.Data = at.next.Data
	at.next = at.next.next
	l.length--

	if l.tail == at.next {
		l.tail = at
	}

	return at
}

func (l *SinglyLinkedList[T]) Find(from, to *Node[T], data *T, comp func(*T, *T) int) *Node[T] {
	if l == nil || from == nil || to == nil {
		return nil
	}

	for from != to {
		if from.Data != nil && comp(from.Data, data) == 0 {
			return from
		}
		from = from.next
	}
	return nil
}

func (l *SinglyLinkedList[T]) ForEach(from, to *Node[T], do func(*T) bool) *Node[T] {
	if l == nil || from == nil || to == nil {
		return nil
	}

	for from != to {
		if from.Data != nil && !do(from.Data) {
			return from
		}
		from = from.next
	}
	return nil
}

func (l *SinglyLinkedList[T]) ToSilce() []T {
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
