package cum_list

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	linkedListObj := &linkedList{dummyHead: newNode(0, nil, nil), size: 0,}
	for i := 0; i < 30; i++ {
		linkedListObj.Add(i + 1)
	}



	fmt.Println(linkedListObj)
	//linkedListObj.dummyHead.nextNode = linkedListObj.ReverseList(linkedListObj.dummyHead.nextNode)
	linkedListObj.AddIndex(10, 189)
	fmt.Println(linkedListObj)

	fmt.Println(linkedListObj.last.val)


}