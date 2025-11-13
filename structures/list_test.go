package structures

import "testing"

func Test_ListIntAddJoin(t *testing.T) {
	list := NewList(1)
	list.Add(2)

	if list.Join() != "1,2" {
		t.Errorf("List error, expected: %v, got: %v", "1,2", list.Join())
	}
}

func Test_ListStringAddJoin(t *testing.T) {
	list := NewList("hello")
	list.Add("world")

	if list.Join() != "hello,world" {
		t.Errorf("List error, expected: %v, got: %v", "hello,world", list.Join())
	}
}
