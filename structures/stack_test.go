package structures

import (
	"testing"
)

func Test_StackIntAdd(t *testing.T) {
	stack, err := NewStackInt()
	if err != nil {
		t.Errorf("Unexpected error on stack creation: %s", err.Error())
	}

	stack.Add(1)
	stack.Add(2)
	stack.Add(5)
	stack.Add(8)
	stack.Add(3)

	expectedStackContent := "1,2,5,8,3"
	if stack.Join() != expectedStackContent {
		t.Errorf("Expected: %v, but got %v", expectedStackContent, stack.List())
	}
}

func Test_StackIntAddPop(t *testing.T) {
	stack, err := NewStackInt()
	if err != nil {
		t.Errorf("Unexpected error on queue creation: %s", err.Error())
	}

	stack.Add(1)
	stack.Add(2)
	stack.Add(5)
	stack.Add(8)
	stack.Add(3)

	lastElement, err := stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error on queue pop: %s", err.Error())
	}
	if *lastElement != 3 {
		t.Errorf("Expected value: 3, got: %d", *lastElement)
	}

	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()

	emptyStackResult, err := stack.Pop()
	if err == nil || emptyStackResult != nil {
		t.Errorf("Incorrect result for an empty queue: %v, %v", emptyStackResult, err)
	}
}
