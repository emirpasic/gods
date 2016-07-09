package main

import (
	"github.com/emirpasic/gods/trees/btree"
	"fmt"
)

func main() {
	tree := btree.NewWithIntComparator(3)
	tree.Put(1, 0)
	tree.Put(2, 1)
	tree.Put(3, 2)
	fmt.Println()
}
