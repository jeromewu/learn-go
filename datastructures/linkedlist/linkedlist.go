package linkedlist

import (
	"fmt"
)

// Node is an element in the linked list
type Node struct {
	value interface{}
	next  *Node
}

// LinkedList is a list of Nodes
type LinkedList struct {
	root *Node
}

// NewNode with assigning next to nil
func NewNode(v interface{}) Node {
	return Node{value: v, next: nil}
}

// NewLinkedList with assigning root to nil
func NewLinkedList() LinkedList {
	return LinkedList{root: nil}
}

// PrintNodes prints all nodes in list
func (list *LinkedList) PrintNodes() {
	it := list.root
	fmt.Printf("list: ")
	for it != nil {
		fmt.Printf("%d ", it.value.(int))
		it = it.next
	}
	fmt.Printf("\n")
}

// AppendNode appends a node to LinkedList
func (list *LinkedList) AppendNode(node *Node) {
	if list.root == nil {
		list.root = node
		return
	}
	it := list.root
	for it.next != nil {
		it = it.next
	}
	it.next = node
}

// InsertNode inserts a node after index
func (list *LinkedList) InsertNode(index int, node *Node) {
	if list.root == nil {
		list.root = node
		return
	}
	it := list.root
	for i := 0; i < index && it.next != nil; i++ {
		it = it.next
	}
	it.next = node
}

// DeleteNode deletes a node
func (list *LinkedList) DeleteNode(index int) {
	if list.root == nil {
		return
	}
	it := list.root
	prev := it
	for i := 0; i < index && it.next != nil; i++ {
		prev = it
		it = it.next
	}

	if list.root == it {
		list.root = it.next
		return
	}
	prev.next = it.next
}

// GetNode returns node at specific index
func (list *LinkedList) GetNode(index int) *Node {
	it := list.root
	for i := 0; i < index; i++ {
		it = it.next
	}
	return it
}
