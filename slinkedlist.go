package main

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (ll *LinkedList) Size() int {
	return ll.length
}

func (ll *LinkedList) isEmpty() bool {
	return ll.length == 0
}
