package main

import (
	"fmt"

	"github.com/utkarsh-pro/ugstl/pkg/LinkedList/singly"
)

func main() {
	ll := singly.New[int]()

	ll.InsertAtTail(1)
	ll.InsertAtTail(100)
	ll.InsertAtTail(1000)
	fmt.Println(ll.Len())

	iter := ll.Iter()
	for {
		val, ok := iter.Next()
		if !ok {
			break
		}

		fmt.Println(val.GetValue())
	}
}
