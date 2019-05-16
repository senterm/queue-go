package queue

import (
	"sync"
)

type NormalQueue struct {
	head     *Node
	tail     *Node
	size     int
	capacity int
	mutex    sync.RWMutex
}

func NewNormalQueue(capacity int) (*NormalQueue, QUEUE_RTV) {
	if capacity <= 0 {
		return nil, CAPACITY_INVALID
	}
	head := &Node{
		value:    nil,
		previous: nil,
	}
	tail := &Node{
		value:    nil,
		previous: head,
	}
	head.next = tail
	return &NormalQueue{
		head:     head,
		tail:     tail,
		capacity: capacity,
	}, SUCCESS
}

func (q *NormalQueue) Size() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.size
}

func (q *NormalQueue) Capacity() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.capacity
}

func (q *NormalQueue) Head() (interface{}, QUEUE_RTV) {
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

func (q *NormalQueue) Tail() (interface{}, QUEUE_RTV) {
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

func (q *NormalQueue) Push(value interface{}) QUEUE_RTV {
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
	if q.size == 0 {
		q.head.next = node
	}
	node.previous = q.tail.previous
	node.next = q.tail
	q.tail.previous.next = node
	q.tail.previous = node
	q.size++
	return SUCCESS
}

func (q *NormalQueue) Pop() (interface{}, QUEUE_RTV) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.size == 0 {
		return nil, QUEUE_EMPTY
	}

	result := q.head.next
	q.head.next = result.next
	result.next = nil
	result.previous = nil
	q.size--

	return result.value, SUCCESS
}
