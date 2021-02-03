package splay

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	key    interface{}
	value  interface{}
	parent *Node
	left   *Node
	right  *Node
}

type SplayTree interface {
	SetRoot(n *Node)
	GetRoot() *Node
	Ord(key1, key2 interface{}) int // 0 => LESS, 1 => EQUAL, 2 => GREATER
}

func Search(ST SplayTree, key interface{}) *Node {
	return SearchNode(ST, key, ST.GetRoot())
}

func SearchNode(ST SplayTree, key interface{}, n *Node) *Node {
	if n == nil {
		return nil
	} else {
		switch ST.Ord(key, n.key) {
		case 0:
			return SearchNode(ST, key, n.left)
		case 1:
			return n
		case 2:
			return SearchNode(ST, key, n.right)
		}
		return nil
	}
}

func Find(ST SplayTree, key interface{}) interface{} {
	return FindNode(ST, key, ST.GetRoot()) // // Returns value associated with key if it exists, or nil if not
}

func FindNode(ST SplayTree, key interface{}, n *Node) interface{} {
	if n == nil {
		return nil
	} else {
		switch ST.Ord(key, n.key) {
		case 0:
			return FindNode(ST, key, n.left)
		case 1:
			return n.value
		case 2:
			return FindNode(ST, key, n.right)
		}
		return nil
	}
}

func Insert(ST SplayTree, key, value interface{}) error {
	if Search(ST, key) != nil {
		s := fmt.Sprintf("splay.Insert Error: key %v already exists", key)
		return errors.New(s)
	}

	n := InsertNode(ST, key, value, ST.GetRoot())
	Splay(ST, n)
	return nil
}

func InsertNode(ST SplayTree, key interface{}, value interface{}, n *Node) *Node {
	if n == nil {
		_n := new(Node)
		_n.key = key
		_n.value = value
		ST.SetRoot(_n)
		return ST.GetRoot()
	}

	switch ST.Ord(key, n.key) {
	case 0:
		if n.left == nil {
			n.left = new(Node)
			n.left.key = key
			n.left.value = value
			n.left.parent = n
			return n.left
		} else {
			return InsertNode(ST, key, value, n.left)
		}
	case 2:
		if n.right == nil {
			n.right = new(Node)
			n.right.key = key
			n.right.value = value
			n.right.parent = n
			return n.right
		} else {
			return InsertNode(ST, key, value, n.right)
		}
	}
	return nil
}

func Delete(ST SplayTree, key interface{}) error {
	n := Search(ST, key)
	if n == nil {
		s := fmt.Sprintf("splay.Delete Error: key %v does not exist", key)
		return errors.New(s)
	} else {
		p := n.parent
		if n.left != nil {
			iop := InOrderPredecessor(n.left)
			Swap(n, iop)
			Remove(ST, iop)
		} else if n.right != nil {
			ios := InOrderSuccessor(n.right)
			Swap(n, ios)
			Remove(ST, ios)
		} else {
			Remove(ST, n)
		}

		if p != nil {
			Splay(ST, p)
		}
		return nil
	}
}

func Swap(n1, n2 *Node) {
	n1.key, n2.key = n2.key, n1.key
	n1.value, n2.value = n2.value, n1.value
}

func Remove(ST SplayTree, n *Node) {
	var isRoot, isLeft bool
	isRoot = (n == ST.GetRoot())
	if isRoot != true {
		isLeft = (n == n.parent.left)
	}

	if isRoot != true {
		if isLeft == true {
			if n.left != nil {
				n.parent.left = n.left
				n.left.parent = n.parent
			} else if n.right != nil {
				n.parent.left = n.right
				n.right.parent = n.parent
			} else {
				n.parent.left = nil
			}
		} else {
			if n.left != nil {
				n.parent.right = n.left
				n.left.parent = n.parent
			} else if n.right != nil {
				n.parent.right = n.right
				n.right.parent = n.parent
			} else {
				n.parent.right = nil
			}
		}
	}
	n = nil
}

func InOrderPredecessor(n *Node) *Node {
	if n.right == nil {
		return n
	} else {
		return InOrderPredecessor(n.right)
	}
}

func InOrderSuccessor(n *Node) *Node {
	if n.left == nil {
		return n
	} else {
		return InOrderSuccessor(n.left)
	}
}

