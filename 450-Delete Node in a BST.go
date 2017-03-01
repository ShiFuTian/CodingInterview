package leetcode

func deleteNode2(root *TreeNode, key int) *TreeNode {
	// find key
	kp := root           // key point to find
	var pre [2]*TreeNode // parent
	for kp != nil {      // find the key
		if key > kp.Val {
			pre[0], pre[1] = nil, kp // parent propagate
			kp = kp.Right
		} else if key < kp.Val {
			pre[0], pre[1] = kp, nil
			kp = kp.Left
		} else {
			break // found
		}
	}
	if kp == nil { // not found
		return root
	}

	// delete key
	dp := kp
	if kp.Left != nil && kp.Right != nil { // kp have two children
		pre[0], pre[1] = nil, dp
		dp = dp.Right
		for dp.Left != nil { // find min of right child
			pre[0], pre[1] = dp, nil
			dp = dp.Left
		}
		kp.Val = dp.Val // delete key
	}

	if dp.Left == nil {
		// dp.Praent.xx = dp.Right
		if pre[1] != nil {
			pre[1].Right = dp.Right
		} else if pre[0] != nil {
			pre[0].Left = dp.Right
		} else { // if dp == root
			root = dp.Right
		}
	} else if dp.Right == nil { // dp have 0 or 1 child
		// dp.Praent.xx = dp.Left
		if pre[1] != nil {
			pre[1].Right = dp.Left
		} else if pre[0] != nil {
			pre[0].Left = dp.Left
		} else { // if dp == root
			root = dp.Left
		}
	}

	return root
}
