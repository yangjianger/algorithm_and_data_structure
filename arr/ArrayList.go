package arr

import (
	"strconv"
	"strings"
)

type arrayList struct {
	size int
	arrayData []int
}

const (
	_DEFAULT_CAPACITY = 10 //默认容量
	_DEFAULT_CAPACITY_SPEED = 2 //默认扩容倍数
	_ELEMENT_NOT_FOUND = -1 //默认元素没找到
)

func NewArrayList(capacity int) IArrayList{
	if capacity < 10{
		capacity = _DEFAULT_CAPACITY
	}
	return &arrayList{
		arrayData: make([]int, capacity),
		size: 0,
	}
}

//元素的数量
func (a *arrayList) Size() int {
	return a.size
}

//是否为空
func (a *arrayList) IsEmpty() bool {
	return a.size == 0
}

//是否包含某个元素
func (a *arrayList) Contains(element int) bool {
	for i := 0; i < a.size; i++ {
		if(element == a.arrayData[i]){
			return true
		}
	}
	return false
}

//判断数组是否已满
func (a *arrayList) isFull() bool{
	if(a.size == cap(a.arrayData)){
		return true
	}
	return false
}

//设置扩容
func (a *arrayList) setResize(){
	//2 倍扩容
	a.resize(a.size * _DEFAULT_CAPACITY_SPEED)
}

//扩容
func (a *arrayList) resize(capacity int){
	newArrayData := make([]int, capacity)
	for i := 0; i < a.size; i++ {
		newArrayData[i] = a.arrayData[i]
	}
	a.arrayData = newArrayData
}

//判断是否越界
func (a *arrayList) isOutOfRange(index int){
	if index < 0 || index >= a.size{
		panic("out of range")
	}
}

func (a *arrayList) isAddOutOfRange(index int){
	if index < 0 || index > a.size{
		panic("out of range")
	}
}

//添加元素到最后面
func (a *arrayList) Add(element int) {
	a.AddIndex(a.size, element)
}

//添加元素到最前面
func (a *arrayList) AddFirst(element int) {
	a.AddIndex(0, element)
}

//添加元素到最后面
func (a *arrayList) AddLast(element int) {
	a.AddIndex(a.size, element)
}

//往index位置添加元素
func (a *arrayList) AddIndex(index int, element int) {

	a.isAddOutOfRange(index)

	//已满，需要扩容
	if a.isFull(){
		a.setResize()
	}
 
	for i := a.size; i > index; i-- {
		a.arrayData[i] = a.arrayData[i-1]
	}

	a.arrayData[index] = element
	a.size++
}

//返回index 位置的对应元素
func (a *arrayList) Get(index int) int {
	a.isOutOfRange(index)

	return a.arrayData[index]
}

//设置index位置的元素
func (a *arrayList) Set(index int, element int) int {
	a.isOutOfRange(index)

	oldEle := a.arrayData[index]
	a.arrayData[index] = element
	return oldEle
}

//删除index 位置对应的元素
func (a *arrayList) Remove(index int) int {
	a.isOutOfRange(index)

	if a.IsEmpty(){
		panic("array is empty")
	}

	oldEle := a.arrayData[index]
	for i := index; i < a.size-1; i++ {
		a.arrayData[i] = a.arrayData[i+1]
	}
	a.size --
	return oldEle
}

//删除第一个元素
func (a *arrayList) RemoveFirst() int{
	return a.Remove(0)
}

//删除最后一个元素
func (a *arrayList) RemoveLast() int{
	return a.Remove(a.size-1)
}

//查看元素位置
func (a *arrayList) IndexOf(element int) int {
	for i := 0; i < a.size; i++ {
		if a.arrayData[i] == element{
			return i
		}
	}

	return _ELEMENT_NOT_FOUND
}

//清除所有元素
func (a *arrayList) Clear() {
	a.size = 0
}

//输出数组
func (a *arrayList) String() string{
	sb := strings.Builder{}
	sb.WriteString("array_list: ")
	for i := 0; i < a.size; i++ {
		if i != 0{
			sb.WriteString(",")
		}
		sb.WriteString(strconv.Itoa(a.arrayData[i]))
	}

	return sb.String()
}

