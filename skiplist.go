package skiplist

import (
	"math/rand"
	"sync"
	"sync/atomic"

	"golang.org/x/exp/constraints"
)

var (
	MAX_LEVEL  = 32
	SKIPLIST_P = 0.25
)

type SkipList[T constraints.Ordered] struct {
	len       atomic.Int32
	level     int
	head      *Node[T]
	skipListP float64
	mutex     sync.RWMutex
}

type Node[T constraints.Ordered] struct {
	Value any
	Key   T
	next  []*Node[T]
}

func WithMaxLevel(maxLevel int) Option {
	return func(o *options) {
		o.MaxLevel = maxLevel
	}
}

func WithSkipListP(skipListP float64) Option {
	return func(o *options) {
		o.SkipListP = skipListP
	}
}

func New[T constraints.Ordered](opts ...Option) *SkipList[T] {

	op := NewOptions()

	for _, opt := range opts {
		opt(op)
	}

	sl := &SkipList[T]{
		level:     op.MaxLevel,
		skipListP: op.SkipListP,
		head: &Node[T]{
			next: make([]*Node[T], op.MaxLevel),
		},
		len: atomic.Int32{},
	}
	return sl
}

func (s *SkipList[T]) Len() int {
	return int(s.len.Load())
}

func (sl *SkipList[T]) Get(key T) any {
	sl.mutex.RLock()
	defer sl.mutex.RUnlock()
	if node := sl.find(key); node != nil {
		return node.Value
	}
	return nil
}

func (sl *SkipList[T]) find(key T) *Node[T] {
	if sl.head == nil || sl.len.Load() == 0 {
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
			return node.next[i]
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

func (sl *SkipList[T]) Put(key T, value any) {
	if value == nil {
		return
	}
	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	if node := sl.find(key); node != nil {
		node.Value = value
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
	sl.len.Add(1)
}

func (sl *SkipList[T]) Delete(key T) {
	if sl.head == nil || sl.len.Load() == 0 {
		return
	}
	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	node := sl.head
	for i := sl.level - 1; i >= 0; i-- {
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
	sl.len.Add(-1)
}

func (sl *SkipList[T]) Range(start, end T) []any {
	if sl.head == nil || sl.len.Load() == 0 {
		return nil
	}
	sl.mutex.RLock()
	defer sl.mutex.RUnlock()

	node := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		// 从最高层开始往下查找
		for node.next[i] != nil && node.next[i].Key < start {
			// 当前层往右查找
			node = node.next[i]
		}
		if node.next[i] != nil && node.next[i].Key == start {
			node = node.next[i]
			break
		}
	}
	var result []any
	for node != nil && node.Key <= end {
		result = append(result, node.Value)
		node = node.next[0]
	}
	return result
}
