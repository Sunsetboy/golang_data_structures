package structures

import "testing"

func Test_ListIntAddJoin(t *testing.T) {
	list := NewListInt(1)
	list.Add(2)

	if list.Join() != "1,2" {
		t.Errorf("List error, expected: %v, got: %v", "1,2", list.Join())
	}
}
