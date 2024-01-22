package skiplist

import (
	"reflect"
	"testing"
)

func TestPutInt1(t *testing.T) {
	sl := New[int]()
	sl.Put(1, 2)
	sl.Put(2, 2)
	sl.Put(10, 2)
	sl.Put(3, 2)
	a := sl.Get(1)
	b := sl.Get(10)
	if !reflect.DeepEqual(a, 2) {
		t.Errorf("Put() = %v, want %v", a, 2)
	}
	if !reflect.DeepEqual(b, 2) {
		t.Errorf("Put() = %v, want %v", b, 2)
	}
}

func TestPutStruct1(t *testing.T) {
	type MyStruct struct {
		A int
		B int
		C string
	}
	sl := New[string]()
	sl.Put("a", MyStruct{A: 1, B: 2, C: "3"})
	sl.Put("b", MyStruct{A: 2, B: 1, C: "3"})
	sl.Put("c", MyStruct{A: 1, B: 2, C: "1"})
	a := sl.Get("a")
	c := sl.Get("c")

	if !reflect.DeepEqual(a, MyStruct{A: 1, B: 2, C: "3"}) {
		t.Errorf("Put() = %v, want %v", a, 2)
	}
	if !reflect.DeepEqual(c, MyStruct{A: 1, B: 2, C: "1"}) {
		t.Errorf("Put() = %v, want %v", c, 2)
	}
	sl.Put("a", MyStruct{A: 11, B: 2, C: "3"})
	a = sl.Get("a")
	if !reflect.DeepEqual(a, MyStruct{A: 11, B: 2, C: "3"}) {
		t.Errorf("Put() = %v, want %v", a, 2)
	}
}

func TestDeleteInt1(t *testing.T) {
	sl := New[int]()
	sl.Put(1, 2)
	sl.Put(2, 2)
	sl.Put(10, 2)
	sl.Put(3, 2)
	sl.Delete(1)
	a := sl.Get(1)
	if !reflect.DeepEqual(a, nil) {
		t.Errorf("Put() = %v, want %v", a, nil)
	}
	sl.Put(1, 2)
	// sl.Delete(1)
	a = sl.Get(1)
	if !reflect.DeepEqual(a, 2) {
		t.Errorf("Put() = %v, want %v", a, nil)
	}
}

func TestRangeInt1(t *testing.T) {
	sl := New[int]()
	sl.Put(1, 2)
	sl.Put(2, 3)
	sl.Put(10, 2)
	sl.Put(3, 2)
	a := sl.Range(1, 10)
	if !reflect.DeepEqual(a, []any{2, 3, 2, 2}) {
		t.Errorf("Put() = %v, want %v", a, []any{2, 3, 2, 2})
	}
}

func TestRangeInt2(t *testing.T) {
	sl := New[int]()
	sl.Put(1, 2)
	sl.Put(2, 3)
	sl.Put(10, 2)
	sl.Put(3, 2)
	a := sl.Range(3, 9)
	if !reflect.DeepEqual(a, []any{2}) {
		t.Errorf("Put() = %v, want %v", a, []any{2})
	}

	a = sl.Range(3, 11)
	if !reflect.DeepEqual(a, []any{2, 2}) {
		t.Errorf("Put() = %v, want %v", a, []any{2, 2})
	}
}
