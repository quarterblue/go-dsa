package main

import "fmt"

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

func (ll *LinkedList) PushBack(data int) (d int, err error) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		ll.length++
		return data, nil
	}

	curr := ll.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	ll.tail = newNode
	ll.length++

	return data, nil
}

func (ll *LinkedList) PushFront(data int) (d int, err error) {
	if ll.head == nil {
		ll.PushBack(data)
		return data, nil
	}

	newNode := &Node{
		data: data,
		next: nil,
	}

	newNode.next = ll.head
	ll.head = newNode
	ll.length++

	return data, nil
}

// Prints the list with styling
// 8 -> 17 -> 2 -> 1 -> 87
func (ll *LinkedList) PrintList() {
	if ll.head == nil {
		fmt.Println("Empty List")
		return
	}

	curr := ll.head

	for curr != nil {
		fmt.Printf("%d", curr.data)
		fmt.Printf(" -> ")
		curr = curr.next
	}
	fmt.Printf("NULL\n")

}

func main() {
	newLL := NewLinkedList()
	// 7 -> NULL
	newLL.PushBack(7)
	// 7 -> 2 -> NULL
	newLL.PushBack(2)
	// 7 -> 2 -> 12 -> NULL
	newLL.PushBack(12)
	// 7 -> 2 -> 12 -> 32 -> NULL
	newLL.PushBack(32)
	// 7 -> 2 -> 12 -> 32 -> 8 -> NULL
	newLL.PushBack(8)
	// 7 -> 2 -> 12 -> 32 -> 8 -> 99 -> NULL
	newLL.PushFront(99)
	// 7 -> 2 -> 12 -> 32 -> 8 -> 99 -> 97 -> NULL
	newLL.PushFront(97)
	// 7 -> 2 -> 12 -> 32 -> 8 -> 99 -> 97 -> 17 -> NULL
	newLL.PushBack(17)
	newLL.PrintList()
}
