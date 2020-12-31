package main

import "fmt"

func rotateRight(head *ListNode, k int) *ListNode {
	// 先计算len，然后遍历获取点，然后截取

	if head == nil || k == 0 {
		return head
	}
	var len int = 0
	p := head
	for p != nil {
		len += 1
		p = p.Next
	}
	fmt.Println("len=", len)
	var n = len - k%len
	if n == 0 || n == len {
		return head
	}
	fmt.Println("n=", n)

	p1 := head
	for head != nil && n-1 > 0 {
		n -= 1
		head = head.Next
	}
	p2 := head.Next
	head.Next = nil
	p3 := p2
	for p2.Next != nil {
		p2 = p2.Next
	}
	p2.Next = p1
	return p3
}
