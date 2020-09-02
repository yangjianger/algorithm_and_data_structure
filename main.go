package main

import (
	"algor3/arr"
	"fmt"
)

func main(){

}

func mainArr(){

	arrayList := arr.NewArrayList(10)
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
