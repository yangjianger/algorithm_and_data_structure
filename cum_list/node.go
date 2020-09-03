package cum_list

type node struct {
	val int
	nextNode *node
	prevNode *node
}

func newNode(val int, nextNode *node, prevNode *node) *node{
	return &node{val: val, nextNode: nextNode, prevNode: prevNode}
}
