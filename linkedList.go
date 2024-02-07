package linkedList

import (
	"fmt"
)

type Node struct {
	Data int   // int value
	Next *Node // pointer to the next node in the list
}

type LinkedList struct {
	Head  *Node // Pointer to the first node of the list, or null if list is empty
	Count int   // len of the list
}

func (l *LinkedList) Push(value int) {
	// Push value to the front of the list
	newNode := &Node{Data: value, Next: l.Head}
	// newNode.next = l.head
	l.Head = newNode
	l.Count++
}

func (l *LinkedList) Append(value int) {
	// Append value to the end of the list
	newNode := &Node{Data: value, Next: nil} // next is null as this will be the last node
	last := l.Head
	if l.Head == nil {
		l.Head = newNode
		return
	}

	for last.Next != nil {
		last = last.Next // traverse to the end of the list
	}

	last.Next = newNode
	l.Count++

}

func (l *LinkedList) InsertAfter(node *Node, value int) {
	// Insert value after the provided Node
	newNode := &Node{Data: value, Next: node.Next}
	node.Next = newNode

}

func (l *LinkedList) PrintList() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Data)
		curr = curr.Next
	}
	fmt.Println("nil")
}
