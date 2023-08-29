package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func AddNode(head *Node, value int) *Node {
	newNode := &Node{value: value, next: head}
	fmt.Println("Node added:", value)
	return newNode
}

func DeleteNode(head *Node, value int) *Node {
	if head == nil {
		fmt.Println("Linked list is empty")
		return nil
	}

	if head.value == value {
		fmt.Println("Node deleted:", value)
		return head.next
	}

	current := head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next == nil {
		fmt.Println("Node not found:", value)
		return head
	}

	current.next = current.next.next
	fmt.Println("Node deleted:", value)
	return head
}

func Traverse(head *Node) {
	if head == nil {
		return
	}
	Traverse(head.next)
	fmt.Print(head.value, " ")
}
