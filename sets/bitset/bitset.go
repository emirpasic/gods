// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bitset implements a set backed by a bit array.
//
// Elements must be non-negative integers. The backing array grows automatically
// to accommodate new elements.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/Bit_array
package bitset

import (
	"fmt"
	"math/bits"
	"strings"

	"github.com/emirpasic/gods/v2/sets"
)

// Assert Set implementation
var _ sets.Set[int] = (*Set)(nil)

// Set holds elements as bits in an array of uint64 words.
type Set struct {
	b    []uint64
	size int
}

// New instantiates a new empty set and adds the passed values, if any, to the set.
func New(values ...int) *Set {
	set := &Set{}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// NewWithCapacity instantiates a new empty set pre-allocated to hold elements
// in the range [0, universe) and adds the passed values, if any, to the set.
func NewWithCapacity(universe int, values ...int) *Set {
	n := 0
	if universe > 0 {
		n = (universe-1)>>6 + 1
	}
	set := &Set{
		b: make([]uint64, n),
	}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// grow ensures the backing array can accommodate element x.
func (set *Set) grow(x int) {
	needed := x>>6 + 1
	if needed <= len(set.b) {
		return
	}
	expanded := make([]uint64, needed)
	copy(expanded, set.b)
	set.b = expanded
}

// Add adds the items (one or more) to the set.
// Panics if any element is negative.
func (set *Set) Add(items ...int) {
	for _, item := range items {
		if item < 0 {
			panic(fmt.Sprintf("bitset: negative element %d", item))
		}
		set.grow(item)
		w, mask := item>>6, uint64(1)<<uint(item&63)
		if set.b[w]&mask == 0 {
			set.b[w] |= mask
			set.size++
		}
	}
}

// Remove removes the items (one or more) from the set.
// Panics if any element is negative.
func (set *Set) Remove(items ...int) {
	for _, item := range items {
		if item < 0 {
			panic(fmt.Sprintf("bitset: negative element %d", item))
		}
		w := item >> 6
		if w >= len(set.b) {
			continue
		}
		mask := uint64(1) << uint(item&63)
		if set.b[w]&mask != 0 {
			set.b[w] &^= mask
			set.size--
		}
	}
}

// Contains checks if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
// Panics if any element is negative.
func (set *Set) Contains(items ...int) bool {
	for _, item := range items {
		if item < 0 {
			panic(fmt.Sprintf("bitset: negative element %d", item))
		}
		w := item >> 6
		if w >= len(set.b) {
			return false
		}
		if set.b[w]&(uint64(1)<<uint(item&63)) == 0 {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set) Empty() bool {
	return set.size == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return set.size
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.b = nil
	set.size = 0
}

// Values returns all items in the set.
// Elements are returned in ascending order.
func (set *Set) Values() []int {
	values := make([]int, 0, set.size)
	for i, word := range set.b {
		if word == 0 {
			continue
		}
		base := i << 6 // i * 64
		for word != 0 {
			bit := bits.TrailingZeros64(word)
			values = append(values, base+bit)
			word &= word - 1 // clear lowest set bit
		}
	}
	return values
}

// String returns a string representation of container.
func (set *Set) String() string {
	str := "BitSet\n"
	items := []string{}
	for _, v := range set.Values() {
		items = append(items, fmt.Sprintf("%v", v))
	}
	str += strings.Join(items, ", ")
	return str
}

// Intersection returns the intersection between two sets.
// The new set consists of all elements that are both in "set" and "another".
// Ref: https://en.wikipedia.org/wiki/Intersection_(set_theory)
func (set *Set) Intersection(another *Set) *Set {
	result := &Set{}
	minLen := min(len(set.b), len(another.b))
	if minLen == 0 {
		return result
	}
	result.b = make([]uint64, minLen)
	for i := 0; i < minLen; i++ {
		result.b[i] = set.b[i] & another.b[i]
		result.size += bits.OnesCount64(result.b[i])
	}
	return result
}

// Union returns the union of two sets.
// The new set consists of all elements that are in "set" or "another" (possibly both).
// Ref: https://en.wikipedia.org/wiki/Union_(set_theory)
func (set *Set) Union(another *Set) *Set {
	result := &Set{}
	longer, shorter := set.b, another.b
	if len(shorter) > len(longer) {
		longer, shorter = shorter, longer
	}
	if len(longer) == 0 {
		return result
	}
	result.b = make([]uint64, len(longer))
	copy(result.b, longer)
	for i, w := range shorter {
		result.b[i] |= w
	}
	for _, w := range result.b {
		result.size += bits.OnesCount64(w)
	}
	return result
}

// Difference returns the difference between two sets.
// The new set consists of all elements that are in "set" but not in "another".
// Ref: https://proofwiki.org/wiki/Definition:Set_Difference
func (set *Set) Difference(another *Set) *Set {
	result := &Set{}
	if len(set.b) == 0 {
		return result
	}
	result.b = make([]uint64, len(set.b))
	minLen := min(len(set.b), len(another.b))
	for i := 0; i < minLen; i++ {
		result.b[i] = set.b[i] &^ another.b[i]
	}
	copy(result.b[minLen:], set.b[minLen:])
	for _, w := range result.b {
		result.size += bits.OnesCount64(w)
	}
	return result
}
