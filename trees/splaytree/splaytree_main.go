package main

// run go run splaytree_main.go in terminal

import (
	"fmt"
	"math/rand"

	"./splay"
)

type Tree struct {
	root *splay.Node
}

func (ST *Tree) SetRoot(n *splay.Node) {
	ST.root = n
}

func (ST *Tree) GetRoot() *splay.Node {
	return ST.root
}

func (ST *Tree) Ord(key1, key2 interface{}) int {
	if key1.(int) < key2.(int) {
		return 0
	} else if key1.(int) == key2.(int) {
		return 1
	} else {
		return 2
	}
}

func main() {
	ST := new(Tree)
	for i := 0; i < 30; i++ {
		err := splay.Insert(ST, rand.Int()%100, "hello")
		if err != nil {
			fmt.Println(err)
		} else {
			splay.Print(ST)
		}
	}
	for i := 0; i < 30; i++ {
		err := splay.Delete(ST, i)
		if err != nil {
			fmt.Println(err)
		} else {
			splay.Print(ST)
		}
	}
}
