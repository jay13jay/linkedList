package jhaxgoll

import "fmt"

type Node struct {
	data int
	next *Node
}

type LL struct {
	head *Node
}

func (l *LL) Push(value int) {
	newNode := &Node{data: value}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) PrintList() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.next
	}
	fmt.Println("nil")
}
