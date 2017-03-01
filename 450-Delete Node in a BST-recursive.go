package leetcode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil { // base case
		return nil // key not found
	}

	if key > root.Val { // >
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val { // <
		root.Left = deleteNode(root.Left, key)
	} else { // = found key
		if root.Left == nil { // 0 or 1 child
			return root.Right
		} else if root.Right == nil { // 1 child
			return root.Left
		}
		min := findMin(root.Right) // 2 children
		root.Val = min.Val
		root.Right = deleteNode(root.Right, root.Val)
	}

	return root
}

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}
