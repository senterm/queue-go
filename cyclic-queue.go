package queue

import (
	"sync"
)

type CyclicQueue struct {
	head     int
	tail     int
	size     int
	capacity int
	nodes    []*Node
	mutex    sync.RWMutex
}

func NewCyclicQueue(capacity int) (*CyclicQueue, QUEUE_RTV) {
	if capacity <= 0 {
		return nil, CAPACITY_INVALID
	}
	nodes := make([]*Node, capacity, capacity)
	return &CyclicQueue{
		head:     -1,
		tail:     -1,
		capacity: capacity,
		nodes:    nodes,
	}, SUCCESS
}

func (q *CyclicQueue) Size() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.size
}

func (q *CyclicQueue) Capacity() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.capacity
}

func (q *CyclicQueue) Head() (interface{}, QUEUE_RTV) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.size == 0 {
		return nil, QUEUE_EMPTY
	}
	if q.nodes[q.head] == nil {
		return nil, VALUE_NIL
	}
	return q.nodes[q.head].value, SUCCESS
}

func (q *CyclicQueue) Tail() (interface{}, QUEUE_RTV) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.size == 0 {
		return nil, QUEUE_EMPTY
	}
	if q.nodes[q.tail] == nil {
		return nil, VALUE_NIL
	}
	return q.nodes[q.tail].value, SUCCESS
}

func (q *CyclicQueue) Push(value interface{}) QUEUE_RTV {
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
	index := (q.tail + 1) % cap(q.nodes)
	q.nodes[index] = node
	q.tail = index
	q.size++
	if q.size == 1 {
		q.head = index
	}
	return SUCCESS
}

func (q *CyclicQueue) Pop() (interface{}, QUEUE_RTV) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.size == 0 {
		return nil, QUEUE_EMPTY
	}
	result := q.nodes[q.head].value
	q.nodes[q.head] = nil
	index := (q.head + 1) % cap(q.nodes)
	q.head = index
	q.size--
	return result, SUCCESS
}
