// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashset_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/emirpasic/gods/v2/sets/linkedhashset"
)

func TestSetNew(t *testing.T) {
	set := linkedhashset.New(2, 1)

	if actualValue := set.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue := set.Contains(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(3); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetAdd(t *testing.T) {
	set := linkedhashset.New[int]()
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
}

func TestSetContains(t *testing.T) {
	set := linkedhashset.New[int]()
	set.Add(3, 1, 2)
	set.Add(2, 3)
	set.Add()
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
	set := linkedhashset.New[int]()
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
	set := linkedhashset.New[string]()
	set.Add("c", "a", "b")
	set.Each(func(index int, value string) {
		switch index {
		case 0:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	})
}

func TestSetMap(t *testing.T) {
	set := linkedhashset.New[string]()
	set.Add("c", "a", "b")
	mappedSet := set.Map(func(index int, value string) string {
		return "mapped: " + value
	})
	if actualValue, expectedValue := mappedSet.Contains("mapped: c", "mapped: b", "mapped: a"), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := mappedSet.Contains("mapped: c", "mapped: b", "mapped: x"), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if mappedSet.Size() != 3 {
		t.Errorf("Got %v expected %v", mappedSet.Size(), 3)
	}
}

func TestSetSelect(t *testing.T) {
	set := linkedhashset.New[string]()
	set.Add("c", "a", "b")
	selectedSet := set.Select(func(index int, value string) bool {
		return value >= "a" && value <= "b"
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
	set := linkedhashset.New[string]()
	set.Add("c", "a", "b")
	any := set.Any(func(index int, value string) bool {
		return value == "c"
	})
	if any != true {
		t.Errorf("Got %v expected %v", any, true)
	}
	any = set.Any(func(index int, value string) bool {
		return value == "x"
	})
	if any != false {
		t.Errorf("Got %v expected %v", any, false)
	}
}

func TestSetAll(t *testing.T) {
	set := linkedhashset.New[string]()
	set.Add("c", "a", "b")
	all := set.All(func(index int, value string) bool {
		return value >= "a" && value <= "c"
	})
	if all != true {
		t.Errorf("Got %v expected %v", all, true)
	}
	all = set.All(func(index int, value string) bool {
		return value >= "a" && value <= "b"
	})
	if all != false {
		t.Errorf("Got %v expected %v", all, false)
	}
}

func TestSetFind(t *testing.T) {
	set := linkedhashset.New[string]()
	set.Add("c", "a", "b")
	foundIndex, foundValue := set.Find(func(index int, value string) bool {
		return value == "c"
	})
	if foundValue != "c" || foundIndex != 0 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, "c", 0)
	}
	foundIndex, foundValue = set.Find(func(index int, value string) bool {
		return value == "x"
	})
	if foundValue != "" || foundIndex != -1 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, nil, nil)
	}
}

func TestSetChaining(t *testing.T) {
	set := linkedhashset.New[string]()
	set.Add("c", "a", "b")
}

func TestSetIteratorPrevOnEmpty(t *testing.T) {
	set := linkedhashset.New[string]()
	it := set.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty set")
	}
}

func TestSetIteratorNext(t *testing.T) {
	set := linkedhashset.New[string]()
	set.Add("c", "a", "b")
	it := set.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
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
	set := linkedhashset.New[string]()
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
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
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
	set := linkedhashset.New[string]()
	it := set.Iterator()
	it.Begin()
	set.Add("a", "b", "c")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestSetIteratorEnd(t *testing.T) {
	set := linkedhashset.New[string]()
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
	set := linkedhashset.New[string]()
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
	set := linkedhashset.New[string]()
	set.Add("a", "b", "c")
	it := set.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 2 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 3, "c")
	}
}

func TestSetIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// NextTo (empty)
	{
		set := linkedhashset.New[string]()
		it := set.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty set")
		}
	}

	// NextTo (not found)
	{
		set := linkedhashset.New[string]()
		set.Add("xx", "yy")
		it := set.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty set")
		}
	}

	// NextTo (found)
	{
		set := linkedhashset.New[string]()
		set.Add("aa", "bb", "cc")
		it := set.Iterator()
		it.Begin()
		if !it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty set")
		}
		if index, value := it.Index(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Next() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Index(), it.Value(); index != 2 || value != "cc" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "cc")
		}
		if it.Next() {
			t.Errorf("Should not go past last element")
		}
	}
}

func TestSetIteratorPrevTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// PrevTo (empty)
	{
		set := linkedhashset.New[string]()
		it := set.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty set")
		}
	}

	// PrevTo (not found)
	{
		set := linkedhashset.New[string]()
		set.Add("xx", "yy")
		it := set.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty set")
		}
	}

	// PrevTo (found)
	{
		set := linkedhashset.New[string]()
		set.Add("aa", "bb", "cc")
		it := set.Iterator()
		it.End()
		if !it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty set")
		}
		if index, value := it.Index(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Prev() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Index(), it.Value(); index != 0 || value != "aa" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "aa")
		}
		if it.Prev() {
			t.Errorf("Should not go before first element")
		}
	}
}

func TestSetSerialization(t *testing.T) {
	set := linkedhashset.New[string]()
	set.Add("a", "b", "c")

	var err error
	assert := func() {
		if actualValue, expectedValue := set.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue := set.Contains("a", "b", "c"); actualValue != true {
			t.Errorf("Got %v expected %v", actualValue, true)
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	bytes, err := set.ToJSON()
	assert()

	err = set.FromJSON(bytes)
	assert()

	bytes, err = json.Marshal([]interface{}{"a", "b", "c", set})
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	err = json.Unmarshal([]byte(`["a","b","c"]`), &set)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	assert()
}

func TestSetString(t *testing.T) {
	c := linkedhashset.New[int]()
	c.Add(1)
	if !strings.HasPrefix(c.String(), "LinkedHashSet") {
		t.Errorf("String should start with container name")
	}
}

func TestSetIntersection(t *testing.T) {
	set := linkedhashset.New[string]()
	another := linkedhashset.New[string]()

	intersection := set.Intersection(another)
	if actualValue, expectedValue := intersection.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	intersection = set.Intersection(another)

	if actualValue, expectedValue := intersection.Size(), 2; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := intersection.Contains("c", "d"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetUnion(t *testing.T) {
	set := linkedhashset.New[string]()
	another := linkedhashset.New[string]()

	union := set.Union(another)
	if actualValue, expectedValue := union.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	union = set.Union(another)

	if actualValue, expectedValue := union.Size(), 6; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := union.Contains("a", "b", "c", "d", "e", "f"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetDifference(t *testing.T) {
	set := linkedhashset.New[string]()
	another := linkedhashset.New[string]()

	difference := set.Difference(another)
	if actualValue, expectedValue := difference.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	difference = set.Difference(another)

	if actualValue, expectedValue := difference.Size(), 2; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := difference.Contains("a", "b"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func benchmarkContains(b *testing.B, set *linkedhashset.Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Contains(n)
		}
	}
}

func benchmarkAdd(b *testing.B, set *linkedhashset.Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, set *linkedhashset.Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Remove(n)
		}
	}
}

func BenchmarkHashSetContains100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetContains1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetContains10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetContains100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := linkedhashset.New[int]()
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkHashSetRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkHashSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkHashSetRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := linkedhashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}
