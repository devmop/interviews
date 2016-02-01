package bst

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var valid []*Tree

func init() {
	valid = append(valid, nil)
	valid = append(valid, Leaf(1))
	valid = append(valid, Node(2, Leaf(1), Leaf(3)))
}

func TestValidBST(t *testing.T) {

	for _, tree := range valid {
		if !tree.IsBST() {
			t.Error(spew.Sprintf("%v should be a bst", tree))
		}
	}
}

func TestInvalidBST(t *testing.T) {
	tree := Node(1, Leaf(2), Leaf(3))

	if tree.IsBST() {
		t.Error(spew.Sprintf("%v should not be a bst", tree))
	}
}

func Node(val uint, left, right *Tree) *Tree {
	return &Tree{value: val, left: left, right: right}
}

func Leaf(val uint) *Tree {
	return &Tree{value: val}
}
