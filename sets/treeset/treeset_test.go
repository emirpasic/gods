// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treeset

import (
	"fmt"
	"testing"
)

func TestSetAdd(t *testing.T) {
	set := NewWithIntComparator()
	set.Add()
	set.Add(1)
	set.Add(2)
	set.Add(2, 3)
	set.Add()
	if actualValue := set.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := fmt.Sprintf("%d%d%d", set.Values()...), "123"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestSetContains(t *testing.T) {
	set := NewWithIntComparator()
	set.Add(3, 1, 2)
	if actualValue := set.Contains(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1, 2, 3); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1, 2, 3, 4); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestSetRemove(t *testing.T) {
	set := NewWithIntComparator()
	set.Add(3, 1, 2)
	set.Remove()
	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	set.Remove(1)
	if actualValue := set.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	set.Remove(3)
	set.Remove(3)
	set.Remove()
	set.Remove(2)
	if actualValue := set.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestSetEach(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")
	set.Each(func(index int, value interface{}) {
		switch index {
		case 0:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	})
}

func TestSetMap(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")
	mappedSet := set.Map(func(index int, value interface{}) interface{} {
		return "mapped: " + value.(string)
	})
	if actualValue, expectedValue := mappedSet.Contains("mapped: a", "mapped: b", "mapped: c"), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := mappedSet.Contains("mapped: a", "mapped: b", "mapped: x"), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if mappedSet.Size() != 3 {
		t.Errorf("Got %v expected %v", mappedSet.Size(), 3)
	}
}

func TestSetSelect(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")
	selectedSet := set.Select(func(index int, value interface{}) bool {
		return value.(string) >= "a" && value.(string) <= "b"
	})
	if actualValue, expectedValue := selectedSet.Contains("a", "b"), true; actualValue != expectedValue {
		fmt.Println("A: ", selectedSet.Contains("b"))
		t.Errorf("Got %v (%v) expected %v (%v)", actualValue, selectedSet.Values(), expectedValue, "[a b]")
	}
	if actualValue, expectedValue := selectedSet.Contains("a", "b", "c"), false; actualValue != expectedValue {
		t.Errorf("Got %v (%v) expected %v (%v)", actualValue, selectedSet.Values(), expectedValue, "[a b]")
	}
	if selectedSet.Size() != 2 {
		t.Errorf("Got %v expected %v", selectedSet.Size(), 3)
	}
}

func TestSetAny(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")
	any := set.Any(func(index int, value interface{}) bool {
		return value.(string) == "c"
	})
	if any != true {
		t.Errorf("Got %v expected %v", any, true)
	}
	any = set.Any(func(index int, value interface{}) bool {
		return value.(string) == "x"
	})
	if any != false {
		t.Errorf("Got %v expected %v", any, false)
	}
}

func TestSetAll(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")
	all := set.All(func(index int, value interface{}) bool {
		return value.(string) >= "a" && value.(string) <= "c"
	})
	if all != true {
		t.Errorf("Got %v expected %v", all, true)
	}
	all = set.All(func(index int, value interface{}) bool {
		return value.(string) >= "a" && value.(string) <= "b"
	})
	if all != false {
		t.Errorf("Got %v expected %v", all, false)
	}
}

func TestSetFind(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")
	foundIndex, foundValue := set.Find(func(index int, value interface{}) bool {
		return value.(string) == "c"
	})
	if foundValue != "c" || foundIndex != 2 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, "c", 2)
	}
	foundIndex, foundValue = set.Find(func(index int, value interface{}) bool {
		return value.(string) == "x"
	})
	if foundValue != nil || foundIndex != -1 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, nil, nil)
	}
}

func TestSetChaining(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")
}

func TestSetIteratorNextOnEmpty(t *testing.T) {
	set := NewWithStringComparator()
	it := set.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty set")
	}
}

func TestSetIteratorPrevOnEmpty(t *testing.T) {
	set := NewWithStringComparator()
	it := set.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty set")
	}
}

func TestSetIteratorNext(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")
	it := set.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := index, count-1; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestSetIteratorPrev(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")
	it := set.Iterator()
	for it.Prev() {
	}
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := index, count-1; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestSetIteratorBegin(t *testing.T) {
	m := NewWithStringComparator()
	it := m.Iterator()
	it.Begin()
	m.Add("a", "b", "c")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestSetIteratorEnd(t *testing.T) {
	set := NewWithStringComparator()
	it := set.Iterator()

	if index := it.Index(); index != -1 {
		t.Errorf("Got %v expected %v", index, -1)
	}

	it.End()
	if index := it.Index(); index != 0 {
		t.Errorf("Got %v expected %v", index, 0)
	}

	set.Add("a", "b", "c")
	it.End()
	if index := it.Index(); index != set.Size() {
		t.Errorf("Got %v expected %v", index, set.Size())
	}

	it.Prev()
	if index, value := it.Index(), it.Value(); index != set.Size()-1 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, set.Size()-1, "c")
	}
}

func TestSetIteratorFirst(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("a", "b", "c")
	it := set.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestSetIteratorLast(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("a", "b", "c")
	it := set.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 3, "c")
	}
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := NewWithIntComparator()
		for n := 0; n < 1000; n++ {
			set.Add(i)
		}
		for n := 0; n < 1000; n++ {
			set.Remove(n)
		}
	}
}
