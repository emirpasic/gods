package intervaltree

import (
	"errors"
)

type Interval struct {
	start int
	end   int
}

func NewInterval(start int, end int) (Interval, error) {
	var i Interval
	if start > end {
		return i, errors.New("Start must be smaller than End")
	} else if start == end {
		return i, errors.New("Start must be equals to end")
	} else {
		return Interval{start, end}, nil
	}
}

func (i Interval) Start() int {
	return i.start
}

func (i Interval) End() int {
	return i.end
}

func Overlaps(i1 Interval, i2 Interval) bool {
	i2_contains_i1 := (i1.start < i2.end && i1.start > i2.start) || (i1.end < i2.end && i1.end > i2.start)
	i1_contains_i2 := (i2.start < i1.end && i2.start > i1.start) || (i2.end < i1.end && i2.end > i1.start)

	return i2_contains_i1 || i1_contains_i2
}

type IntervalTreeNode struct {
	i          Interval
	subTreeMax int
	left       *IntervalTreeNode
	right      *IntervalTreeNode
}

type IntervalTree struct {
	root *IntervalTreeNode
}

func NewIntervalTree() *IntervalTree {
	return &IntervalTree{nil}
}

func (tree *IntervalTree) Empty() bool {
	return tree.root == nil
}

func (tree *IntervalTree) Insert(i Interval) {
	if tree.Empty() {
		tree.root = newIntervalTreeNode(i)
	} else {
		tree.root.insert(i)
	}
}

func (tree *IntervalTree) FindOverlap(i Interval) []Interval {
	if tree.Empty() {
		var overlaps []Interval
		return overlaps
	} else {
		overlaps := tree.root.findOverlap(i)
		return overlaps
	}
}

func (tree *IntervalTree) Overlaps(i Interval) bool {
	if tree.Empty() {
		return false
	} else {
		return tree.root.overlaps(i)
	}
}

func newIntervalTreeNode(i Interval) *IntervalTreeNode {
	node := new(IntervalTreeNode)
	node.i = i
	node.subTreeMax = i.End()
	return node
}

func (node *IntervalTreeNode) insert(i Interval) *IntervalTreeNode {
	start := node.i.Start()

	if i.End() < start {
		if node.left == nil {
			node.left = newIntervalTreeNode(i)
		} else {
			node.left.insert(i)
		}
	} else {
		if node.right == nil {
			node.right = newIntervalTreeNode(i)
		} else {
			node.right.insert(i)
		}
	}

	if node.subTreeMax < i.End() {
		node.subTreeMax = i.End()
	}

	return node
}

func (node *IntervalTreeNode) findOverlap(i Interval) []Interval {
	var overlaps []Interval

	if Overlaps(node.i, i) {
		overlaps = append(overlaps, node.i)
	}

	if node.left != nil {
		overlaps = append(overlaps, node.left.findOverlap(i)...)
	}

	if node.right != nil {
		overlaps = append(overlaps, node.right.findOverlap(i)...)
	}

	return overlaps
}

func (node *IntervalTreeNode) overlaps(i Interval) bool {
	if Overlaps(node.i, i) {
		return true
	} else if node.left != nil && node.left.overlaps(i) {
		return true
	} else if node.right != nil && node.right.overlaps(i) {
		return true
	} else {
		return false
	}
}
