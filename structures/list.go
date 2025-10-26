package structures

import (
	"strconv"
	"strings"
)

type listNode struct {
	val  int
	next *listNode
}

type ListInt struct {
	root *listNode
}

func NewListInt(val int) ListInt {
	return ListInt{&listNode{val: val}}
}

func (l *ListInt) List() []int {
	data := []int{}
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

func (l *ListInt) Add(el int) {
	currentNode := l.root
	for {
		if currentNode.next == nil {
			currentNode.next = &listNode{val: el}
			return
		}
		currentNode = currentNode.next
	}
}

func (l *ListInt) Join() string {
	elementsAsStrings := []string{}
	for _, i := range l.List() {
		elementsAsStrings = append(elementsAsStrings, strconv.Itoa(i))
	}
	return strings.Join(elementsAsStrings, ",")
}
