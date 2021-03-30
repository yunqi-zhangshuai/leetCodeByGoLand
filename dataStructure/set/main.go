package main

import (
	"fmt"
	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

func main() {
	list := sll.New()
	list.Append(1, 2, 3, 4)
	cur := list.Iterator()
	for cur.Next() {
		fmt.Println(cur.Value())
	}
}
