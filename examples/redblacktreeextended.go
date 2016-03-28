/*
Copyright (c) 2015, Emir Pasic
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package examples

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type RedBlackTreeExtended struct {
	*rbt.Tree
}

func (tree *RedBlackTreeExtended) GetMin() (value interface{}, found bool) {
	node, found := tree.getMinFromNode(tree.Root)
	if node != nil {
		return node.Value, found
	} else {
		return nil, false
	}
}

func (tree *RedBlackTreeExtended) GetMax() (value interface{}, found bool) {
	node, found := tree.getMaxFromNode(tree.Root)
	if node != nil {
		return node.Value, found
	} else {
		return nil, false
	}
}

func (tree *RedBlackTreeExtended) RemoveMin() (value interface{}, deleted bool) {
	node, found := tree.getMinFromNode(tree.Root)
	if found {
		tree.Remove(node.Key)
		return node.Value, found
	} else {
		return nil, false
	}
}

func (tree *RedBlackTreeExtended) RemoveMax() (value interface{}, deleted bool) {
	node, found := tree.getMaxFromNode(tree.Root)
	if found {
		tree.Remove(node.Key)
		return node.Value, found
	} else {
		return nil, false
	}
}

func (tree *RedBlackTreeExtended) getMinFromNode(node *rbt.Node) (foundNode *rbt.Node, found bool) {
	if node == nil {
		return nil, false
	}
	if node.Left == nil {
		return node, true
	} else {
		return tree.getMinFromNode(node.Left)
	}
}

func (tree *RedBlackTreeExtended) getMaxFromNode(node *rbt.Node) (foundNode *rbt.Node, found bool) {
	if node == nil {
		return nil, false
	}
	if node.Right == nil {
		return node, true
	} else {
		return tree.getMaxFromNode(node.Right)
	}
}

func print(tree *RedBlackTreeExtended) {
	max, _ := tree.GetMax()
	min, _ := tree.GetMin()
	fmt.Printf("Value for max key: %v \n", max)
	fmt.Printf("Value for min key: %v \n", min)
	fmt.Println(tree)
}

// Find ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree is smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *RedBlackTreeExtended) Ceiling(key interface{}) (ceiling *rbt.Node, found bool) {
	found = false
	comparator := tree.Comparator()

	node := tree.Root
	for node != nil {
		compare := comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			ceiling, found = node, true
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	if found {
		return ceiling, true
	}
	return nil, false
}

// Find floor node of the input key, return the floor node or nil if no ceiling is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree is larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *RedBlackTreeExtended) Floor(key interface{}) (floor *rbt.Node, found bool) {
	found = false
	comparator := tree.Comparator()

	node := tree.Root
	for node != nil {
		compare := comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			node = node.Left
		case compare > 0:
			floor, found = node, true
			node = node.Right
		}
	}
	if found {
		return floor, true
	}
	return nil, false
}

func RedBlackTreeExtendedExample() {
	tree := RedBlackTreeExtended{rbt.NewWithIntComparator()}

	tree.Put(1, "a") // 1->x (in order)
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(3, "c") // 1->x, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->x, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->x, 2->b, 3->c, 4->d, 5->e (in order)

	print(&tree)
	// Value for max key: e
	// Value for min key: a
	// RedBlackTree
	// │       ┌── 5
	// │   ┌── 4
	// │   │   └── 3
	// └── 2
	//     └── 1

	tree.RemoveMin() // 2->b, 3->c, 4->d, 5->e (in order)
	tree.RemoveMax() // 2->b, 3->c, 4->d (in order)
	tree.RemoveMin() // 3->c, 4->d (in order)
	tree.RemoveMax() // 3->c (in order)

	print(&tree)
	// Value for max key: c
	// Value for min key: c
	// RedBlackTree
	// └── 3

	/*
	 * Ceiling and Floor functions
	 */
	tree = RedBlackTreeExtended{rbt.NewWithIntComparator()}
	tree.Put(1, "a")
	tree.Put(2, "b")
	tree.Put(4, "d")
	tree.Put(6, "f")
	tree.Put(7, "g")

	//index, ceiling, floor
	testValues := [][]interface{}{
		{0, 1, nil},
		{1, 1, 1},
		{2, 2, 2},
		{3, 4, 2},
		{4, 4, 4},
		{5, 6, 4},
		{6, 6, 6},
		{7, 7, 7},
		{8, nil, 7},
	}
	for _, tt := range testValues {
		actualCeiling, _ := tree.Ceiling(tt[0])
		actualFloor, _ := tree.Floor(tt[0])
		fmt.Printf("test key %d, expected (%d, %d), actual (%d, %d)\n", tt[0], tt[1], tt[2], actualCeiling.Key, actualFloor.Key)
	}
}
