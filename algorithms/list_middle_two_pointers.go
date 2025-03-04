package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func middle(list *ListNode) int {
	slow, fast := list, list

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.value
}

func main() {
	listEven := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	fmt.Println(middle(listEven))

	listOdd := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}
	fmt.Println(middle(listOdd))
}
