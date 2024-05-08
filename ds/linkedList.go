package ds

type Node struct {
	value any
	next  *Node
}

type LinkedList struct {
	First *Node
	Last  *Node
	Size  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Size: 0}

}

func (l *LinkedList) Insert(data any) {
	newNode := Node{value: data, next: nil}
	// if list size is zero, newNode is the first node
	if l.Size == 0 {
		l.First, l.Last = &newNode, &newNode
		l.Size = 1
		return
	}

	// if list size != 0, add newNode to last node pointer
	l.Last.next = &newNode
	l.Last = &newNode
	l.Size++
}
