package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	minNode := root
	if leftMin := BTreeMin(root.Left); leftMin != nil {
		if leftMin.Data < minNode.Data {
			minNode = leftMin
		}
	}
	if rightMin := BTreeMin(root.Left); rightMin != nil {
		if rightMin.Data < minNode.Data {
			minNode = rightMin
		}
	}
	return minNode
}
