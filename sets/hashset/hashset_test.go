// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashset_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/emirpasic/gods/v2/sets/hashset"
)

func TestSetNew(t *testing.T) {
	set := hashset.New(2, 1)

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
	set := hashset.New[int]()
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
	set := hashset.New[int]()
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
	set := hashset.New[int]()
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

func TestSetSerialization(t *testing.T) {
	set := hashset.New[string]()
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
	c := hashset.New[int]()
	c.Add(1)
	if !strings.HasPrefix(c.String(), "HashSet") {
		t.Errorf("String should start with container name")
	}
}

func TestSetIntersection(t *testing.T) {
	set := hashset.New[string]()
	another := hashset.New[string]()

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
	set := hashset.New[string]()
	another := hashset.New[string]()

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
	set := hashset.New[string]()
	another := hashset.New[string]()

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

func benchmarkContains(b *testing.B, set *hashset.Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Contains(n)
		}
	}
}

func benchmarkAdd(b *testing.B, set *hashset.Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, set *hashset.Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Remove(n)
		}
	}
}

func BenchmarkHashSetContains100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetContains1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetContains10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetContains100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := hashset.New[int]()
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkHashSetRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkHashSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkHashSetRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := hashset.New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}
