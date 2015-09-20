package bst

import (
	"log"
	"testing"
)

func TestCreateBST(t *testing.T) {

	bst := &BST{}
	bst.Insert(20)
	bst.Insert(4)
	bst.Insert(5)
	bst.Insert(6)
	bst.Insert(9)
	bst.Insert(100)
	bst.Insert(3)

	log.Println(bst.WalkInOrder())
}
