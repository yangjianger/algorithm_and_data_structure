package cum_list

import (
	"strconv"
	"strings"
)

type linkedList struct {
	//虚拟头结点
	dummyHead *node
	size int
}

func NewLinkedList() IList{
	return &linkedList{dummyHead: newNode(0, nil, nil), size: 0,}
}

//元素的数量
func (l *linkedList) Size() int {
	return l.size
}

//是否为空
func (l *linkedList) IsEmpty() bool {
	return l.size == 0
}

//是否包含某个元素
func (l *linkedList) Contains(element int) bool {

	cur := l.dummyHead.nextNode
	for cur != nil {

		if cur.val == element{
			return true
		}
		cur = cur.nextNode
	}

	return false
}

//添加元素到最前面
func (l *linkedList) Add(element int) {
	l.AddFirst(element)
}

//链表是否为空异常
func (l *linkedList) isEmptyPanic(){
	if l.IsEmpty(){
		panic("linked list is empty")
	}
}

//链表范围判断异常
func (l *linkedList) linedListRangPanic(index int){

	l.isEmptyPanic()
	if index >= l.size {
		panic("linked list is range out")
	}

	if index < 0{
		panic("index is invalid")
	}
}

//链表添加范围判断异常
func (l *linkedList) linedListAddRangPanic(index int){

	if index < 0{
		panic("index is invalid")
	}

	if index > l.size {
		panic("linked list is range out")
	}
}
//返回第一个 位置的对应元素
func (l *linkedList) GetFirst() int{
	return l.Get(0)
}

//返回最后一个 位置的对应元素
func (l *linkedList) GetLast() int{
	return l.Get(l.size - 1)
}

//返回index 位置的对应元素
func (l *linkedList) Get(index int) int {
	l.linedListRangPanic(index)

	cur := l.dummyHead.nextNode
	for i := 0; i < index; i++ {
		cur = cur.nextNode
	}

	return cur.val
	
}

//设置index位置的元素
func (l *linkedList) Set(index int, element int) int {
	l.linedListRangPanic(index)

	cur := l.dummyHead.nextNode
	for i := 0; i < index; i++ {
		cur = cur.nextNode
	}
	oldVal := cur.val
	cur.val = element

	return oldVal
}

//往index位置添加元素
func (l *linkedList) AddIndex(index int, element int) {
	l.linedListAddRangPanic(index)

	cur := l.dummyHead

	for i := 0; i < index; i++ {
		cur = cur.nextNode
	}
	
	cur.nextNode = newNode(element, cur.nextNode, cur)
	l.size ++
}

//添加元素到最前面
func (l *linkedList) AddFirst(element int) {
	l.AddIndex(0, element)
}

//添加元素到最后面
func (l *linkedList) AddLast(element int) {
	l.AddIndex(l.size, element)
}

//删除index 位置对应的元素
func (l *linkedList) Remove(index int) int {
	l.linedListRangPanic(index)

	prev := l.dummyHead

	for i := 0; i < index; i++ {
		prev = prev.nextNode
	}

	oldVal := prev.nextNode.val
	prev.nextNode =  prev.nextNode.nextNode
	l.size --
	return oldVal

}

//删除最后一个元素
func (l *linkedList) RemoveLast() int {
	return l.Remove(l.size - 1)
}

//删除第一个元素
func (l *linkedList) RemoveFirst() int {
	return l.Remove(0)
}

//查看元素位置
func (l *linkedList) IndexOf(element int) int {
	l.isEmptyPanic()

	cur := l.dummyHead
	for i := 0; i < l.size; i++ {
		cur = cur.nextNode
		if cur.val == element{
			return i
		}
	}

	return -1

}

//清除所有元素
func (l *linkedList) Clear() {
	l.size = 0
}

//输出链表
func (l *linkedList) String() string {
	sb := strings.Builder{}
	sb.WriteString("linked_list: ")
	cur := l.dummyHead

	for i := 0; i < l.size; i++ {
		if i != 0{
			sb.WriteString(",")
		}
		sb.WriteString(strconv.Itoa(cur.nextNode.val))
		cur = cur.nextNode
	}

	return sb.String()

}