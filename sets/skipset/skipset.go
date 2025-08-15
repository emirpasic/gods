// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skipset

import (
	"cmp"
	"fmt"
	"math/bits"
	"math/rand/v2"
	"strings"

	"github.com/emirpasic/gods/v2/sets"
	"github.com/emirpasic/gods/v2/utils"
)

// Assert Set implementation
var _ sets.Set[int] = (*Set[int])(nil)

const defaultSkipSetHeight = 64

type (
	// Set holds a sentinel element and the size of the set.
	// Elements are kept in sorted order and duplicates are not allowed.
	Set[T comparable] struct {
		sentinel   *element[T]
		size       int
		comparator utils.Comparator[T]
	}

	element[T comparable] struct {
		value T
		nexts []*element[T]
		prev  *element[T]
	}
)

// New instantiates a new set with the default comparator and adds the passed values, if any, to the set
func New[T cmp.Ordered](ts ...T) *Set[T] {
	return NewWith(cmp.Compare[T], ts...)
}

// NewWith instantiates a new set with the provided comparator and adds the passed values, if any, to the set
func NewWith[T comparable](comparator utils.Comparator[T], ts ...T) *Set[T] {
	var t T
	elem := &element[T]{
		value: t,
		nexts: make([]*element[T], defaultSkipSetHeight),
	}
	for i := range elem.nexts {
		elem.nexts[i] = elem
	}
	s := &Set[T]{
		sentinel:   elem,
		comparator: comparator,
	}
	s.Add(ts...)
	return s
}

// Add adds the items (one or more) to the set. If an item already exists, it is not added again.
func (s *Set[T]) Add(values ...T) {
	for _, value := range values {
		if s.Contains(value) {
			continue // Skip duplicates
		}

		elem := &element[T]{
			value: value,
			nexts: make([]*element[T], bits.TrailingZeros64(uint64(rand.Int64())<<1)),
		}

		prev := s.sentinel
		for height := len(prev.nexts) - 1; height >= 0; height-- {
			for prev.nexts[height] != s.sentinel && s.comparator(prev.nexts[height].value, value) < 0 {
				prev = prev.nexts[height]
			}
			if height < len(elem.nexts) {
				elem.nexts[height] = prev.nexts[height]
				prev.nexts[height] = elem
			}
		}
		elem.prev = prev
		elem.nexts[0].prev = elem
		s.size++
	}
}

// Remove removes the items (one or more) from the set.
func (s *Set[T]) Remove(values ...T) {
	for _, value := range values {
		for height, prev := len(s.sentinel.nexts)-1, s.sentinel; height >= 0; height-- {
			for prev.nexts[height] != s.sentinel && s.comparator(prev.nexts[height].value, value) < 0 {
				prev = prev.nexts[height]
			}

			// If this level points to the element we want to remove, update the pointer
			if prev.nexts[height] != s.sentinel && s.comparator(prev.nexts[height].value, value) == 0 {
				prev.nexts[height] = prev.nexts[height].nexts[height]

				// Update prev pointer on level 0
				if height == 0 {
					prev.nexts[height].nexts[0].prev = prev
					s.size--
				}
			}
		}
	}
}

// Contains checks if values (one or more) are present in the set.
// All values have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (s *Set[T]) Contains(values ...T) bool {
	for _, value := range values {
		// Search for the value in the skip list
		current := s.sentinel
		for level := len(s.sentinel.nexts) - 1; level >= 0; level-- {
			for current.nexts[level] != s.sentinel && s.comparator(current.nexts[level].value, value) < 0 {
				current = current.nexts[level]
			}
		}
		current = current.nexts[0]
		if current == s.sentinel || s.comparator(current.value, value) != 0 {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (s *Set[T]) Empty() bool {
	return s.size == 0
}

// Size returns number of elements within the set.
func (s *Set[T]) Size() int {
	return s.size
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	s.size = 0
	for i := range s.sentinel.nexts {
		s.sentinel.nexts[i] = s.sentinel
	}
	s.sentinel.prev = s.sentinel
}

// Values returns all elements in the set as a slice in sorted order.
func (s *Set[T]) Values() []T {
	values := make([]T, s.size)
	elem := s.sentinel
	for i := 0; i < s.size; i++ {
		elem = elem.nexts[0]
		values[i] = elem.value
	}
	return values
}

// String returns a string representation of container
func (s *Set[T]) String() string {
	str := "SkipSet\n"
	values := make([]string, s.size)

	elem := s.sentinel
	for i := 0; i < s.size; i++ {
		elem = elem.nexts[0]
		values[i] = fmt.Sprintf("%v", elem.value)
	}

	str += strings.Join(values, ", ")
	return str
}

// Union returns a new set containing all elements from both sets
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	unionSet := NewWith(s.comparator)
	unionSet.Add(s.Values()...)
	unionSet.Add(other.Values()...)
	return unionSet
}

// Intersection returns a new set containing elements present in both sets
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersection := NewWith(s.comparator)

	// Iterate through the smaller set for efficiency
	smaller, larger := s, other
	if other.Size() < s.Size() {
		smaller, larger = other, s
	}

	for _, value := range smaller.Values() {
		if larger.Contains(value) {
			intersection.Add(value)
		}
	}
	return intersection
}

// Difference returns a new set containing elements present in this set but not in the other set
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	smaller, larger := s, other
	if other.Size() < s.Size() {
		smaller, larger = other, s
	}

	difference := NewWith(s.comparator)
	for _, value := range smaller.Values() {
		if !larger.Contains(value) {
			difference.Add(value)
		}
	}
	return difference
}
