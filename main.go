package main

import (
	"fmt"
)

type Node struct {
	name string
	next *Node
}

type List struct {
	head *Node
}

func (l *List) AddItem(name string) {
	newNode := &Node{name: name}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	if current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func main() {
	linkedList := &List{}
	linkedList.AddItem("Jose")
	linkedList.AddItem("Henrique")
	linkedList.AddItem("Mariana")

	fmt.Println("Lista encadeada:")
	curr := linkedList.head
	for curr != nil {
		//Operação de busca O(n)
		if curr.name == "Henrique" {
			fmt.Printf("Nome %s presente na lista\n", curr.name)
		}
		fmt.Printf("%s \n", curr.name)
		curr = curr.next
	}
}