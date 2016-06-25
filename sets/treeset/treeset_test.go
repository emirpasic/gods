/*
Copyright (c) Emir Pasic, All rights reserved.

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 3.0 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library. See the file LICENSE included
with this distribution for more information.
*/

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

func TestSetIterator(t *testing.T) {
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

	set.Clear()
	it = set.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty set")
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
