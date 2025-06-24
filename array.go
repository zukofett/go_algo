package main

import (
	"unsafe"
)

type Array[T any] struct {
	data    unsafe.Pointer
	backing []byte
	size    int
}

func findOffset[T any](i int) uintptr {
	var elem T
	return unsafe.Sizeof(elem) * uintptr(i)
}

func NewArray[T any](numElements int) *Array[T] {
	if numElements == 0 {
		return &Array[T]{}
	}

	backing := make([]byte, findOffset[T](numElements))

	return &Array[T]{
		data:    unsafe.Pointer(&backing[0]),
		size:    numElements,
		backing: backing,
	}
}

func (a *Array[T]) Get(i int) T {
	if i < 0 || i >= a.size {
		panic("segmentation fault")
	}

	offset := findOffset[T](i)
	return *(*T)(unsafe.Pointer(uintptr(a.data) + offset))
}

func (a *Array[T]) Set(val T, i int) {
	if i < 0 || i >= a.size {
		panic("segmentation fault")
	}

	offset := findOffset[T](i)
	*(*T)(unsafe.Pointer(uintptr(a.data) + offset)) = val
}
