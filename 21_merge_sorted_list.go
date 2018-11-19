package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//// 优先判断边界条件
	//if l1 == nil && l2 == nil {
	//	return nil
	//}
	//head := &ListNode{}
	//l := head
	//
	//// 复制参数，避免移动过程中对传入的外界变量造成影响
	//tl1, tl2 := l1, l2
	//for tl1 != nil && tl2 != nil {
	//	if tl1.Val > tl2.Val {
	//		l.Next = &ListNode{tl2.Val, nil}
	//		tl2 = tl2.Next
	//	} else {
	//		l.Next = &ListNode{tl1.Val, nil}
	//		tl1 = tl1.Next
	//	}
	//	l = l.Next
	//
	//}
	//
	//for tl1 != nil {
	//	// 或 l.Next = ll1，但这样会导致返回结果中包含外界变量l1的地址
	//	l.Next = &ListNode{tl1.Val, nil}
	//	l = l.Next
	//	tl1 = tl1.Next
	//}
	//for tl2 != nil {
	//	l.Next = &ListNode{tl2.Val, nil}
	//	l = l.Next
	//	tl2 = tl2.Next
	//}

	// another solution
	// try to do not copy value
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, current *ListNode
	if l1.Val > l2.Val {
		current = l2
		l2 = l2.Next
	} else {
		current = l1
		l1 = l1.Next
	}
	head = current

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			current.Next = l2
			current = l2
			l2 = l2.Next
		} else {
			current.Next = l1
			current = l1
			l1 = l1.Next
		}
	}

	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return head
}

func main() {
	l1Head := &ListNode{1, nil}
	l1 := l1Head
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{3, nil}

	l2Head := &ListNode{1, nil}
	l2 := l2Head
	l2.Next = &ListNode{2, nil}
	l2.Next.Next = &ListNode{3, nil}

	mergedList := mergeTwoLists(l1Head, l2Head)
	for mergedList != nil {
		fmt.Printf("%d ", mergedList.Val)
		mergedList = mergedList.Next
	}
}
