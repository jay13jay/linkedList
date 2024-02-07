package linkedList

import (
	"fmt"
)

type Node struct {
	data int   // int value
	next *Node // pointer to the next node in the list
}

type LinkedList struct {
	head  *Node // Pointer to the first node of the list, or null if list is empty
	count int   // len of the list
}

func (l *LinkedList) Push(value int) {
	// Push value to the front of the list
	newNode := &Node{data: value, next: l.head}
	// newNode.next = l.head
	l.head = newNode
	l.count++
	return
}

func (l *LinkedList) Append(value int) {
	// Append value to the end of the list
	newNode := &Node{data: value, next: nil} // next is null as this will be the last node
	last := l.head
	if l.head == nil {
		l.head = newNode
		return
	}

	for last.next != nil {
		last = last.next // traverse to the end of the list
	}

	last.next = newNode
	l.count++
	return

}

func (l *LinkedList) InsertAfter(node *Node, value int) {
	// Insert value after the provided Node
	newNode := &Node{data: value, next: node.next}
	node.next = newNode

}

func (l *LinkedList) PrintList() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.next
	}
	fmt.Println("nil")
}
func (l *LinkedList) Len() int {
	// fmt.Println(l.count)
	return l.count
}
