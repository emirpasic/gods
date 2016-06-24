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

package treemap

import (
	"fmt"
	"testing"
)

func TestTreeMap(t *testing.T) {

	m := NewWithIntComparator()

	// insertions
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	// Test Size()
	if actualValue := m.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}

	// test Keys()
	if actualValue, expectedValue := fmt.Sprintf("%d%d%d%d%d%d%d", m.Keys()...), "1234567"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	// test Values()
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s%s%s%s", m.Values()...), "abcdefg"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	// test Min()
	if key, value := m.Min(); key != 1 || value != "a" {
		t.Errorf("Got %v expected %v", key, 1)
	}

	// test Max()
	if key, value := m.Max(); key != 7 || value != "g" {
		t.Errorf("Got %v expected %v", key, 7)
	}

	// key,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, nil, false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := m.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	// removals
	m.Remove(5)
	m.Remove(6)
	m.Remove(7)
	m.Remove(8)
	m.Remove(5)

	// Test Keys()
	if actualValue, expectedValue := fmt.Sprintf("%d%d%d%d", m.Keys()...), "1234"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	// test Values()
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s", m.Values()...), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	// Test Size()
	if actualValue := m.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}

	tests2 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, nil, false},
		{6, nil, false},
		{7, nil, false},
		{8, nil, false},
	}

	for _, test := range tests2 {
		// retrievals
		actualValue, actualFound := m.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	// removals
	m.Remove(1)
	m.Remove(4)
	m.Remove(2)
	m.Remove(3)
	m.Remove(2)
	m.Remove(2)

	// Test Keys()
	if actualValue, expectedValue := fmt.Sprintf("%s", m.Keys()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	// test Values()
	if actualValue, expectedValue := fmt.Sprintf("%s", m.Values()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	// Test Size()
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	// Test Empty()
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	m.Put(1, "a")
	m.Put(2, "b")
	m.Clear()

	// Test Empty()
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	// test Min()
	if key, value := m.Min(); key != nil || value != nil {
		t.Errorf("Got %v expected %v", key, nil)
	}

	// test Max()
	if key, value := m.Max(); key != nil || value != nil {
		t.Errorf("Got %v expected %v", key, nil)
	}
}

func TestTreeMapEnumerableAndIterator(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)

	// Each
	count := 0
	m.Each(func(key interface{}, value interface{}) {
		count += 1
		if actualValue, expectedValue := count, value; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		switch value {
		case 1:
			if actualValue, expectedValue := key, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := key, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 3:
			if actualValue, expectedValue := key, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	})

	// Map
	mappedMap := m.Map(func(key1 interface{}, value1 interface{}) (key2 interface{}, value2 interface{}) {
		return key1, "mapped: " + key1.(string)
	})
	if actualValue, _ := mappedMap.Get("a"); actualValue != "mapped: a" {
		t.Errorf("Got %v expected %v", actualValue, "mapped: a")
	}
	if actualValue, _ := mappedMap.Get("b"); actualValue != "mapped: b" {
		t.Errorf("Got %v expected %v", actualValue, "mapped: b")
	}
	if actualValue, _ := mappedMap.Get("c"); actualValue != "mapped: c" {
		t.Errorf("Got %v expected %v", actualValue, "mapped: c")
	}
	if mappedMap.Size() != 3 {
		t.Errorf("Got %v expected %v", mappedMap.Size(), 3)
	}

	// Select
	selectedMap := m.Select(func(key interface{}, value interface{}) bool {
		return key.(string) >= "a" && key.(string) <= "b"
	})
	if actualValue, _ := selectedMap.Get("a"); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, "value: a")
	}
	if actualValue, _ := selectedMap.Get("b"); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, "value: b")
	}
	if selectedMap.Size() != 2 {
		t.Errorf("Got %v expected %v", selectedMap.Size(), 3)
	}

	// Any
	any := m.Any(func(key interface{}, value interface{}) bool {
		return value.(int) == 3
	})
	if any != true {
		t.Errorf("Got %v expected %v", any, true)
	}
	any = m.Any(func(key interface{}, value interface{}) bool {
		return value.(int) == 4
	})
	if any != false {
		t.Errorf("Got %v expected %v", any, false)
	}

	// All
	all := m.All(func(key interface{}, value interface{}) bool {
		return key.(string) >= "a" && key.(string) <= "c"
	})
	if all != true {
		t.Errorf("Got %v expected %v", all, true)
	}
	all = m.All(func(key interface{}, value interface{}) bool {
		return key.(string) >= "a" && key.(string) <= "b"
	})
	if all != false {
		t.Errorf("Got %v expected %v", all, false)
	}

	// Find
	foundKey, foundValue := m.Find(func(key interface{}, value interface{}) bool {
		return key.(string) == "c"
	})
	if foundKey != "c" || foundValue != 3 {
		t.Errorf("Got %v -> %v expected %v -> %v", foundKey, foundValue, "c", 3)
	}
	foundKey, foundValue = m.Find(func(key interface{}, value interface{}) bool {
		return key.(string) == "x"
	})
	if foundKey != nil || foundValue != nil {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundKey, nil, nil)
	}

	// Iterator
	it := m.Iterator()
	for it.Next() {
		key := it.Key()
		value := it.Value()
		switch key {
		case "a":
			if actualValue, expectedValue := value, 1; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "b":
			if actualValue, expectedValue := value, 2; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "c":
			if actualValue, expectedValue := value, 3; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	}
	m.Clear()
	it = m.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty map")
	}
}

func BenchmarkTreeMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewWithIntComparator()
		for n := 0; n < 1000; n++ {
			m.Put(n, n)
		}
		for n := 0; n < 1000; n++ {
			m.Remove(n)
		}
	}
}
