package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := new(ByteStack)
	s2 := new(ByteStack)
	for p := l1; p != nil; p = p.Next {
		s1.Push(byte(p.Val))
	}
	for p := l2; p != nil; p = p.Next {
		s2.Push(byte(p.Val))
	}

	var sum, val byte
	var head *ListNode // nice code to pushFront

	for !s1.Empty() || !s2.Empty() || sum != 0 { // sum need to be emptyed too
		if !s1.Empty() {
			sum += s1.Pop()
		}
		if !s2.Empty() {
			sum += s2.Pop()
		}

		sum, val = sum/10, sum%10

		node := &ListNode{int(val), head} // nice code to pushFront
		head = node                       // nice code to pushFront
	}

	return head
}
