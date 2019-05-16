package queue

type Queue interface {
	Size() QUEUE_RTV
	Capacity() QUEUE_RTV
	Head() (interface{}, QUEUE_RTV)
	Tail() (interface{}, QUEUE_RTV)
	Push(value interface{}) QUEUE_RTV
	Pop() (interface{}, QUEUE_RTV)
}
