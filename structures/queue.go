package structures

import (
	"errors"
	"fmt"
	"strings"
)

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() (Queue[T], error) {
	return Queue[T]{data: []T{}}, nil
}

func (q *Queue[T]) Add(el T) error {
	tmpQueue := []T{el}
	q.data = append(tmpQueue, q.data...)
	return nil
}

func (q *Queue[T]) Pop() (*T, error) {
	if len(q.data) == 0 {
		return nil, errors.New("empty queue")
	}
	lastElement := q.data[len(q.data)-1]
	data := q.data[0 : len(q.data)-1]
	q.data = data
	return &lastElement, nil
}

func (q *Queue[T]) List() []T {
	return q.data
}

func (q *Queue[T]) Join() string {
	elementsAsStrings := []string{}
	for _, i := range q.data {
		elementsAsStrings = append(elementsAsStrings, fmt.Sprintf("%v", i))
	}
	return strings.Join(elementsAsStrings, ",")
}
