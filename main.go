package main

import (
	"algor3/cum_list"
	"fmt"
)

func main(){
	linkedListObj := cum_list.NewLinkedList()
	listTest(linkedListObj)
}

func listTest(listObj cum_list.IList){
	for i := 0; i < 10; i++ {
		listObj.Add(i + 1)
	}
	listObj.AddLast(12345)
	fmt.Println(listObj.Get(9))
	fmt.Println(listObj.Set(9, 123))
	fmt.Println(listObj.IndexOf(3))
	fmt.Println(listObj.Contains(1231))
	//fmt.Println(listObj.RemoveFirst())
	fmt.Println(listObj.RemoveLast())
	//fmt.Println(listObj.Set(8, 123))

	fmt.Println(listObj)
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
