package piscine

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                 string
// }

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}
// 	if data < root.Data {
// 		if root.Left == nil {
// 			root.Left = &TreeNode{Data: data, Parent: root}
// 		} else {
// 			BTreeInsertData(root.Left, data)
// 		}
// 	} else {
// 		if root.Right == nil {
// 			root.Right = &TreeNode{Data: data, Parent: root}
// 		} else {
// 			BTreeInsertData(root.Right, data)
// 		}
// 	}
// 	return root
// }

//BTreeSearchItem searches for a node with matching data
// func BTreeSearchItem(root *TreeNode, data string) *TreeNode {
// 	if root == nil || root.Data == data {
// 		return root
// 	}
// 	if data < root.Data {
// 		return BTreeSearchItem(root.Left, data)
// 	}
// 	return BTreeSearchItem(root.Right, data)
// }

// func BTreeIsBinary(root *TreeNode) bool {
// 	return isBinaryHelper(root, "", "")
// }

func isBinaryHelper(node *TreeNode, minVal, maxVal string) bool {
	if node == nil {
		return true
	}

	if minVal != "" && node.Data <= minVal {
		return false
	}
	if maxVal != "" && node.Data >= maxVal {
		return false
	}

	return isBinaryHelper(node.Left, minVal, node.Data) &&
		isBinaryHelper(node.Right, node.Data, maxVal)
}

// func BTreeMin(root *TreeNode) *TreeNode {
// 	current := root
// 	for current != nil && current.Left != nil {
// 		current = current.Left
// 	}
// 	return current
// }

func replaceNodeInParent(root, oldNode, newNode *TreeNode) *TreeNode {
	if oldNode.Parent == nil {
		root = newNode
	} else if oldNode == oldNode.Parent.Left {
		oldNode.Parent.Left = newNode
	} else {
		oldNode.Parent.Right = newNode
	}

	if newNode != nil {
		newNode.Parent = oldNode.Parent
	}

	return root
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	if node.Left == nil {
		root = replaceNodeInParent(root, node, node.Right)
	} else if node.Right == nil {
		root = replaceNodeInParent(root, node, node.Left)
	} else {
		successor := BTreeMin(node.Right)

		if successor.Parent != node {
			root = replaceNodeInParent(root, successor, successor.Right)
			successor.Right = node.Right
			if successor.Right != nil {
				successor.Right.Parent = successor
			}
		}

		root = replaceNodeInParent(root, node, successor)
		successor.Left = node.Left
		if successor.Left != nil {
			successor.Left.Parent = successor
		}
	}

	return root
}
