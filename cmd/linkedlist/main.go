package main

import (
	"github.com/ravitezu/al-go-rithms/linkedlist/singly_linked_list"
)

func main() {
	sl := singly_linked_list.NewList()
	for i := range "123456789"  {
		sl.Append(i)
	}
	sl.Append("a")
	sl.Append("b")
	sl.Display()
	sl.Reverse()
	sl.Display()
	sl.Append(-1)
	sl.Display()
	sl.Reverse()
	sl.Display()
	sl.Prepend(-2)
	sl.Display()
	n1 := singly_linked_list.NewNode("c")
	sl.Tail.Next = n1
	sl.Tail = n1
	sl.Display()
	n2 := singly_linked_list.NewNode("e")
	sl.Tail.Next = n2
	sl.Tail = n2
	sl.Display()
	err := sl.Insert(n1, "d")
	if err != nil {
		panic(err)
	}
	sl.Display()
}