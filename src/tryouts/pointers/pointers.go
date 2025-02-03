package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	dummy := new(ListNode)
	tail := dummy

	for _, val := range arr {
		node := &ListNode{val: val}
		tail.next = node
		tail = tail.next
	}

	tail = dummy.next

	for tail != nil {
		fmt.Println(tail.val)
		tail = tail.next
	}

}
