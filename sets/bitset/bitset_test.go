// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitset

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSetNew(t *testing.T) {
	set := New(2, 1)

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
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestSetNewWithCapacity(t *testing.T) {
	set := NewWithCapacity(100, 2, 1)

	if actualValue := set.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue := set.Contains(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(99); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestSetAdd(t *testing.T) {
	set := New()
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

func TestSetAddAutoGrow(t *testing.T) {
	set := NewWithCapacity(10)
	set.Add(5)
	set.Add(100) // beyond initial capacity, should grow
	if actualValue := set.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue := set.Contains(5); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(100); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetAddNegativePanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for negative element")
		}
	}()
	set := New()
	set.Add(-1)
}

func TestSetContains(t *testing.T) {
	set := New()
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

func TestSetContainsBeyondCapacity(t *testing.T) {
	set := NewWithCapacity(10)
	set.Add(5)
	// Element beyond backing array should return false, not panic
	if actualValue := set.Contains(1000); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestSetRemove(t *testing.T) {
	set := New()
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
	set.Remove(3) // remove non-existent, should be no-op
	set.Remove()
	set.Remove(2)
	if actualValue := set.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestSetRemoveBeyondCapacity(t *testing.T) {
	set := NewWithCapacity(10)
	set.Add(5)
	// Remove beyond backing array should be no-op, not panic
	set.Remove(1000)
	if actualValue := set.Size(); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestSetValues(t *testing.T) {
	set := New()
	set.Add(3, 1, 2, 100, 65)
	values := set.Values()
	if len(values) != 5 {
		t.Errorf("Got %v expected %v", len(values), 5)
	}
	// Values should be in ascending order
	expected := []int{1, 2, 3, 65, 100}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf("Got %v expected %v at index %d", v, expected[i], i)
		}
	}
}

func TestSetEmpty(t *testing.T) {
	set := New()
	if actualValue := set.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	set.Add(1)
	if actualValue := set.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	set.Remove(1)
	if actualValue := set.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetClear(t *testing.T) {
	set := New()
	set.Add(1, 2, 3)
	set.Clear()
	if actualValue := set.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := set.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetSerialization(t *testing.T) {
	set := New()
	set.Add(1, 2, 3)

	var err error
	assert := func() {
		if actualValue, expectedValue := set.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue := set.Contains(1, 2, 3); actualValue != true {
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

	bytes, err = json.Marshal([]interface{}{1, 2, 3, set})
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	err = json.Unmarshal([]byte(`[1,2,3]`), &set)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	assert()
}

func TestSetString(t *testing.T) {
	c := New()
	c.Add(1)
	if !strings.HasPrefix(c.String(), "BitSet") {
		t.Errorf("String should start with container name")
	}
}

func TestSetIntersection(t *testing.T) {
	set := New()
	another := New()

	intersection := set.Intersection(another)
	if actualValue, expectedValue := intersection.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add(1, 2, 3, 4)
	another.Add(3, 4, 5, 6)

	intersection = set.Intersection(another)

	if actualValue, expectedValue := intersection.Size(), 2; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := intersection.Contains(3, 4); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetUnion(t *testing.T) {
	set := New()
	another := New()

	union := set.Union(another)
	if actualValue, expectedValue := union.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add(1, 2, 3, 4)
	another.Add(3, 4, 5, 6)

	union = set.Union(another)

	if actualValue, expectedValue := union.Size(), 6; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := union.Contains(1, 2, 3, 4, 5, 6); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetDifference(t *testing.T) {
	set := New()
	another := New()

	difference := set.Difference(another)
	if actualValue, expectedValue := difference.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add(1, 2, 3, 4)
	another.Add(3, 4, 5, 6)

	difference = set.Difference(another)

	if actualValue, expectedValue := difference.Size(), 2; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := difference.Contains(1, 2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetIntersectionMismatchedSize(t *testing.T) {
	set := New()
	another := New()
	set.Add(1, 2, 100)
	another.Add(2, 3)

	intersection := set.Intersection(another)
	if actualValue, expectedValue := intersection.Size(), 1; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := intersection.Contains(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetUnionMismatchedSize(t *testing.T) {
	set := New()
	another := New()
	set.Add(1, 2, 100)
	another.Add(2, 3)

	union := set.Union(another)
	if actualValue, expectedValue := union.Size(), 4; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := union.Contains(1, 2, 3, 100); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetDifferenceMismatchedSize(t *testing.T) {
	set := New()
	another := New()
	set.Add(1, 2, 100)
	another.Add(2, 3)

	difference := set.Difference(another)
	if actualValue, expectedValue := difference.Size(), 2; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := difference.Contains(1, 100); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func benchmarkContains(b *testing.B, set *Set, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Contains(n)
		}
	}
}

func benchmarkAdd(b *testing.B, set *Set, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, set *Set, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Remove(n)
		}
	}
}

func BenchmarkBitSetContains100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkBitSetContains1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkBitSetContains10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkBitSetContains100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkBitSetAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New()
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkBitSetAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkBitSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkBitSetAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkBitSetRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkBitSetRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkBitSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkBitSetRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}
