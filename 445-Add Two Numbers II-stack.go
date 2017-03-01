package leetcode

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	newStack := func(l *ListNode) []int {
		var stack []int
		for p := l; p != nil; p = p.Next {
			stack = append(stack, p.Val)
		}
		return stack
	}
	s1, s2 := newStack(l1), newStack(l2)

	// popStack := func(stack []int) int {
	// 	slen := len(stack)
	// 	ret := stack[slen-1]
	// 	stack = stack[:slen-1]
	// 	return ret
	// }

	var i, j = len(s1) - 1, len(s2) - 1
	var sum, carry = 0, 0
	var sumStack = make([]int, 0, i)
	for i >= 0 { // s1 not empty
		if j >= 0 { // s2 not empty
			sum, carry = s1[i]+s2[j]+carry, 0
			if sum > 9 {
				sum, carry = sum-10, 1
			}
			sumStack = append(sumStack, sum)
			j--
		} else if carry > 0 { // add carry
			sum, carry = s1[i]+carry, 0
			if sum > 9 {
				sum, carry = sum-10, 1
			}
			sumStack = append(sumStack, sum)
		} else { // empty s1
			sumStack = append(sumStack, s1[i])
		}
		i--
	}

	for j >= 0 { // s2 not empty
		if carry > 0 { // add carry
			sum, carry = s2[j]+carry, 0
			if sum > 9 {
				sum, carry = sum-10, 1
			}
			sumStack = append(sumStack, sum)
		} else { // empty s2
			sumStack = append(sumStack, s2[j])
		}
		j--
	}

	if carry > 0 { // add carry
		sumStack = append(sumStack, carry)
	}

	var ret = &ListNode{Next: nil} //head, numpy node
	pre := ret
	for i = len(sumStack) - 1; i >= 0; i-- {
		pre.Next = &ListNode{sumStack[i], nil}
		pre = pre.Next
	}

	return ret.Next // drop head
}
