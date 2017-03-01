package leetcode

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {

	reverse := func(l *ListNode) *ListNode { // reverse list
		var p1, p2, p3 *ListNode = l, nil, nil
		for p1 != nil {
			p2, p1 = p1, p1.Next
			p2.Next, p3 = p3, p2
		}
		return p2
	}
	l1 = reverse(l1)
	l2 = reverse(l2)

	carry := 0

	var pre, p1, p2 *ListNode = nil, l1, l2
	for p1 != nil && p2 != nil {
		p1.Val, carry = p1.Val+p2.Val+carry, 0
		if p1.Val > 9 {
			// p1.Val, carry = p1.Val%10, p1.Val/10
			p1.Val, carry = p1.Val-10, 1
		}
		pre, p1, p2 = p1, p1.Next, p2.Next
	}

	if p1 == nil && p2 != nil { // check non-nil
		pre.Next, p1 = p2, p2
	}

	for p1 != nil { // add carry
		if carry == 0 {
			break
		}
		p1.Val, carry = p1.Val+carry, 0
		if p1.Val > 9 {
			// p1.Val, carry = p1.Val%10, p1.Val/10
			p1.Val, carry = p1.Val-10, 1
		}
		pre, p1 = p1, p1.Next
	}

	if carry != 0 {
		pre.Next = &ListNode{carry, nil}
	}

	l1 = reverse(l1)

	return l1
}
