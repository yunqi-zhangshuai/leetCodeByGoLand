package linkedList

// SingleNode
// 链表节点
type SingleNode struct {
	Item interface{}
	Next *SingleNode
}

// NewSingleNode
// new 链表节点
func NewSingleNode(item interface{}) *SingleNode {
	return &SingleNode{Item: item}
}

//----------------------------------------------------------------------

// SingleList
// 单链表
type SingleList struct {
	// 链表头尾节点
	head, tail *SingleNode
	size       uint32
}

// Append
// 追加到链表尾部
func (l *SingleList) Append(node *SingleNode) {
	if l.head == nil {
		l.head = node
		l.tail = node
		l.size++
		return
	}
	l.tail.Next = node
	l.size++
	l.tail = node
}

// GetHead
// 获取链表头
func (l *SingleList) GetHead() *SingleNode {
	return l.head
}

// GetTail
// 获取链表尾部节点
func (l *SingleList) GetTail() *SingleNode {
	return l.tail
}

// RemoveNode
// 删除指定位置链表节点
func (l *SingleList) RemoveNode(num uint32) (bool, error) {

	curNode := l.head

	var index uint32
	for curNode != nil {
		index++
		if index == num {

		}

		curNode = curNode.Next
	}

	return true, nil
}

// InsertNode
// 指定位置插入节点
func (l *SingleList) InsertNode(sort uint32, node *SingleNode) (bool, error) {

	return true, nil
}

// Length
// 获取链表长度
func (l *SingleList) Length() uint32 {
	return l.size
}

// BatchAppend
// 批量追加链表节点
func (l *SingleList) BatchAppend(nodes ...*SingleNode) {
	for _, node := range nodes {
		l.Append(node)
	}
}
