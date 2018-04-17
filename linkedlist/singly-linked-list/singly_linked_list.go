// This package implements a Singly Linked List.

// ToDo: Add Unit Tests.

package linkedlist

import (
	"fmt"
	"errors"
)

// List represents a Singly Linked List.
type List struct {
	Head *Node
	Tail *Node
}

// NewList returns a new List.
func NewList() *List {
	return &List{}
}

// Prepend accepts new data(interface{}) and creates a Node with it
// and adds the Node before the HEAD Node.
func (l *List) Prepend(d interface{}) {
	// Create a new Node with the data received.
	newNode := &Node{
		Data: d,
		Next: nil,
	}

	// Make the current HEAD as Next to the new Node.
	newNode.Next = l.Head

	// Update the HEAD to point the new node.
	l.Head = newNode
}

// Insert accepts a Node and new data. It inserts a new Node(with supplied data d)
// after to the passed Node. Also, it returns an error if any or nil.
func (l *List) Insert(after *Node, d interface{}) error {
	if after == l.Head || after == l.Tail {
		return errors.New("use the Prepend/Append methods to Insert new Nodes at Start/End")
	}
	newNode := &Node{
		Data: d,
		Next: nil,
	}

	nextNode := after.Next
	after.Next = newNode
	newNode.Next = nextNode
	return nil
}

// Append accepts new data(interface{}) and creates a Node with it
// and appends the Node to the end of the List.
func (l *List) Append(d interface{}) {
	// Create a new Node with the data received.
	newNode := &Node{
		Data: d,
		Next: nil,
	}

	if l.Head == nil {
		// This is the first Node in the List.
		l.Head = newNode
		l.Tail = newNode
		return
	}

	// Update the current tail next to point to the new node.
	l.Tail.Next = newNode

	// Update the tail to the new node added.
	l.Tail = newNode
}

// Display prints the Linked List.
func (l *List) Display() {
	// Assume the head is the tail of the List.
	curr := l.Head
	for {
		if curr.Next == nil {
			fmt.Printf("%#v\n", curr.Data)
			break
		}
		fmt.Printf("%#v |=> ", curr.Data)
		curr = curr.Next
	}
}

// Reverse reverses the Linked List.
func (l *List) Reverse() {
	var prev *Node
	tempTail := l.Head
	curNode := l.Head
	for {
		if curNode.Next == nil {
			curNode.Next = prev
			break
		} else {
			nextNode := curNode.Next
			curNode.Next = prev
			prev = curNode
			curNode = nextNode
		}
	}

	l.Head = curNode
	l.Tail = tempTail
}

// Node represents a Node in the Singly Linked List.
type Node struct {
	Data interface{}
	Next *Node
}

// NewNode accepts new data(interface{}) and creates a Node with it and
// returns a pointer to it.
func NewNode(d interface{}) *Node {
	return &Node{
		Data: d,
		Next: nil,
	}
}



