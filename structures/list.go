package structures

import (
	"fmt"
	"strings"
)

type listNode[T any] struct {
	val  T
	next *listNode[T]
}

type List[T any] struct {
	root *listNode[T]
}

func NewList[T any](val T) List[T] {
	return List[T]{&listNode[T]{val: val}}
}

func (l *List[T]) List() []T {
	data := []T{}
	currentNode := l.root
	if currentNode == nil {
		return data
	}
	for {
		data = append(data, currentNode.val)
		if currentNode.next == nil {
			return data
		}
		currentNode = currentNode.next
	}
}

func (l *List[T]) Add(el T) {
	currentNode := l.root
	for {
		if currentNode.next == nil {
			currentNode.next = &listNode[T]{val: el}
			return
		}
		currentNode = currentNode.next
	}
}

func (l *List[T]) Join() string {
	elementsAsStrings := []string{}
	for _, i := range l.List() {
		elementsAsStrings = append(elementsAsStrings, fmt.Sprintf("%v", i))
	}
	return strings.Join(elementsAsStrings, ",")
}
