package bst

type Tree struct {
	value       uint
	left, right *Tree
}

func (t *Tree) IsBST() bool {
	return t.isBST(0, ^uint(0))
}

func (t *Tree) isBST(min, max uint) bool {
	if t == nil {
		return true
	}

	if t.value < min || t.value > max {
		return false
	}

	return t.left.isBST(min, t.value) && t.right.isBST(t.value, max)
}
