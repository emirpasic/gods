package twofourtree

import (
	"github.com/emirpasic/gods/trees"
	"github.com/emirpasic/gods/utils"
)

func assertTreeImplementation() {
	var _ trees.Tree = new(Tree)
}

// Tree holds elements of the two-four tree.
type Tree struct {
	root       *Node
	Comparator utils.Comparator
}

// Node is a single element within the tree
type Node struct {
	numItems int
	parent   *Node
	children [4]*Node
	nodeData [3]*NodeData
}

// NodeData contains Node's data
type NodeData struct {
	data int
}

// NewWithIntComparator instantiates a two-four tree with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator() *Tree {
	return &Tree{root: &Node{}, Comparator: utils.IntComparator}
}

func (t *Tree) insert(value int) {
	curNode := t.root
	nodeData := &NodeData{data: value}
	for {
		if curNode.isFull() {
			t.split(curNode)
			curNode = curNode.parent
			curNode = getNextChild(curNode, value)
		} else if curNode.isLeaf() {
			break
		} else {
			curNode = getNextChild(curNode, value)
		}
	}
	curNode.insertItem(nodeData)
}

func (t *Tree) split(node *Node) {
	itemC := node.removeItem()
	itemB := node.removeItem()
	child2 := node.disconnectChild(2)
	child3 := node.disconnectChild(3)
	var parent *Node
	newRight := &Node{}
	if node == t.root {
		t.root = &Node{}
		parent = t.root
		t.root.connectChild(0, node)
	} else {
		parent = node.parent
	}
	itemIndex := parent.insertItem(itemB)
	n := parent.numItems
	for i := n - 1; i > itemIndex; i-- {
		temp := parent.disconnectChild(i)
		parent.connectChild(i+1, temp)
	}
	parent.connectChild(itemIndex+1, newRight)
	newRight.insertItem(itemC)
	newRight.connectChild(0, child2)
	newRight.connectChild(1, child3)
}

func getNextChild(node *Node, value int) *Node {
	numItems := node.numItems
	for i := 0; i < numItems; i++ {
		if value < node.nodeData[i].data {
			return node.children[i]
		}
	}
	return node.children[numItems]
}

func (t *Tree) find(value int) *Node {
	return findValue(t.root, value)
}

func findValue(node *Node, value int) *Node {
	numItems := node.numItems
	for i := 0; i < numItems; i++ {
		if value == node.nodeData[i].data {
			return node
		} else if value < node.nodeData[i].data && !node.isLeaf() {
			return findValue(node.children[i], value)
		} else if value > node.nodeData[i].data && !node.isLeaf() {
			return findValue(node.children[i+1], value)
		}
	}
	return nil
}

func (t *Tree) delete(node *Node, value int) *Node {
	if node.isLeaf() {
		if node.numItems > 1 {
			node.deleteNodeValue(value)
			return node
		}
		return t.deleteLeaf(node, value)
	}
	n := getNextChild(node, value)
	c := getInOrderNode(n)
	d := c.nodeData[0]
	k := d.data
	t.delete(c, k)
	found := t.find(value)
	for i := 0; i < found.numItems; i++ {
		if found.nodeData[i].data == value {
			found.nodeData[i].data = k
		}
	}
	return found
}

func (t *Tree) deleteLeaf(node *Node, value int) *Node {
	siblingSide := "l"
	p := node.parent
	sibling := node.getSibling(value)
	if sibling == nil {
		siblingSide = "r"
		sibling = p.children[1]
	}
	if sibling.numItems == 1 {
		for i := 0; i <= p.numItems; i++ {
			if p.children[i] == sibling && siblingSide == "l" {
				node.nodeData[node.numItems-1] = nil
				node.numItems--
				d := p.nodeData[i]
				sibling.insertItem(d)
				p.disconnectChild(i + 1)
				for j := i; j < p.numItems; j++ {
					if j+1 < p.numItems {
						p.nodeData[j] = p.nodeData[j+1]
						if j+2 <= p.numItems {
							p.connectChild(j+1, p.disconnectChild(j+2))
						}
					}
				}
				p.nodeData[p.numItems-1] = nil
				p.numItems--
				if p.numItems == 0 {
					if p != t.root {
						p = t.balanceTree(p)
					} else {
						t.root = sibling
					}
				}
				return p
			} else if p.children[i] == sibling && siblingSide == "r" {
				node.nodeData[node.numItems-1] = nil
				node.numItems--
				d := p.nodeData[i-1]
				sibling.insertItem(d)
				p.disconnectChild(0)
				p.connectChild(0, p.disconnectChild(1))
				for j := i; j < p.numItems; j++ {
					p.nodeData[j-1] = p.nodeData[j]
					if j+1 <= p.numItems {
						p.connectChild(j, p.disconnectChild(j+1))
					}
				}
				p.nodeData[p.numItems-1] = nil
				p.numItems--
				if p.numItems == 0 {
					if p != t.root {
						p = t.balanceTree(p)
					} else {
						t.root = sibling
					}
				}
				return p
			}
		}
	} else if sibling.numItems > 1 {
		f := 0
		if siblingSide == "l" {
			f = sibling.numItems - 1
		}
		for i := 0; i <= p.numItems; i++ {
			if p.children[i] == sibling && siblingSide == "l" {
				node.nodeData[0].data = p.nodeData[i].data
				p.nodeData[i].data = sibling.nodeData[f].data
				sibling.deleteNodeValue(sibling.nodeData[f].data)
				return p
			}
			if p.children[i] == sibling && siblingSide == "r" {
				node.nodeData[0].data = p.nodeData[i-1].data
				p.nodeData[i-1].data = sibling.nodeData[f].data
				sibling.deleteNodeValue(sibling.nodeData[f].data)
				return p
			}
		}
	}
	return nil
}

