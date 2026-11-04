package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	maxNode := root
	if leftMax := BTreeMax(root.Left); leftMax != nil {
		if leftMax.Data > maxNode.Data {
			maxNode = leftMax
		}
	}
	if rightMax := BTreeMax(root.Right); rightMax != nil {
		if rightMax.Data > maxNode.Data {
			maxNode = rightMax
		}
	}

	return maxNode
}
