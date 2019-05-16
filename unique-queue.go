package queue

import (
	"reflect"
	"sync"
)

type UniqueQueue struct {
	head     *Node
	tail     *Node
	size     int
	capacity int
	nodeMap  map[interface{}]bool
	mutex    sync.RWMutex
}

func NewUniqueQueue(capacity int) (*UniqueQueue, QUEUE_RTV) {
	if capacity <= 0 {
		return nil, CAPACITY_INVALID
	}
	head := &Node{
		value: nil,
	}
	tail := &Node{
		value:    nil,
		previous: head,
	}
	head.next = tail
	nodeMap := make(map[interface{}]bool)
	return &UniqueQueue{
		head:     head,
		tail:     tail,
		capacity: capacity,
		nodeMap:  nodeMap,
	}, SUCCESS
}

func (q *UniqueQueue) Size() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.size
}

func (q *UniqueQueue) Capacity() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.capacity
}

func (q *UniqueQueue) Head() (interface{}, QUEUE_RTV) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.size == 0 {
		return nil, QUEUE_EMPTY
	}
	if q.head.next == nil {
		return nil, VALUE_NIL
	}
	return q.head.next.value, SUCCESS
}

func (q *UniqueQueue) Tail() (interface{}, QUEUE_RTV) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.size == 0 {
		return nil, QUEUE_EMPTY
	}
	if q.tail.previous == nil {
		return nil, VALUE_NIL
	}
	return q.tail.previous.value, SUCCESS
}

func (q *UniqueQueue) Push(value interface{}) QUEUE_RTV {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.size == q.capacity {
		return REACH_CAPACITY
	}
	if value == nil {
		return VALUE_NIL
	}
	node := &Node{
		value: value,
	}
	if kind := reflect.TypeOf(value).Kind(); kind == reflect.Map || kind == reflect.Slice || kind == reflect.Func {
		return VALUE_TYPE_UNCOMPARABLE
	}
	if v, ok := q.nodeMap[value]; ok || v {
		return VALUE_EXISTED
	}
	if q.size == 0 {
		q.head.next = node
	}
	node.previous = q.tail.previous
	node.next = q.tail
	q.tail.previous.next = node
	q.tail.previous = node
	q.nodeMap[value] = true
	q.size++
	return SUCCESS
}

func (q *UniqueQueue) Pop() (interface{}, QUEUE_RTV) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.size == 0 {
		return nil, QUEUE_EMPTY
	}
	result := q.head.next
	delete(q.nodeMap, result.value)
	q.head.next = result.next
	result.next = nil
	result.previous = nil
	q.size--
	return result.value, SUCCESS
}
