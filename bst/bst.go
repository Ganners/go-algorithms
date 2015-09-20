package bst

import "container/list"

type BST struct {
	Root *Node
	Size int
}

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// Iterative in order tree traversal
func (bst *BST) WalkInOrder() []int {

	// Push = stack.PushBack()
	// Pop = stack.Remove(stack.Back())
	// Peek = stack.Back()
	returnList := make([]int, 0)
	stack := list.New()
	node := bst.Root

	for stack.Len() > 0 || node != nil {

		// For each non nil node, add it to a stack
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else {

			// For each node on that stack, pop it, look at the current
			// node then look at it's right node (which will add to the
			// stack)
			node = stack.Remove(stack.Back()).(*Node)
			returnList = append(returnList, node.Data)
			node = node.Right
		}
	}

	return returnList
}

// Searches the tree for a given node
func (bst *BST) Search(needle int) *Node {

	searchNode := bst.Root

	for searchNode != nil {

		if searchNode.Data == needle {

			return searchNode
		}

		if searchNode.Data > needle {
			if searchNode.Left == nil {
				return searchNode
			}
			searchNode = searchNode.Left
		} else {
			if searchNode.Right == nil {
				return searchNode
			}
			searchNode = searchNode.Right
		}
	}
	return searchNode
}

// Inserts the node
func (bst *BST) Insert(data int) {

	node := &Node{Data: data}

	if bst.Root == nil {
		bst.Root = node
		return
	}

	// Find best place to insert and do so
	bst.insertAt(node, bst.Search(data))
}

// Inserts into the tree at a given node
func (bst *BST) insertAt(node, atNode *Node) {

	if atNode.Data < node.Data {
		atNode.Right = node
	} else {
		atNode.Left = node
	}
}
