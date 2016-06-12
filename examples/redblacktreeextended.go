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
}
