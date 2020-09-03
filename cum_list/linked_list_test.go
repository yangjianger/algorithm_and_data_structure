package cum_list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	linkedListObj := NewLinkedList()
	for i := 0; i < 10; i++ {
		linkedListObj.Add(i + 1)
	}

	if linkedListObj.Get(0) != 10{
		t.Error("Get error")
	}

	linkedListObj.Set(0, 123)
	if linkedListObj.Get(0) != 123{
		t.Error("Set error")
	}

	linkedListObj.AddLast(12345)
	if linkedListObj.Get(10) != 12345{
		t.Error("AddLast error")
	}

	if linkedListObj.RemoveFirst() != 123{
		t.Error("RemoveFirst error")
	}

}