package arr

//动态数组接口
type IArrayList interface {
	Size() int //元素的数量
	IsEmpty() bool //是否为空
	Contains(element int) bool //是否包含某个元素
	Add(element int) //添加元素到最后面
	Get(index int) int //返回index 位置的对应元素
	Set(index int, element int) int //设置index位置的元素
	AddIndex(index int, element int) //往index位置添加元素
	AddFirst(element int) //添加元素到最前面
	AddLast(element int) //添加元素到最后面
	Remove(index int) int //删除index 位置对应的元素
	RemoveLast() int //删除最后一个元素
	RemoveFirst() int //删除第一个元素
	IndexOf(element int) int //查看元素位置
	Clear() //清除所有元素
	String() string //输出数组
}