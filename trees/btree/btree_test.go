// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package btree

import (
	_ "fmt"
	"testing"
)

func TestBTree_search(t *testing.T) {
	{
		tree := NewWithIntComparator(3)
		tree.Root = &Node{Entries: []*Entry{}, Children: make([]*Node, 0)}
		tests := [][]interface{}{
			{0, 0, false},
		}
		for _, test := range tests {
			index, found := tree.search(tree.Root, &Entry{test[0], nil})
			if actualValue, expectedValue := index, test[1]; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
			if actualValue, expectedValue := found, test[2]; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	{
		tree := NewWithIntComparator(3)
		tree.Root = &Node{Entries: []*Entry{{2, 0}, {4, 1}, {6, 2}}, Children: []*Node{}}
		tests := [][]interface{}{
			{0, 0, false},
			{1, 0, false},
			{2, 0, true},
			{3, 1, false},
			{4, 1, true},
			{5, 2, false},
			{6, 2, true},
			{7, 3, false},
		}
		for _, test := range tests {
			index, found := tree.search(tree.Root, &Entry{test[0], nil})
			if actualValue, expectedValue := index, test[1]; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
			if actualValue, expectedValue := found, test[2]; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
}

func TestBTree_insert1(t *testing.T) {
	// https://upload.wikimedia.org/wikipedia/commons/3/33/B_tree_insertion_example.png
	tree := NewWithIntComparator(3)
	assertValidTree(t, tree, 0)

	tree.Put(1, 0)
	assertValidTree(t, tree, 1)
	assertValidTreeNode(t, tree.Root, 1, 0, []int{1})

	tree.Put(2, 1)
	assertValidTree(t, tree, 2)
	assertValidTreeNode(t, tree.Root, 2, 0, []int{1, 2})

	tree.Put(3, 2)
	assertValidTree(t, tree, 3)
	assertValidTreeNode(t, tree.Root, 1, 2, []int{2})
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []int{1})
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []int{3})

	tree.Put(4, 2)
	assertValidTree(t, tree, 4)
	assertValidTreeNode(t, tree.Root, 1, 2, []int{2})
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []int{1})
	assertValidTreeNode(t, tree.Root.Children[1], 2, 0, []int{3, 4})

	tree.Put(5, 2)
	assertValidTree(t, tree, 5)
	assertValidTreeNode(t, tree.Root, 2, 3, []int{2, 4})
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []int{1})
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []int{3})
	assertValidTreeNode(t, tree.Root.Children[2], 1, 0, []int{5})

	tree.Put(6, 2)
	assertValidTree(t, tree, 6)
	assertValidTreeNode(t, tree.Root, 2, 3, []int{2, 4})
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []int{1})
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []int{3})
	assertValidTreeNode(t, tree.Root.Children[2], 2, 0, []int{5, 6})

	tree.Put(7, 2)
	assertValidTree(t, tree, 7)
	assertValidTreeNode(t, tree.Root, 1, 2, []int{4})
	assertValidTreeNode(t, tree.Root.Children[0], 1, 2, []int{2})
	assertValidTreeNode(t, tree.Root.Children[1], 1, 2, []int{6})
	assertValidTreeNode(t, tree.Root.Children[0].Children[0], 1, 0, []int{1})
	assertValidTreeNode(t, tree.Root.Children[0].Children[1], 1, 0, []int{3})
	assertValidTreeNode(t, tree.Root.Children[1].Children[0], 1, 0, []int{5})
	assertValidTreeNode(t, tree.Root.Children[1].Children[1], 1, 0, []int{7})
}

func TestBTree_insert2(t *testing.T) {
	tree := NewWithIntComparator(4)
	assertValidTree(t, tree, 0)

	tree.Put(0, 0)
	assertValidTree(t, tree, 1)
	assertValidTreeNode(t, tree.Root, 1, 0, []int{0})

	tree.Put(2, 2)
	assertValidTree(t, tree, 2)
	assertValidTreeNode(t, tree.Root, 2, 0, []int{0, 2})

	tree.Put(1, 1)
	assertValidTree(t, tree, 3)
	assertValidTreeNode(t, tree.Root, 3, 0, []int{0, 1, 2})

	tree.Put(1, 1)
	assertValidTree(t, tree, 3)
	assertValidTreeNode(t, tree.Root, 3, 0, []int{0, 1, 2})

	tree.Put(3, 3)
	assertValidTree(t, tree, 4)
	assertValidTreeNode(t, tree.Root, 1, 2, []int{1})
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []int{0})
	assertValidTreeNode(t, tree.Root.Children[1], 2, 0, []int{2, 3})

	tree.Put(4, 4)
	assertValidTree(t, tree, 5)
	assertValidTreeNode(t, tree.Root, 1, 2, []int{1})
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []int{0})
	assertValidTreeNode(t, tree.Root.Children[1], 3, 0, []int{2, 3, 4})

	tree.Put(5, 5)
	assertValidTree(t, tree, 6)
	assertValidTreeNode(t, tree.Root, 2, 3, []int{1, 3})
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []int{0})
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []int{2})
	assertValidTreeNode(t, tree.Root.Children[2], 2, 0, []int{4, 5})
}

