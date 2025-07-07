package array

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestNewArray(t *testing.T) {
	elements := 5
	arr := NewArray[int](elements)

	if arr == nil {
		t.Fatal("failed to allocate array")
	}

	if arr.size != elements {
		t.Errorf("expected size %d; got %d", elements, arr.size)
	}

	if arr.data == nil {
		t.Fatal("array data pointer is nil")
	}

	if arr.backing == nil {
		t.Fatal("array backing array is nil")
	}

	expectedSize := int(unsafe.Sizeof(int(0))) * elements
	if len(arr.backing) != expectedSize {
		t.Errorf("expected backing array size %d; got %d", expectedSize, len(arr.backing))
	}
}

func TestZeroSizeArray(t *testing.T) {
	elements := 0
	arr := NewArray[int](elements)

	if arr == nil {
		t.Fatal("failed to allocate array")
	}

	if arr.size != elements {
		t.Errorf("expected size %d; got %d", elements, arr.size)
	}

	if arr.data != nil {
		t.Fatal("array data pointer is not nil")
	}

	if arr.backing != nil {
		t.Fatal("array backing array is not nil")
	}

	expectedSize := int(unsafe.Sizeof(int(0))) * elements
	if len(arr.backing) != expectedSize {
		t.Errorf("expected backing array size %d; got %d", expectedSize, len(arr.backing))
	}

	defer func() {
		if err := recover(); err == nil {
			t.Fatal("expected to panic at access but didn't")
		}
	}()

	_ = arr.Get(0)
}
func TestSetGet(t *testing.T) {
	arr := NewArray[int](3)
	vals := []int{100, 52, 0}

	for i, v := range vals {
		arr.Set(v, i)
	}

	for i, want := range vals {
		got := arr.Get(i)
		if want != got {
			t.Errorf("in index %d, got: %d; expected: %d", i, got, want)
		}
	}
}

func TestSetGetStrings(t *testing.T) {
	arr := NewArray[string](3)
	vals := []string{"hello", "world", "gophers"}

	for i, v := range vals {
		arr.Set(v, i)
	}

	for i, want := range vals {
		got := arr.Get(i)
		if want != got {
			t.Errorf("in index %d, got: %q; expected: %q", i, got, want)
		}
	}
}

func TestSetGetCustomType(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	arr := NewArray[person](3)
	vals := []person{
		{
			name: "idan",
			age:  100,
		},
		{
			name: "other",
			age:  -1,
		},
		{
			name: "",
			age:  0,
		},
	}

	for i, v := range vals {
		arr.Set(v, i)
	}

	for i, want := range vals {
		got := arr.Get(i)
		if want.name != got.name || want.age != got.age {
			t.Errorf("in index %d, got: %v; expected: %v", i, got, want)
		}
	}
}

func TestZeroValues(t *testing.T) {
	arr := NewArray[int](3)

	for i := range arr.size {
		if got := arr.Get(i); got != 0 {
			t.Errorf("in index %d, got: %d; expected zero", i, got)
		}
	}
}

func TestGetBoundsCheck(t *testing.T) {
	arr := NewArray[int](3)
	cases := []struct {
		name  string
		index int
	}{
		{
			name:  "negative index",
			index: -1,
		},
		{
			name:  "large index",
			index: 1000,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Fatal("expected panic but got nil")
				}
			}()
			arr.Get(tt.index)
		})
	}
}

func TestSetBoundsCheck(t *testing.T) {
	arr := NewArray[int](3)
	cases := []struct {
		name  string
		index int
	}{
		{
			name:  "negative index",
			index: -1,
		},
		{
			name:  "large index",
			index: 1000,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Fatal("expected panic but got nil")
				}
			}()
			arr.Set(15, tt.index)
		})
	}
}

func TestMemoryIsAlive(t *testing.T) {
	elements := 1000
	arr := NewArray[int](elements)

	for i := range elements {
		arr.Set(i*i, i)
	}

	runtime.GC()
	runtime.GC()

	for i := range elements {
		if got := arr.Get(i); got != i*i {
			t.Errorf("got: %d; want: %d", got, i*i)
			break
		}
	}
}

func BenchmarkArraySet(b *testing.B) {
	arr := NewArray[int](1000)
	for i := 0; i < b.N; i++ {
		arr.Set(i, i%1000)
	}
}

func BenchmarkSliceSet(b *testing.B) {
	slice := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		slice[i%1000] = i
	}
}

func BenchmarkArrayGet(b *testing.B) {
	arr := NewArray[int](1000)
	for i := range 1000 {
		arr.Set(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = arr.Get(i % 1000)
	}
}

func BenchmarkSliceGet(b *testing.B) {
	slice := make([]int, 1000)
	for i := range 1000 {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = slice[i%1000]
	}
}
