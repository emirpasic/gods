// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arraylist implements the array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package arraylist

import (
	"fmt"
	"slices"
	"strings"

	"github.com/emirpasic/gods/v2/lists"
	"github.com/emirpasic/gods/v2/utils"
)

// Assert List implementation
var _ lists.List[int] = (*List[int])(nil)

// List holds the elements in a slice
type List[T comparable] struct {
	elements []T
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New instantiates a new list and adds the passed values, if any, to the list
func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends a value at the end of the list
func (list *List[T]) Add(values ...T) {
	l := len(list.elements)
	list.growBy(len(values))
	for i := range values {
		list.elements[l+i] = values[i]
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List[T]) Get(index int) (T, bool) {

	if !list.withinRange(index) {
		var t T
		return t, false
	}

	return list.elements[index], true
}

// Remove removes the element at the given index from the list.
func (list *List[T]) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	list.elements = slices.Delete(list.elements, index, index+1)
	list.shrink()
}

// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List[T]) Contains(values ...T) bool {
	for _, searchValue := range values {
		if !slices.Contains(list.elements, searchValue) {
			return false
		}
	}
	return true
}

// Values returns all elements in the list.
func (list *List[T]) Values() []T {
	return slices.Clone(list.elements)
}

// IndexOf returns index of provided element
func (list *List[T]) IndexOf(value T) int {
	return slices.Index(list.elements, value)
}

// Empty returns true if list does not contain any elements.
func (list *List[T]) Empty() bool {
	return len(list.elements) == 0
}

// Size returns number of elements within the list.
func (list *List[T]) Size() int {
	return len(list.elements)
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	clear(list.elements[:cap(list.elements)])
	list.elements = list.elements[:0]
}

// Sort sorts values (in-place) using.
func (list *List[T]) Sort(comparator utils.Comparator[T]) {
	if len(list.elements) < 2 {
		return
	}
	slices.SortFunc(list.elements, comparator)
}

// Swap swaps the two values at the specified positions.
func (list *List[T]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Insert(index int, values ...T) {
	if !list.withinRange(index) {
		// Append
		if index == len(list.elements) {
			list.Add(values...)
		}
		return
	}

	l := len(list.elements)
	list.growBy(len(values))
	list.elements = slices.Insert(list.elements[:l], index, values...)
}

// Set the value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Set(index int, value T) {

	if !list.withinRange(index) {
		// Append
		if index == len(list.elements) {
			list.Add(value)
		}
		return
	}

	list.elements[index] = value
}

// String returns a string representation of container
func (list *List[T]) String() string {
	str := "ArrayList\n"
	values := make([]string, 0, len(list.elements))
	for _, value := range list.elements {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < len(list.elements)
}

func (list *List[T]) resize(len, cap int) {
	newElements := make([]T, len, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *List[T]) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(list.elements)

	if newLength := len(list.elements) + n; newLength >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newLength, newCapacity)
	} else {
		list.elements = list.elements[:newLength]
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (list *List[T]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(list.elements)
	if len(list.elements) <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(len(list.elements), len(list.elements))
	}
}
