package skiplist

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

var (
	MAX_LEVEL  = 32
	SKIPLIST_P = 0.25
)

type SkipList[T constraints.Ordered] struct {
	len       int
	level     int
	head      *Node[T]
	skipListP float64
}

type Node[T constraints.Ordered] struct {
	Value any
	Key   T
	next  []*Node[T]
}

func New[T constraints.Ordered]() *SkipList[T] {
	defaultSL := &SkipList[T]{
		level:     MAX_LEVEL,
		skipListP: SKIPLIST_P,
		head: &Node[T]{
			next: make([]*Node[T], MAX_LEVEL),
		},
	}
	return defaultSL
}

func (s *SkipList[T]) Len() int {
	return s.len
}

func (sl *SkipList[T]) Get(key T) any {
	if sl.head == nil || sl.len == 0 {
		return nil
	}
	node := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		// 从最高层开始往下查找
		for node.next[i] != nil && node.next[i].Key < key {
			// 当前层往右查找
			node = node.next[i]
		}
		if node.next[i] != nil && node.next[i].Key == key {
			return node.next[i].Value
		}
	}
	return nil
}

func (sl *SkipList[T]) randomLevel() int {
	level := 0

	for rand.Float64() < sl.skipListP && level < sl.level {
		level += 1
	}
	return level
}

func (sl *SkipList[T]) Insert(key T, value any) {
	if value == nil {
		return
	}
	level := sl.randomLevel()
	newNode := &Node[T]{
		Value: value,
		Key:   key,
		next:  make([]*Node[T], sl.level),
	}
	node := sl.head
	for i := level; i >= 0; i-- {
		// 向右遍历，直到右侧节点不存在或者 key 值大于 key
		for node.next[i] != nil && node.next[i].Key < key {
			node = node.next[i]
		}

		// 调整指针关系，完成新节点的插入
		newNode.next[i] = node.next[i]
		node.next[i] = newNode
	}
	sl.len += 1
}

func (sl *SkipList[T]) Delete(key T) {
	if sl.head == nil || sl.len == 0 {
		return
	}
	node := sl.head
	for i := sl.level; i >= 0; i-- {
		// 从最高层开始往下查找
		for node.next[i] != nil && node.next[i].Key < key {
			// 当前层往右查找
			node = node.next[i]
		}
		if node.next[i] != nil && node.next[i].Key == key {
			// 找到待删除节点
			node.next[i] = node.next[i].next[i]
		}
	}
	sl.len -= 1
}

func (sl *SkipList[T]) Range(start, end T) []any {

	if sl.head == nil || sl.len == 0 {
		return nil
	}
	node := sl.head
	for i := sl.level; i >= 0; i-- {
		// 从最高层开始往下查找
		for node.next[i] != nil && node.next[i].Key < start {
			// 当前层往右查找
			node = node.next[i]
		}
	}
	var result []any
	for node != nil && node.Key <= end {
		result = append(result, node.Value)
		node = node.next[0]
	}
	return result
}
