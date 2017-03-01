package leetcode

// pass 1558/1560
// FAIL at
// l1 = [2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,
// 			2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9]
// l2 = [5,6,4,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,
// 			2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9,9,9,9]
// w = [8,0,7,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,
// 			8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,7,2,4,3,8]
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	newInt := func(l *ListNode) (ret int) {
		for p := l; p != nil; p = p.Next {
			ret = ret*10 + p.Val
		}
		return ret
	}
	i1, i2 := newInt(l1), newInt(l2)

	sum, a := i1+i2, 0

	var p1, p2 *ListNode = nil, nil
	for sum >= 10 {
		sum, a = sum/10, sum%10
		p1 = &ListNode{a, p2} // backward growth
		p2 = p1
	}
	p1 = &ListNode{sum, p2} // backward growth
	p2 = p1

	return p1
}
