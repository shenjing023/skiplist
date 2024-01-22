package main

import (
	"fmt"

	"github.com/shenjing023/skiplist"
)

func main() {
	sl := skiplist.New[int]()
	sl.Put(1, 2)
	sl.Put(2, 3)
	sl.Put(10, 2)
	sl.Put(3, 2)
	a := sl.Get(1)
	b := sl.Get(10)
	fmt.Printf("%v, %v\n", a, b)
	sl.Delete(10)
	c := sl.Get(10)
	fmt.Printf("%v, %v\n", a, c)
}
