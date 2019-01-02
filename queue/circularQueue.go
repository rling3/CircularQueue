package queue

import (
	"github.com/sirupsen/logrus"
)

type qError string

const (
	DEFAULT_SIZE = 4
	E_EMPTY = qError("Queue is empty!")
)

type Queue interface {
	Push(interface{})
	Pop() (interface{}, error)
	Size() int
	Length() int
}

func (e qError) Error() string {
	return string(e)
}

type queue struct {
	head int
	size int
	data []interface{}
}

func New(data []interface{}) Queue {
	logrus.SetLevel(logrus.DebugLevel)
	if data != nil {
		return &queue{
			head: 0,
			size: len(data),
			data: data,
		}
	}
	return &queue{
		head: 0,
		size: 0,
		data: make([]interface{}, DEFAULT_SIZE),
	}
}

// If more than half of queue is filled, double the queue length
func (q *queue) Push(item interface{}) {
	q.debug("Before push")
	if q.size+1 > len(q.data)/2 {
		logrus.Debugf("Current size is %d. Resizing current length %d to new length %d", q.size, len(q.data), 2*len(q.data))
		// make double length queue
		newQueue := make([]interface{}, len(q.data) * 2)
		switch {
		case q.head + q.size > len(q.data):
			wrapLength := q.head + q.size - len(q.data)
			copy(newQueue, q.data[q.head:])
			copy(newQueue[:len(q.data[q.head:])], q.data[:wrapLength-1])
		case q.head + q.size < len(q.data):
			copy(newQueue, q.data[q.head:])
		}
		q.data = newQueue
	}
	q.data[(q.size + q.head) % len(q.data)] = item
	q.size++
	q.debug("After push")
}

func (q *queue) Pop() (interface{}, error) {
	q.debug("Before pop")
	switch {
	case q.size <= 0:
		return nil, E_EMPTY
	case q.size-1 <= len(q.data)/4 && len(q.data) > DEFAULT_SIZE:
		val := q.data[q.head]
		newQueue := make([]interface{}, len(q.data)/2)
		switch {
		case q.head+q.size > len(q.data):
			wrapLength := q.head + q.size - len(q.data)
			copy(newQueue, q.data[q.head+1:])
			copy(newQueue[:len(q.data[q.head-1:])], q.data[:wrapLength-1])
			q.data = newQueue
		case q.head+q.size < len(q.data):
			copy(newQueue, q.data[q.head+1:])
			q.data = newQueue
		}
		q.head = 0
		q.size--
		q.debug("After pop")
		return val, nil
	default:
		val := q.data[q.head]
		q.head++
		q.size--
		q.debug("After pop")
		return val, nil
	}
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) Length() int {
	return len(q.data)
}

func (q *queue) debug(step string) {
	logrus.Debugf(step + "| Data: %v | Length: %d | Size: %d | Head: %d", q.data, len(q.data), q.size, q.head)
}