func (t *Tree) balanceTree(node *Node) *Node {
	siblingSide := "l"
	p := node.parent
	sibling := node.getSibling(-1)
	if sibling == nil {
		siblingSide = "r"
		sibling = p.children[1]
	}
	if sibling.numItems == 1 {
		for i := 0; i <= p.numItems; i++ {
			if p.children[i] == sibling && siblingSide == "l" {
				d := p.nodeData[i]
				sibling.insertItem(d)
				sibling.connectChild(sibling.numItems, node.disconnectChild(0))
				p.disconnectChild(i + 1)
				for j := i; j < p.numItems; j++ {
					if j+1 < p.numItems {
						p.nodeData[j] = p.nodeData[j+1]
						if j+2 <= p.numItems {
							p.connectChild(j+1, p.disconnectChild(j+2))
						}
					}
				}
				p.nodeData[p.numItems-1] = nil
				p.numItems--
				if p.numItems == 0 {
					if p != t.root {
						p = t.balanceTree(p)
					} else {
						t.root = sibling
					}
				}
				return p
			} else if p.children[i] == sibling && siblingSide == "r" {
				d := p.nodeData[i-1]
				sibling.insertItem(d)
				sibling.connectChild(0, node.disconnectChild(0))
				p.disconnectChild(0)
				p.connectChild(0, p.disconnectChild(1))
				for j := i; j < p.numItems; j++ {
					p.nodeData[j-1] = p.nodeData[j]
					if j+1 <= p.numItems {
						p.connectChild(j, p.disconnectChild(j+1))
					}
				}
				p.nodeData[p.numItems-1] = nil
				p.numItems--
				if p.numItems == 0 {
					if p != t.root {
						p = t.balanceTree(p)
					} else {
						t.root = sibling
					}
				}
				return p
			}
		}
	} else if sibling.numItems > 1 {
		f := 0
		if siblingSide == "l" {
			f = sibling.numItems - 1
		}
		for i := 0; i <= p.numItems; i++ {
			if p.children[i] == sibling && siblingSide == "l" {
				node.numItems++
				node.connectChild(1, node.disconnectChild(0))
				node.connectChild(0, sibling.disconnectChild(sibling.numItems))
				node.nodeData[0] = p.nodeData[i]
				p.nodeData[i] = sibling.nodeData[f]
				sibling.nodeData[sibling.numItems-1] = nil
				sibling.numItems--
				return p
			}
			if p.children[i] == sibling && siblingSide == "r" {
				node.numItems++
				node.nodeData[0] = p.nodeData[i-1]
				p.nodeData[i-1] = sibling.nodeData[f]
				node.connectChild(1, sibling.disconnectChild(f))
				for j := 0; j < sibling.numItems; j++ {
					if j+1 < sibling.numItems {
						sibling.nodeData[j] = sibling.nodeData[j+1]
					}
					sibling.connectChild(j, sibling.disconnectChild(j+1))
				}
				sibling.nodeData[sibling.numItems-1] = nil
				sibling.numItems--
				return p
			}
		}
	}
	return nil
}

func getInOrderNode(node *Node) *Node {
	if node.isLeaf() {
		return node
	}
	return getInOrderNode(node.children[0])
}

func (n *Node) connectChild(childNumber int, child *Node) {
	n.children[childNumber] = child
	if child != nil {
		child.parent = n
	}
}

func (n *Node) disconnectChild(childNumber int) *Node {
	temp := n.children[childNumber]
	n.children[childNumber] = nil
	return temp
}

func (n *Node) isLeaf() bool {
	if n.children[0] == nil {
		return true
	}
	return false
}

func (n *Node) isFull() bool {
	if n.numItems == 3 {
		return true
	}
	return false
}
func (n *Node) insertItem(value *NodeData) int {
	n.numItems++
	for i := 2; i >= 0; i-- {
		if n.nodeData[i] == nil {
			continue
		}
		if n.nodeData[i].data < value.data {
			n.nodeData[i+1] = value
			return i + 1
		}
		n.nodeData[i+1] = n.nodeData[i]
	}
	n.nodeData[0] = value
	return 0
}
func (n *Node) insertAtFront(value *NodeData) {
	n.numItems++
	for i := n.numItems - 1; i > 0; i-- {
		n.nodeData[i] = n.nodeData[i-1]
		n.connectChild(i+1, n.disconnectChild(i))
	}
	n.connectChild(1, n.disconnectChild(0))
	n.nodeData[0] = value
	n.connectChild(0, nil)
}
func (n *Node) removeItem() *NodeData {
	temp := n.nodeData[n.numItems-1]
	n.nodeData[n.numItems-1] = nil
	n.numItems--
	return temp
}
func (n *Node) deleteNodeValue(value int) {
	//only for leafs
	flag := -1
	for i := 0; i < n.numItems; i++ {
		if n.nodeData[i].data == value {
			flag = i
		}
		if flag != -1 && i+1 < n.numItems {
			n.nodeData[i].data = n.nodeData[i+1].data
		}
	}
	n.nodeData[n.numItems-1] = nil
	n.numItems--
}
func (n *Node) getSibling(value int) *Node {
	var result *Node
	p := n.parent
	if n.numItems != 0 {
		for i := 0; i <= p.numItems; i++ {
			if p.children[i].nodeData[0].data < value {
				result = p.children[i]
			}
		}
	} else {
		for i := 0; i <= p.numItems; i++ {
			if p.children[i].nodeData[0] == nil && i != 0 {
				result = p.children[i-1]
			}
		}
	}
	return result
}
