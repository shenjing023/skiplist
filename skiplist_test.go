package skiplist

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *SkipList[int]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertInt1(t *testing.T) {
	sl := New[int]()
	sl.Insert(1, 2)
	sl.Insert(2, 2)
	sl.Insert(10, 2)
	sl.Insert(3, 2)
	a := sl.Get(1)
	b := sl.Get(10)
	if !reflect.DeepEqual(a, 2) {
		t.Errorf("Insert() = %v, want %v", a, 2)
	}
	if !reflect.DeepEqual(b, 2) {
		t.Errorf("Insert() = %v, want %v", b, 2)
	}
}
