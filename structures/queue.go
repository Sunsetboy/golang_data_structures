package structures

import (
	"errors"
	"strconv"
	"strings"
)

type QueueInt struct {
	data []int
}

func NewQueueInt() (QueueInt, error) {
	return QueueInt{data: []int{}}, nil
}

func (q *QueueInt) Add(el int) error {
	tmpQueue := []int{el}
	q.data = append(tmpQueue, q.data...)
	return nil
}

func (q *QueueInt) Pop() (*int, error) {
	if len(q.data) == 0 {
		return nil, errors.New("empty queue")
	}
	lastElement := q.data[len(q.data)-1]
	data := q.data[0 : len(q.data)-1]
	q.data = data
	return &lastElement, nil
}

func (q *QueueInt) List() []int {
	return q.data
}

func (q *QueueInt) Join() string {
	elementsAsStrings := []string{}
	for _, i := range q.data {
		elementsAsStrings = append(elementsAsStrings, strconv.Itoa(i))
	}
	return strings.Join(elementsAsStrings, ",")
}
