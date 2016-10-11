package redblacktree

import (
	"testing"
)

func TestRedBlackTreeIteratorNextOnEmpty(t *testing.T) {
	tree := NewWithIntComparator()
	it := tree.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestRedBlackTreeIteratorPrevOnEmpty(t *testing.T) {
	tree := NewWithIntComparator()
	it := tree.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestRedBlackTreeIterator1Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	//│       ┌── 7
	//│   ┌── 6
	//│   │   └── 5
	//└── 4
	//	  │   ┌── 3
	//	  └── 2
	//		  └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator1Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	//│       ┌── 7
	//│   ┌── 6
	//│   │   └── 5
	//└── 4
	//    │   ┌── 3
	//    └── 2
	//        └── 1
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator2Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator2Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator3Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(1, "a")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator3Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(1, "a")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator4Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	//│       ┌── 27
	//│   ┌── 25
	//│   │   │   ┌── 22
	//│   │   └──(17)
	//│   │       └── 15
	//└── 13
	//    │   ┌── 11
	//    └── 8
	//        └── 6
	//            └──(1)
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		value := it.Value()
		switch value {
		case count:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator4Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	//│       ┌── 27
	//│   ┌── 25
	//│   │   │   ┌── 22
	//│   │   └──(17)
	//│   │       └── 15
	//└── 13
	//		│   ┌── 11
	//		└── 8
	//			└── 6
	//				└──(1)
	it := tree.Iterator()
	count := tree.Size()
	for it.Next() {
	}
	for it.Prev() {
		value := it.Value()
		switch value {
		case count:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		count--
	}
	if actualValue, expectedValue := count, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIteratorBegin(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Begin()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	for it.Next() {
	}

	it.Begin()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Next()
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestRedBlackTreeIteratorEnd(t *testing.T) {
	tree := NewWithIntComparator()
	it := tree.Iterator()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.End()
	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it.End()
	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Prev()
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestRedBlackTreeIteratorFirst(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestRedBlackTreeIteratorLast(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestRedBlackTreeIteratorCombinedPrevNext(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	//│       ┌── 27
	//│   ┌── 25
	//│   │   │   ┌── 22
	//│   │   └──(17)
	//│   │       └── 15
	//└── 13
	//	  │   ┌── 11
	//	  └── 8
	//		  └── 6
	//			  └──(1)
	it := tree.Iterator()
	it.First()
	for i:=1 ; i<9 ; i++ {
		if (it.Value() != i) {
			t.Errorf("Expected %v and instead got %v", i, it.Value())
		}
		it.Next()
	}
	for i:=9 ; i>3 ; i-- {
		if (it.Value() != i) {
			t.Errorf("Expected %v and instead got %v", i, it.Value())
		}
		it.Prev()
	}
	for i:=4 ; it.Next() ; i++ {
		if (it.Value() != i) {
			t.Errorf("Expected %v and instead got %v", i, it.Value())
		}
	}
}

func TestRedBlackTreeRangedIteratorOnEmptyTree(t *testing.T) {
	tree := NewWithIntComparator()
	it := tree.IteratorWithin(10, 15)
	if (it.First()) {
		t.Errorf("There should be no key available through the iterator")
	}
}

func TestRedBlackTreeRangedIteratorNext(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	//│       ┌── 27
	//│   ┌── 25
	//│   │   │   ┌── 22
	//│   │   └──(17)
	//│   │       └── 15
	//└── 13
	//	  │   ┌── 11
	//	  └── 8
	//		  └── 6
	//			  └──(1)
	it := tree.IteratorWithin(8, 15)
	it.First()
	for i:=3 ; i<7 ; i++ {
		if (it.Value() != i) {
			t.Errorf("Expected %v and instead got %v", i, it.Value())
		}
		it.Next()
	}
	if (it.Next()) {
		t.Errorf("There should be no more values from this iterator")
	}
}

func TestRedBlackTreeRangedIteratorEnd(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	//│       ┌── 27
	//│   ┌── 25
	//│   │   │   ┌── 22
	//│   │   └──(17)
	//│   │       └── 15
	//└── 13
	//	  │   ┌── 11
	//	  └── 8
	//		  └── 6
	//			  └──(1)
	it := tree.IteratorWithin(8, 15)
	it.First()
	it.End()
	if (it.Next()) {
		t.Errorf("There should be no more values from this iterator")
	}
}

func TestRedBlackTreeRangedIteratorPrev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	//│       ┌── 27
	//│   ┌── 25
	//│   │   │   ┌── 22
	//│   │   └──(17)
	//│   │       └── 15
	//└── 13
	//	  │   ┌── 11
	//	  └── 8
	//		  └── 6
	//			  └──(1)
	it := tree.IteratorWithin(8, 15)
	it.Last()
	for i:=6 ; i>2 ; i-- {
		if (it.Value() != i) {
			t.Errorf("Expected %v and instead got %v", i, it.Value())
		}
		it.Prev()
	}
	if (it.Prev()) {
		t.Errorf("There should be no more values from this iterator")
	}
}


func TestRedBlackTreeRangedIteratorCombinedPrevNext(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	//│       ┌── 27
	//│   ┌── 25
	//│   │   │   ┌── 22
	//│   │   └──(17)
	//│   │       └── 15
	//└── 13
	//	  │   ┌── 11
	//	  └── 8
	//		  └── 6
	//			  └──(1)
	it := tree.IteratorWithin(8, 15)
	it.First()
	for i:=3 ; i<6 ; i++ {
		if (it.Value() != i) {
			t.Errorf("Expected %v and instead got %v", i, it.Value())
		}
		it.Next()
	}
	for i:=6 ; i>3 ; i-- {
		if (it.Value() != i) {
			t.Errorf("Expected %v and instead got %v", i, it.Value())
		}
		it.Prev()
	}
	for i:=4 ; it.Next() ; i++ {
		if (it.Value() != i) {
			t.Errorf("Expected %v and instead got %v", i, it.Value())
		}
	}
}
