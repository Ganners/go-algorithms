package bst

import "container/list"

type BST struct {
	Root *Node
	Size int
}

type Node struct {
	Data   int
	Parent *Node
	Left   *Node
	Right  *Node
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

		// Needs to work for duplicate nodes and go deeper
		if searchNode.Data == needle {
			if searchNode.Left != nil && searchNode.Left.Data == needle {
				searchNode = searchNode.Left
				continue
			} else if searchNode.Right != nil && searchNode.Right.Data == needle {
				searchNode = searchNode.Right
				continue
			}
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
func (bst *BST) Insert(data int) *Node {

	node := &Node{Data: data}

	if bst.Root == nil {
		bst.Root = node
		return nil
	}

	// Find best place to insert and do so
	return bst.insertAt(node, bst.Search(data))
}

// Inserts into the tree at a given node
func (bst *BST) insertAt(node, atNode *Node) *Node {

	if atNode.Data < node.Data {
		atNode.Right = node
	} else {
		atNode.Left = node
	}
	node.Parent = atNode

	return node
}

func (bst *BST) Delete(data int) {

	bst.deleteNode(bst.Root, data)
}

func (bst *BST) deleteNode(node *Node, data int) *Node {

	if node == nil {
		return node
	}

	node = bst.Search(data)
	if node.Data != data {
		// Cannot be found
		return node
	}

	if node.Left == nil || node.Right == nil {
		if node.Left == nil && node.Right == nil {
			bst.linkParent(node.Parent, node, nil)
		} else if node.Left != nil {
			bst.linkParent(node.Parent, node, node.Left)
		} else if node.Right != nil {
			bst.linkParent(node.Parent, node, node.Right)
		}
	} else {
		minNode := bst.MinNode(node.Right)
		node.Data = minNode.Data
		node.Right = bst.deleteNode(node.Right, minNode.Data)
	}

	return node
}

func (bst *BST) linkParent(parent, from, to *Node) {

	if parent.Left == from {
		parent.Left = to
	} else if parent.Right == from {
		parent.Right = to
	}
}

func (bst *BST) MinNode(node *Node) *Node {
	minNode := node
	for {
		if minNode.Left == nil {
			break
		}
		minNode = minNode.Left
	}
	return minNode
}

func (bst *BST) MaxNode(node *Node) *Node {
	maxNode := node
	for {
		if maxNode.Right == nil {
			break
		}
		maxNode = maxNode.Right
	}
	return maxNode
}
