[![Build Status](https://travis-ci.org/emirpasic/gods.svg)](https://travis-ci.org/emirpasic/gods) 

# GoDS (Go Data Structures)

Implementation of various data structures in Go. 

## Data Structures

- [Sets](#sets)
  - [HashSet](#hashset)
  - [TreeSet](#treeset)
- [Lists](#lists)
  - [ArrayList](#arraylist)
- [Stacks](#stacks)
  - [LinkedListStack](#linkedliststack)
  - [ArrayStack](#arraystack)
- [Maps](#maps)
  - [HashMap](#hashmap)
  - [TreeMap](#treemap)
- [Trees](#trees)
  - [RedBlackTree](#redblacktree)
- [Functions](#functions)
  - [Comparator](#comparator)
  

###Sets 

A set is a data structure that can store elements and no repeated values. It is a computer implementation of the mathematical concept of a finite set. Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests an element for membership in a set. This structed is often used to ensure that no duplicates are present in a collection.

All sets implement the set interface with the following methods:

```go
	Add(items ...interface{})
	Remove(items ...interface{})
	Contains(items ...interface{}) bool
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}

```

####HashSet

This structure implements the Set interface and is backed by a hash table (actually a Go's map). It makes no guarantees as to the iteration order of the set, since Go randomizes this iteration order on maps.

This structure offers constant time performance for the basic operations (add, remove, contains and size).

```go
package main

import "github.com/emirpasic/gods/sets/hashset"

func main() {
	set := hashset.New()   // empty
	set.Add(1)             // 1
	set.Add(2, 2, 3, 4, 5) // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	set.Remove(4)          // 5, 3, 2, 1 (random order)
	set.Remove(2, 3)       // 1, 5 (random order)
	set.Contains(1)        // true
	set.Contains(1, 5)     // true
	set.Contains(1, 6)     // false
	_ = set.Values()       // []int{5,1} (random order)
	set.Clear()            // empty
	set.Empty()            // true
	set.Size()             // 0
}

	
```

####TreeSet

This structure implements the Set interface and is backed by a red-black tree to keep the elements sorted with respect to the comparator.

This implementation provides guaranteed log(n) time cost for the basic operations (add, remove and contains).

```go
package main

import "github.com/emirpasic/gods/sets/treeset"

func main() {
	set := treeset.NewWithIntComparator() // empty (keys are of type int)
	set.Add(1)                            // 1
	set.Add(2, 2, 3, 4, 5)                // 1, 2, 3, 4, 5 (in order, duplicates ignored)
	set.Remove(4)                         // 1, 2, 3, 5 (in order)
	set.Remove(2, 3)                      // 1, 5 (in order)
	set.Contains(1)                       // true
	set.Contains(1, 5)                    // true
	set.Contains(1, 6)                    // false
	_ = set.Values()                      // []int{1,5} (in order)
	set.Clear()                           // empty
	set.Empty()                           // true
	set.Size()                            // 0
}

```

###Lists

####ArrayList

```go
package main

import "github.com/emirpasic/gods/lists/arraylist"

func main() {
	list := arraylist.New()

	list.Add("a")      // ["a"]
	list.Add("b", "c") // ["a","b","c"]

	_, _ = list.Get(0)   // "a",true
	_, _ = list.Get(100) // nil,false

	_ = list.Contains("a", "b", "c")      //true
	_ = list.Contains("a", "b", "c", "d") //false

	list.Remove(2) // ["a","b"]
	list.Remove(1) // ["a"]
	list.Remove(0) // []
	list.Remove(0) // [] (ignored)

	_ = list.Empty() // true
	_ = list.Size()  // 0

	list.Add("a") // ["a"]
	list.Clear()  // []

}
```

###Stacks

The stack interface represents a last-in-first-out (LIFO) collection of objects. The usual push and pop operations are provided, as well as a method to peek at the top item on the stack, a method to check whether the stack is empty and the size (number of elements).

All stacks implement the stack interface with the following methods:
```go
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)
	Empty() bool
	Size() int
	Clear()
```

####LinkedListStack

This stack structure is based on a linked list, i.e. each previous element has a point to the next. 

All operations are guaranted constant time performance.

```go
package main

import lls "github.com/emirpasic/gods/stacks/linkedliststack"

func main() {
	stack := lls.New()  // empty
	stack.Push(1)       // 1
	stack.Push(2)       // 1, 2
	_, _ = stack.Peek() // 2,true
	_, _ = stack.Pop()  // 2, true
	_, _ = stack.Pop()  // 1, true
	_, _ = stack.Pop()  // nil, false (nothing to pop)
	stack.Push(1)       // 1
	stack.Clear()       // empty
	stack.Empty()       // true
	stack.Size()        // 0
}

```

####ArrayStack

This stack structure is based on a array.

All operations are guaranted constant time performance.

```go
package main

import "github.com/emirpasic/gods/stacks/arraystack"

func main() {
	stack := arraystack.New() // empty
	stack.Push(1)             // 1
	stack.Push(2)             // 1, 2
	_, _ = stack.Peek()       // 2,true
	_, _ = stack.Pop()        // 2, true
	_, _ = stack.Pop()        // 1, true
	_, _ = stack.Pop()        // nil, false (nothing to pop)
	stack.Push(1)             // 1
	stack.Clear()             // empty
	stack.Empty()             // true
	stack.Size()              // 0
}


```

###Maps

Structure that maps keys to values. A map cannot contain duplicate keys and each key can map to at most one value.

All maps implement the map interface with the following methods:
```go
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, found bool)
	Remove(key interface{})
	Empty() bool
	Size() int
	Clear()
	Keys() []interface{}
	Values() []interface{}
```

####HashMap

Map structure based on hash tables, more exactly, Go's map. Keys are unordered. 

All operations are guaranted constant time performance, except _Key()_ and _Values()_ retrieval that of linear time performance.

```go
package main

import "github.com/emirpasic/gods/maps/hashmap"

func main() {
	m := hashmap.New() // empty
	m.Put(1, "x")      // 1->x
	m.Put(2, "b")      // 2->b, 1->x  (random order)
	m.Put(1, "a")      // 2->b, 1->a (random order)
	_, _ = m.Get(2)    // b, true
	_, _ = m.Get(3)    // nil, false
	_ = m.Values()     // []interface {}{"b", "a"} (random order)
	_ = m.Keys()       // []interface {}{1, 2} (random order)
	m.Remove(1)        // 2->b
	m.Clear()          // empty
	m.Empty()          // true
	m.Size()           // 0
}

```

####TreeMap

Map structure based on our red-black tree implementation. Keys are ordered with respect to the passed comparator. 

_Put()_, _Get()_ and _Remove()_ are guaranteed log(n) time performance.

_Key()_ and _Values()_ methods return keys and values respectively in order of the keys. These meethods are quaranteed linear time performance.

```go
package main

import "github.com/emirpasic/gods/maps/treemap"

func main() {
	m := treemap.NewWithIntComparator() // empty (keys are of type int)
	m.Put(1, "x")                       // 1->x
	m.Put(2, "b")                       // 1->x, 2->b (in order)
	m.Put(1, "a")                       // 1->a, 2->b (in order)
	_, _ = m.Get(2)                     // b, true
	_, _ = m.Get(3)                     // nil, false
	_ = m.Values()                      // []interface {}{"a", "b"} (in order)
	_ = m.Keys()                        // []interface {}{1, 2} (in order)
	m.Remove(1)                         // 2->b
	m.Clear()                           // empty
	m.Empty()                           // true
	m.Size()                            // 0
}


```

###Trees

A tree is a widely used data data structure that simulates a hierarchical tree structure, with a root value and subtrees of children, represented as a set of linked nodes; thus no cyclic links.
 
####RedBlackTree

A red–black tree is a binary search tree with an extra bit of data per node, its color, which can be either red or black. The extra bit of storage ensures an approximately balanced tree by constraining how nodes are colored from any path from the root to the leaf. Thus, it is a data structure which is a type of self-balancing binary search tree.[Wikipedia](http://en.wikipedia.org/wiki/Red%E2%80%93black_tree)

The balancing of the tree is not perfect but it is good enough to allow it to guarantee searching in O(log n) time, where n is the total number of elements in the tree. The insertion and deletion operations, along with the tree rearrangement and recoloring, are also performed in O(log n) time.[Wikipedia](http://en.wikipedia.org/wiki/Red%E2%80%93black_tree)

[![Build Status](http://upload.wikimedia.org/wikipedia/commons/thumb/6/66/Red-black_tree_example.svg/500px-Red-black_tree_example.svg.png)](http://en.wikipedia.org/wiki/Red%E2%80%93black_tree)

```go
package main

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	tree := rbt.NewWithIntComparator() // empty(keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)

	fmt.Println(m)
	//
	//  RedBlackTree
	//  │           ┌── 6
	//	│       ┌── 5
	//	│   ┌── 4
	//	│   │   └── 3
	//	└── 2
	//		└── 1

	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(m)
	//
	//  RedBlackTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      │   ┌── 3
	//      └── 1

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0
}

```

### Functions

Various helper functions used throughout the library.

#### Comparator

Some data structures (e.g. TreeMap, TreeSet) require a comparator function to sort their contained elements. This comparator is necessary during the initalization.

Comparator is defined as:


```go
Return values:

  -1, if a < b
   0, if a == b
   1, if a > b
     
Comparator signature:

  type Comparator func(a, b interface{}) int
```

Two common comparators are included in the library:

#####IntComparator
```go
func IntComparator(a, b interface{}) int {
	aInt := a.(int)
	bInt := b.(int)
	switch {
	case aInt > bInt:
		return 1
	case aInt < bInt:
		return -1
	default:
		return 0
	}
}
```

#####StringComparator
```go
func StringComparator(a, b interface{}) int {
	s1 := a.(string)
	s2 := b.(string)
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}
```

#####CustomComparator
```go
package main

import (
	"fmt"
	"github.com/emirpasic/gods/sets/treeset"
)

type User struct {
	id   int
	name string
}

// Comparator function (sort by IDs)
func byID(a, b interface{}) int {

	// Type assertion, program will panic if this is not respected
	c1 := a.(User)
	c2 := b.(User)

	switch {
	case c1.id > c2.id:
		return 1
	case c1.id < c2.id:
		return -1
	default:
		return 0
	}
}

func main() {
	set := treeset.NewWith(byID) 

	set.Add(User{2, "Second"})
	set.Add(User{3, "Third"})
	set.Add(User{1, "First"})
	set.Add(User{4, "Fourth"})
	
	fmt.Println(set) // {1 First}, {2 Second}, {3 Third}, {4 Fourth}
}
```


## Motivations

Collections and data structures found in other languages: Java Collections, C++ Standard Template Library (STL) containers, Qt Containers, etc. 

## Goals

**Fast algorithms**: 

  - Based on decades of knowledge and experiences of other libraries mentioned below.

**Memory efficient algorithms**: 
  
  - Avoiding to consume memory by using optimal algorithms and data structures for the given set of problems, e.g. red-black tree in case of TreeMap to avoid keeping redundant sorted array of keys in memory.

**Easy to use library**: 
  
  - Well-structued library with minimalistic set of atomic operations from which more complex operations can be crafted.

**Stable library**: 
  
  - Only additions are permitted keeping the library backward compatible.

**Solid documentation and examples**: 
  
  - Learning by example.

**Production ready**: 

  - Still waiting for the project to mature and be used in some heavy back-end tasks.

There is often a tug of war between speed and memory when crafting algorithms. We choose to optimize for speed in most cases within reasonable limits on memory consumption.

Thread safety is not a concern of this project, this should be handled at a higher level.

## Testing and Benchmarking

`go test -v -bench . -benchmem  -benchtime 1s ./...`

## Contributing

Biggest contribution towards this library is to use it and give us feedback for further improvements and additions.

For direct contributions, branch of from master and do _pull request_.

## License

This library is distributed under the BSD-style license found in the [LICENSE](https://github.com/emirpasic/gods/blob/master/LICENSE) file.
