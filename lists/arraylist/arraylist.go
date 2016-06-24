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

// Implementation of list using a slice.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/List_%28abstract_data_type%29

package arraylist

import (
	"fmt"
	"github.com/emirpasic/gods/containers"
	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/utils"
	"strings"
)

func assertInterfaceImplementation() {
	var _ lists.List = (*List)(nil)
	var _ containers.EnumerableWithIndex = (*List)(nil)
	var _ containers.IteratorWithIndex = (*Iterator)(nil)
}

type List struct {
	elements []interface{}
	size     int
}

const (
	GROWTH_FACTOR = float32(2.0)  // growth by 100%
	SHRINK_FACTOR = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// Instantiates a new empty list
func New() *List {
	return &List{}
}

// Appends a value at the end of the list
func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size += 1
	}
}

// Returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	return list.elements[index], true
}

// Removes one or more elements from the list with the supplied indices.
func (list *List) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	list.elements[index] = nil                                    // cleanup reference
	copy(list.elements[index:], list.elements[index+1:list.size]) // shift to the left by one (slow operation, need ways to optimize this)
	list.size -= 1

	list.shrink()
}

// Check if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List) Contains(values ...interface{}) bool {

	for _, searchValue := range values {
		found := false
		for _, element := range list.elements {
			if element == searchValue {
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

// Returns all elements in the list.
func (list *List) Values() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// Returns true if list does not contain any elements.
func (list *List) Empty() bool {
	return list.size == 0
}

// Returns number of elements within the list.
func (list *List) Size() int {
	return list.size
}

// Removes all elements from the list.
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

// Sorts values (in-place) using timsort.
func (list *List) Sort(comparator utils.Comparator) {
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

// Swaps the two values at the specified positions.
func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// Inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List) Insert(index int, values ...interface{}) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	// Shift old to right
	for i := list.size - 1; i >= index+l; i-- {
		list.elements[i] = list.elements[i-l]
	}
	// Insert new
	for i, value := range values {
		list.elements[index+i] = value
	}
}

type Iterator struct {
	list  *List
	index int
}

// Returns a stateful iterator whose values can be fetched by an index.
func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1}
}

// Moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) Next() bool {
	iterator.index += 1
	return iterator.list.withinRange(iterator.index)
}

// Returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Value() interface{} {
	return iterator.list.elements[iterator.index]
}

// Returns the current element's index.
// Does not modify the state of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}

// Calls the given function once for each element, passing that element's index and value.
func (list *List) Each(f func(index int, value interface{})) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func (list *List) Map(f func(index int, value interface{}) interface{}) *List {
	newList := &List{}
	iterator := list.Iterator()
	for iterator.Next() {
		newList.Add(f(iterator.Index(), iterator.Value()))
	}
	return newList
}

// Returns a new container containing all elements for which the given function returns a true value.
func (list *List) Select(f func(index int, value interface{}) bool) *List {
	newList := &List{}
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newList.Add(iterator.Value())
		}
	}
	return newList
}

// Passes each element of the collection to the given function and
// returns true if the function ever returns true for any element.
func (list *List) Any(f func(index int, value interface{}) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

// Passes each element of the collection to the given function and
// returns true if the function returns true for all elements.
func (list *List) All(f func(index int, value interface{}) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (list *List) Find(f func(index int, value interface{}) bool) (int, interface{}) {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, nil
}

func (list *List) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is withing bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *List) growBy(n int) {
	// When capacity is reached, grow by a factor of GROWTH_FACTOR and add number of elements
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(GROWTH_FACTOR * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is SHRINK_FACTOR percent of current capacity
func (list *List) shrink() {
	if SHRINK_FACTOR == 0.0 {
		return
	}
	// Shrink when size is at SHRINK_FACTOR * capacity
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*SHRINK_FACTOR) {
		list.resize(list.size)
	}
}
