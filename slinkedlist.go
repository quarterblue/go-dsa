package main

import (
	"errors"
	"fmt"
)

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

// TODO
func (ll *LinkedList) ValueAt(idx int) int {
	return 0
}

func (ll *LinkedList) PeakFront() (d int, err error) {
	if ll.isEmpty() {
		return 0, errors.New("List is empty.")
	}
	return ll.head.data, nil
}

func (ll *LinkedList) PeakBack() (d int, err error) {
	if ll.isEmpty() {
		return 0, errors.New("List is empty.")
	}
	return ll.tail.data, nil
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

func (ll *LinkedList) PopFront() (d int, err error) {
	if ll.isEmpty() {
		return 0, errors.New("List is empty.")
	}

	if ll.head.next == nil {
		data := ll.head.data
		ll.head = nil
		ll.tail = nil
		ll.length--
		return data, nil
	} else {
		data := ll.head.data
		ll.head = ll.head.next
		ll.length--
		return data, nil
	}
}

// TODO
func (ll *LinkedList) PopBack() (d int, err error) {
	if ll.isEmpty() {
		return 0, errors.New("List is empty.")
	}

	return 0, nil
}

// TODO
func (ll *LinkedList) Delete(idx int) (d int, err error) {
	return 0, nil
}

// TODO
func (ll *LinkedList) ValueFromEnd(n int) (d int, err error) {
	return 0, nil
}

// TODO
func (ll *LinkedList) Reverse() error {
	return nil
}

// TODO
func (ll *LinkedList) RemoveValue(data int) (d int, err error) {
	return 0, nil
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
	newLL.PopFront()
	newLL.PrintList()
	newLL.PopFront()
	newLL.PrintList()
	newLL.PopFront()
	newLL.PrintList()
	newLL.PopFront()
	newLL.PrintList()
	newLL.PopFront()
	newLL.PrintList()
	newLL.PopFront()
	newLL.PrintList()
	newLL.PopFront()
	newLL.PrintList()
}