func Splay(ST SplayTree, n *Node) {
	for n != ST.GetRoot() {
		if n.parent == ST.GetRoot() && n.parent.left == n {
			ZigL(ST, n)
		} else if n.parent == ST.GetRoot() && n.parent.right == n {
			ZigR(ST, n)
		} else if n.parent.left == n && n.parent.parent.left == n.parent {
			ZigZigL(ST, n)
		} else if n.parent.right == n && n.parent.parent.right == n.parent {
			ZigZigR(ST, n)
		} else if n.parent.right == n && n.parent.parent.left == n.parent {
			ZigZagLR(ST, n)
		} else {
			ZigZagRL(ST, n)
		}
	}
}

func ZigL(ST SplayTree, n *Node) {
	n.parent.left = n.right
	if n.right != nil {
		n.right.parent = n.parent
	}
	n.parent.parent = n
	n.right = n.parent
	n.parent = nil

	ST.SetRoot(n)
}

func ZigR(ST SplayTree, n *Node) {
	n.parent.right = n.left
	if n.left != nil {
		n.left.parent = n.parent
	}
	n.parent.parent = n
	n.left = n.parent
	n.parent = nil

	ST.SetRoot(n)
}

func ZigZigL(ST SplayTree, n *Node) {
	gg := n.parent.parent.parent

	var isRoot, isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = (gg.left == n.parent.parent)
	}

	n.parent.parent.left = n.parent.right
	if n.parent.right != nil {
		n.parent.right.parent = n.parent.parent
	}
	n.parent.left = n.right
	if n.right != nil {
		n.right.parent = n.parent
	}
	n.parent.right = n.parent.parent
	n.parent.parent.parent = n.parent
	n.right = n.parent
	n.parent.parent = n
	n.parent = gg

	if isRoot == true {
		ST.SetRoot(n)
	} else if isLeft == true {
		gg.left = n
	} else {
		gg.right = n
	}
}

func ZigZigR(ST SplayTree, n *Node) {
	gg := n.parent.parent.parent

	var isRoot, isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = (gg.left == n.parent.parent)
	}

	n.parent.parent.right = n.parent.left
	if n.parent.left != nil {
		n.parent.left.parent = n.parent.parent
	}
	n.parent.right = n.left
	if n.left != nil {
		n.left.parent = n.parent
	}
	n.parent.left = n.parent.parent
	n.parent.parent.parent = n.parent
	n.left = n.parent
	n.parent.parent = n
	n.parent = gg

	if isRoot == true {
		ST.SetRoot(n)
	} else if isLeft == true {
		gg.left = n
	} else {
		gg.right = n
	}
}

func ZigZagLR(ST SplayTree, n *Node) {
	gg := n.parent.parent.parent

	var isRoot, isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = (gg.left == n.parent.parent)
	}

	n.parent.parent.left = n.right
	if n.right != nil {
		n.right.parent = n.parent.parent
	}
	n.parent.right = n.left
	if n.left != nil {
		n.left.parent = n.parent
	}
	n.left = n.parent
	n.right = n.parent.parent
	n.parent.parent.parent = n
	n.parent.parent = n
	n.parent = gg

	if isRoot == true {
		ST.SetRoot(n)
	} else if isLeft == true {
		gg.left = n
	} else {
		gg.right = n
	}
}

func ZigZagRL(ST SplayTree, n *Node) {
	gg := n.parent.parent.parent

	var isRoot, isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = (gg.left == n.parent.parent)
	}

	n.parent.parent.right = n.left
	if n.left != nil {
		n.left.parent = n.parent.parent
	}
	n.parent.left = n.right
	if n.right != nil {
		n.right.parent = n.parent
	}
	n.right = n.parent
	n.left = n.parent.parent
	n.parent.parent.parent = n
	n.parent.parent = n
	n.parent = gg

	if isRoot == true {
		ST.SetRoot(n)
	} else if isLeft == true {
		gg.left = n
	} else {
		gg.right = n
	}
}

func Print(ST SplayTree) {
	PrintNode(ST.GetRoot(), 0) // Prints a directory-looking picture of ST to stdout
}

func PrintNode(n *Node, d int) {
	if n == nil {
		return
	}
	fmt.Println(strings.Repeat("-", 2*d), n.key, n.value)
	PrintNode(n.left, d+1)
	PrintNode(n.right, d+1)
}
