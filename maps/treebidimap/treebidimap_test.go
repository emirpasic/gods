// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treebidimap

import (
	"fmt"
	"testing"

	"github.com/habohai/gods/utils"
)

func TestMapPut(t *testing.T) {
	m := NewWith(utils.IntComparator, utils.StringComparator)
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	if actualValue := m.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}
	if actualValue, expectedValue := m.Keys(), []interface{}{1, 2, 3, 4, 5, 6, 7}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := m.Values(), []interface{}{"a", "b", "c", "d", "e", "f", "g"}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
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
}

func TestMapRemove(t *testing.T) {
	m := NewWith(utils.IntComparator, utils.StringComparator)
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	m.Remove(5)
	m.Remove(6)
	m.Remove(7)
	m.Remove(8)
	m.Remove(5)

	if actualValue, expectedValue := m.Keys(), []interface{}{1, 2, 3, 4}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, expectedValue := m.Values(), []interface{}{"a", "b", "c", "d"}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := m.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
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
		actualValue, actualFound := m.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	m.Remove(1)
	m.Remove(4)
	m.Remove(2)
	m.Remove(3)
	m.Remove(2)
	m.Remove(2)

	if actualValue, expectedValue := fmt.Sprintf("%s", m.Keys()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s", m.Values()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestMapGetKey(t *testing.T) {
	m := NewWith(utils.IntComparator, utils.StringComparator)
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	// key,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{nil, "x", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := m.GetKey(test[1])
		if actualValue != test[0] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[0])
		}
	}
}

func sameElements(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestMapEach(t *testing.T) {
	m := NewWith(utils.StringComparator, utils.IntComparator)
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	count := 0
	m.Each(func(key interface{}, value interface{}) {
		count++
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
}

func TestMapMap(t *testing.T) {
	m := NewWith(utils.StringComparator, utils.IntComparator)
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	mappedMap := m.Map(func(key1 interface{}, value1 interface{}) (key2 interface{}, value2 interface{}) {
		return key1, value1.(int) * value1.(int)
	})
	if actualValue, _ := mappedMap.Get("a"); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, "mapped: a")
	}
	if actualValue, _ := mappedMap.Get("b"); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, "mapped: b")
	}
	if actualValue, _ := mappedMap.Get("c"); actualValue != 9 {
		t.Errorf("Got %v expected %v", actualValue, "mapped: c")
	}
	if mappedMap.Size() != 3 {
		t.Errorf("Got %v expected %v", mappedMap.Size(), 3)
	}
}

func TestMapSelect(t *testing.T) {
	m := NewWith(utils.StringComparator, utils.IntComparator)
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
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
		t.Errorf("Got %v expected %v", selectedMap.Size(), 2)
	}
}

func TestMapAny(t *testing.T) {
	m := NewWith(utils.StringComparator, utils.IntComparator)
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
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
}

func TestMapAll(t *testing.T) {
	m := NewWith(utils.StringComparator, utils.IntComparator)
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
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
}

func TestMapFind(t *testing.T) {
	m := NewWith(utils.StringComparator, utils.IntComparator)
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
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
}

func TestMapChaining(t *testing.T) {
	m := NewWith(utils.StringComparator, utils.IntComparator)
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	chainedMap := m.Select(func(key interface{}, value interface{}) bool {
		return value.(int) > 1
	}).Map(func(key interface{}, value interface{}) (interface{}, interface{}) {
		return key.(string) + key.(string), value.(int) * value.(int)
	})
	if actualValue := chainedMap.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, found := chainedMap.Get("aa"); actualValue != nil || found {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue, found := chainedMap.Get("bb"); actualValue != 4 || !found {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue, found := chainedMap.Get("cc"); actualValue != 9 || !found {
		t.Errorf("Got %v expected %v", actualValue, 9)
	}
}

func TestMapIteratorNextOnEmpty(t *testing.T) {
	m := NewWithStringComparators()
	it := m.Iterator()
	it = m.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty map")
	}
}

func TestMapIteratorPrevOnEmpty(t *testing.T) {
	m := NewWithStringComparators()
	it := m.Iterator()
	it = m.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty map")
	}
}

func TestMapIteratorNext(t *testing.T) {
	m := NewWith(utils.StringComparator, utils.IntComparator)
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)

	it := m.Iterator()
	count := 0
	for it.Next() {
		count++
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
		if actualValue, expectedValue := value, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapIteratorPrev(t *testing.T) {
	m := NewWith(utils.StringComparator, utils.IntComparator)
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)

	it := m.Iterator()
	for it.Next() {
	}
	countDown := m.Size()
	for it.Prev() {
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
		if actualValue, expectedValue := value, countDown; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapIteratorBegin(t *testing.T) {
	m := NewWith(utils.IntComparator, utils.StringComparator)
	it := m.Iterator()
	it.Begin()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestMapTreeIteratorEnd(t *testing.T) {
	m := NewWith(utils.IntComparator, utils.StringComparator)
	it := m.Iterator()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it.End()
	it.Prev()
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestMapIteratorFirst(t *testing.T) {
	m := NewWith(utils.IntComparator, utils.StringComparator)
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it := m.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestMapIteratorLast(t *testing.T) {
	m := NewWith(utils.IntComparator, utils.StringComparator)
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it := m.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestMapSerialization(t *testing.T) {
	for i := 0; i < 10; i++ {
		original := NewWith(utils.StringComparator, utils.StringComparator)
		original.Put("d", "4")
		original.Put("e", "5")
		original.Put("c", "3")
		original.Put("b", "2")
		original.Put("a", "1")

		assertSerialization(original, "A", t)

		serialized, err := original.ToJSON()
		if err != nil {
			t.Errorf("Got error %v", err)
		}
		assertSerialization(original, "B", t)

		deserialized := NewWith(utils.StringComparator, utils.StringComparator)
		err = deserialized.FromJSON(serialized)
		if err != nil {
			t.Errorf("Got error %v", err)
		}
		assertSerialization(deserialized, "C", t)
	}
}

//noinspection GoBoolExpressions
func assertSerialization(m *Map, txt string, t *testing.T) {
	if actualValue := m.Keys(); false ||
		actualValue[0].(string) != "a" ||
		actualValue[1].(string) != "b" ||
		actualValue[2].(string) != "c" ||
		actualValue[3].(string) != "d" ||
		actualValue[4].(string) != "e" {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, "[a,b,c,d,e]")
	}
	if actualValue := m.Values(); false ||
		actualValue[0].(string) != "1" ||
		actualValue[1].(string) != "2" ||
		actualValue[2].(string) != "3" ||
		actualValue[3].(string) != "4" ||
		actualValue[4].(string) != "5" {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, "[1,2,3,4,5]")
	}
	if actualValue, expectedValue := m.Size(), 5; actualValue != expectedValue {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, expectedValue)
	}
}

func benchmarkGet(b *testing.B, m *Map, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *Map, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, n)
		}
	}
}

func benchmarkRemove(b *testing.B, m *Map, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n)
		}
	}
}

func BenchmarkTreeBidiMapGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeBidiMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeBidiMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeBidiMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeBidiMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWithIntComparators()
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeBidiMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeBidiMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeBidiMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeBidiMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeBidiMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeBidiMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeBidiMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWithIntComparators()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
