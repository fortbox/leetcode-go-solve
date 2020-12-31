/*
 * Copyright (c) 2020
 * @Author: xiaoweixiang
 */

package main

func deleteHead(head *ListNode) (bool, *ListNode) {
	if head == nil {
		return false, head
	}
	if head.Next == nil {
		return false, head
	}
	canDelete := false
	i := 0
	for head.Next != nil && head.Val == head.Next.Val {
		//fmt.Println("i:",i)
		i += 1
		head = head.Next
		canDelete = true
	}
	//fmt.Println("head:", head.Val)
	//fmt.Println("canDelete:", canDelete)
	if canDelete {
		head = head.Next
		return true, head
	}
	return false, head
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 需要记录前一个结点、当前结点、下一个结点
	// 先处理首节点
	for res, h := deleteHead(head); res == true; res, h = deleteHead(head) {
		head = h
	}
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	canDelete := false
	pre := head
	cur := pre.Next
	next := cur.Next
	for next != nil {
		//fmt.Println("pre:", pre.Val)
		//fmt.Println("cur:", cur.Val)
		//fmt.Println("next:", next.Val)
		if cur.Val == next.Val {
			cur.Next = next.Next
			next = next.Next
			canDelete = true
		} else {
			if canDelete {
				pre.Next = next
				cur = pre.Next
				next = cur.Next
				canDelete = false
			} else {
				pre = pre.Next
				cur = cur.Next
				next = next.Next
			}
		}
	}
	if canDelete {
		pre.Next = next
		cur = pre.Next
	}
	return head
}
