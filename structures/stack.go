package structures

import (
	"errors"
	"strconv"
	"strings"
)

type StackInt struct {
	data []int
}

func NewStackInt() (StackInt, error) {
	return StackInt{data: []int{}}, nil
}

func (q *StackInt) Add(el int) error {
	q.data = append(q.data, el)
	return nil
}

func (q *StackInt) Pop() (*int, error) {
	if len(q.data) == 0 {
		return nil, errors.New("empty stack")
	}
	lastElement := q.data[len(q.data)-1]
	data := q.data[0 : len(q.data)-1]
	q.data = data
	return &lastElement, nil
}

func (q *StackInt) List() []int {
	return q.data
}

func (q *StackInt) Join() string {
	elementsAsStrings := []string{}
	for _, i := range q.data {
		elementsAsStrings = append(elementsAsStrings, strconv.Itoa(i))
	}
	return strings.Join(elementsAsStrings, ",")
}