func TestBTree_insert3(t *testing.T) {
	// http://www.geeksforgeeks.org/b-tree-set-1-insert-2/
	tree := NewWithIntComparator(6)
	assertValidTree(t, tree, 0)

	tree.Put(10, 0)
	assertValidTree(t, tree, 1)
	assertValidTreeNode(t, tree.Root, 1, 0, []int{10})

	tree.Put(20, 1)
	assertValidTree(t, tree, 2)
	assertValidTreeNode(t, tree.Root, 2, 0, []int{10, 20})

	tree.Put(30, 2)
	assertValidTree(t, tree, 3)
	assertValidTreeNode(t, tree.Root, 3, 0, []int{10, 20, 30})

	tree.Put(40, 3)
	assertValidTree(t, tree, 4)
	assertValidTreeNode(t, tree.Root, 4, 0, []int{10, 20, 30, 40})

	tree.Put(50, 4)
	assertValidTree(t, tree, 5)
	assertValidTreeNode(t, tree.Root, 5, 0, []int{10, 20, 30, 40, 50})

	tree.Put(60, 5)
	assertValidTree(t, tree, 6)
	assertValidTreeNode(t, tree.Root, 1, 2, []int{30})
	assertValidTreeNode(t, tree.Root.Children[0], 2, 0, []int{10, 20})
	assertValidTreeNode(t, tree.Root.Children[1], 3, 0, []int{40, 50, 60})

	tree.Put(70, 6)
	assertValidTree(t, tree, 7)
	assertValidTreeNode(t, tree.Root, 1, 2, []int{30})
	assertValidTreeNode(t, tree.Root.Children[0], 2, 0, []int{10, 20})
	assertValidTreeNode(t, tree.Root.Children[1], 4, 0, []int{40, 50, 60, 70})

	tree.Put(80, 7)
	assertValidTree(t, tree, 8)
	assertValidTreeNode(t, tree.Root, 1, 2, []int{30})
	assertValidTreeNode(t, tree.Root.Children[0], 2, 0, []int{10, 20})
	assertValidTreeNode(t, tree.Root.Children[1], 5, 0, []int{40, 50, 60, 70, 80})

	tree.Put(90, 8)
	assertValidTree(t, tree, 9)
	assertValidTreeNode(t, tree.Root, 2, 3, []int{30, 60})
	assertValidTreeNode(t, tree.Root.Children[0], 2, 0, []int{10, 20})
	assertValidTreeNode(t, tree.Root.Children[1], 2, 0, []int{40, 50})
	assertValidTreeNode(t, tree.Root.Children[2], 3, 0, []int{70, 80, 90})
}

func assertValidTree(t *testing.T, tree *Tree, expectedSize int) {
	if actualValue, expectedValue := tree.size, expectedSize; actualValue != expectedValue {
		t.Errorf("Got %v expected %v for tree size", actualValue, expectedValue)
	}
}

func assertValidTreeNode(t *testing.T, node *Node, expectedEntries int, expectedChildren int, keys []int) {
	if actualValue, expectedValue := len(node.Entries), expectedEntries; actualValue != expectedValue {
		t.Errorf("Got %v expected %v for entries size", actualValue, expectedValue)
	}
	if actualValue, expectedValue := len(node.Children), expectedChildren; actualValue != expectedValue {
		t.Errorf("Got %v expected %v for children size", actualValue, expectedValue)
	}
	for i, key := range keys {
		if actualValue, expectedValue := node.Entries[i].Key, key; actualValue != expectedValue {
			t.Errorf("Got %v expected %v for key", actualValue, expectedValue)
		}
	}
}

func BenchmarkBTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := NewWithIntComparator(32)
		for n := 0; n < 1000; n++ {
			tree.Put(n, n)
		}
		for n := 0; n < 1000; n++ {
			tree.Remove(n)
		}
	}
}
