package bst

import "testing"

func TestCreateBST(t *testing.T) {

	bst := &BST{}
	bst.Insert(20)
	bst.Insert(4)
	bst.Insert(5)
	bst.Insert(6)
	bst.Insert(9)
	bst.Insert(100)
	bst.Insert(3)

	bst.Delete(20)
	bst.Delete(3)
	sliced := bst.WalkInOrder()

}

func TestMinNode(t *testing.T) {

	bst := &BST{}
	bst.Insert(20)
	bst.Insert(4)
	bst.Insert(5)
	bst.Insert(6)
	bst.Insert(9)
	bst.Insert(100)
	minNode := bst.Insert(3)
	result := bst.MinNode(bst.Root)

	if result != minNode {
		t.Errorf("MinNode %d did not match expected %d", result, minNode)
	}
}

func TestMaxNode(t *testing.T) {

	bst := &BST{}
	bst.Insert(20)
	bst.Insert(4)
	bst.Insert(5)
	bst.Insert(6)
	bst.Insert(9)
	maxNode := bst.Insert(100)
	bst.Insert(3)
	result := bst.MaxNode(bst.Root)

	if result != maxNode {
		t.Errorf("MaxNode %d did not match expected %d", result, maxNode)
	}
}
