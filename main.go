package main

import (
	"algor3/cum_list"
	"fmt"
)

func main(){
	linkedListObj := cum_list.NewLinkedList()

	for i := 0; i < 10; i++ {
		linkedListObj.Add(i + 1)
	}
	linkedListObj.AddLast(12345)
	fmt.Println(linkedListObj.Get(9))
	fmt.Println(linkedListObj.Set(9, 123))
	fmt.Println(linkedListObj.IndexOf(3))
	fmt.Println(linkedListObj.Contains(1231))
	//fmt.Println(linkedListObj.RemoveFirst())
	fmt.Println(linkedListObj.RemoveLast())
	//fmt.Println(linkedListObj.Set(8, 123))

	fmt.Println(linkedListObj)


}

func mainArr(){

	arrayList := cum_list.NewArrayList(10)
	for i := 0; i < 10; i++ {
		arrayList.Add(i)
	}
	arrayList.Add(111)
	arrayList.Add(345)
	arrayList.Add(235)
	//arrayList.Remove(9)
	arrayList.Set(9,  98)
	arrayList.AddIndex(9, 123)
	//arrayList.Clear()
	arrayList.AddFirst(90)
	arrayList.AddLast(901)
	arrayList.RemoveFirst()
	arrayList.RemoveLast()
	fmt.Println(arrayList.String())
	fmt.Println(arrayList.Size())
	//fmt.Println(arrayList.Get(9))

}
