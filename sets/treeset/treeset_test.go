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

func TestTreeSet(t *testing.T) {

	set := NewWithIntComparator()

	// insertions
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

	// Asking if a set is superset of nothing, thus it's true
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

	// repeat 10 time since map in golang has a random iteration order each time and we want to make sure that the set is ordered
	for i := 1; i <= 10; i++ {
		if actualValue, expectedValue := fmt.Sprintf("%d%d%d", set.Values()...), "123"; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}

	set.Remove()
	set.Remove(1)

	if actualValue := set.Contains(1); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	set.Remove(1, 2, 3)

	if actualValue := set.Contains(3); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := set.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestTreeSetEnumerableAndIterator(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")

	// Each
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

	// Map
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

	// Select
	selectedSet := set.Select(func(index int, value interface{}) bool {
		return value.(string) >= "a" && value.(string) <= "b"
	})
	if actualValue, expectedValue := selectedSet.Contains("a", "b"), true; actualValue != expectedValue {
		fmt.Println("A: ", mappedSet.Contains("b"))
		t.Errorf("Got %v (%v) expected %v (%v)", actualValue, selectedSet.Values(), expectedValue, "[a b]")
	}
	if actualValue, expectedValue := selectedSet.Contains("a", "b", "c"), false; actualValue != expectedValue {
		t.Errorf("Got %v (%v) expected %v (%v)", actualValue, selectedSet.Values(), expectedValue, "[a b]")
	}
	if selectedSet.Size() != 2 {
		t.Errorf("Got %v expected %v", selectedSet.Size(), 3)
	}

	// Any
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

	// All
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

	// Find
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

	// Iterator
	it := set.Iterator()
	for it.Next() {
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
	}
	set.Clear()
	it = set.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty set")
	}
}

func BenchmarkTreeSet(b *testing.B) {
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
