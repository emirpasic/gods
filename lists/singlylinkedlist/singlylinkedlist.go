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

// Implementation of doubly linked list.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Linked_list#Singly_linked_list

package singlylinkedlist

import (
	"fmt"
	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/utils"
	"strings"
)

func assertInterfaceImplementation() {
	var _ lists.Interface = (*List)(nil)
}

type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	next  *element
}

// Instantiates a new empty list
func New() *List {
	return &List{}
}

// Appends a value (one or more) at the end of the list (same as Append())
func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

// Appends a value (one or more) at the end of the list (same as Add())
func (list *List) Append(values ...interface{}) {
	list.Add(values...)
}

// Prepends a values (or more)
func (list *List) Prepend(values ...interface{}) {
	// in reverse to keep passed order i.e. ["c","d"] -> Prepend(["a","b"]) -> ["a","b","c",d"]
	for v := len(values) - 1; v >= 0; v-- {
		newElement := &element{value: values[v], next: list.first}
		list.first = newElement
		if list.size == 0 {
			list.last = newElement
		}
		list.size++
	}
}

// Returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
	}

	return element.value, true
}

// Removes one or more elements from the list with the supplied indices.
func (list *List) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var beforeElement *element
	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
		beforeElement = element
	}

	if element == list.first {
		list.first = element.next
	}
	if element == list.last {
		list.last = beforeElement
	}
	if beforeElement != nil {
		beforeElement.next = element.next
	}

	element = nil

	list.size--
}

// Check if values (one or more) are present in the set.
// All values have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List) Contains(values ...interface{}) bool {

	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if element.value == value {
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
	values := make([]interface{}, list.size, list.size)
	for e, element := 0, list.first; element != nil; e, element = e+1, element.next {
		values[e] = element.value
	}
	return values
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
	list.first = nil
	list.last = nil
}

// Sorts values (in-place) using timsort.
func (list *List) Sort(comparator utils.Comparator) {

	if list.size < 2 {
		return
	}

	values := list.Values()
	utils.Sort(values, comparator)

	list.Clear()

	list.Add(values...)

}

// Swaps values of two elements at the given indices.
func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) && i != j {
		var element1, element2 *element
		for e, currentElement := 0, list.first; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
			switch e {
			case i:
				element1 = currentElement
			case j:
				element2 = currentElement
			}
		}
		element1.value, element2.value = element2.value, element1.value
	}
}

func (list *List) String() string {
	str := "SinglyLinkedList\n"
	values := []string{}
	for element := list.first; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is withing bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size && list.size != 0
}
