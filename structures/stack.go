package structures

import (
	"errors"
	"fmt"
	"strings"
)

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() (Stack[T], error) {
	return Stack[T]{data: []T{}}, nil
}

func (q *Stack[T]) Add(el T) error {
	q.data = append(q.data, el)
	return nil
}

func (q *Stack[T]) Pop() (*T, error) {
	if len(q.data) == 0 {
		return nil, errors.New("empty stack")
	}
	lastElement := q.data[len(q.data)-1]
	data := q.data[0 : len(q.data)-1]
	q.data = data
	return &lastElement, nil
}

func (q *Stack[T]) List() []T {
	return q.data
}

func (q *Stack[T]) Join() string {
	elementsAsStrings := []string{}
	for _, i := range q.data {
		elementsAsStrings = append(elementsAsStrings, fmt.Sprintf("%v", i))
	}
	return strings.Join(elementsAsStrings, ",")
}
