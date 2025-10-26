package structures

import (
	"testing"
)

func Test_QueueIntAdd(t *testing.T) {
	queue, err := NewQueueInt()
	if err != nil {
		t.Errorf("Unexpected error on queue creation: %s", err.Error())
	}

	queue.Add(1)
	queue.Add(2)
	queue.Add(5)
	queue.Add(8)
	queue.Add(3)

	expectedQueueContent := "3,8,5,2,1"
	if queue.Join() != expectedQueueContent {
		t.Errorf("Expected: %v, but got %v", expectedQueueContent, queue.List())
	}
}

func Test_QueueIntAddPop(t *testing.T) {
	queue, err := NewQueueInt()
	if err != nil {
		t.Errorf("Unexpected error on queue creation: %s", err.Error())
	}

	queue.Add(1)
	queue.Add(2)
	queue.Add(5)
	queue.Add(8)
	queue.Add(3)

	firstElement, err := queue.Pop()
	if err != nil {
		t.Errorf("Unexpected error on queue pop: %s", err.Error())
	}
	if *firstElement != 1 {
		t.Errorf("Expected value: 3, got: %d", *firstElement)
	}

	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Pop()

	emptyQueueResult, err := queue.Pop()
	if err == nil || emptyQueueResult != nil {
		t.Errorf("Incorrect result for an empty queue: %v, %v", emptyQueueResult, err)
	}
}
