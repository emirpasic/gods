/*
Copyright (c) Emir Pasic, All rights reserved.

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 3.0 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library. See the file LICENSE included
with this distribution for more information.
*/

// Implementation of stack using linked list.
// Used by red-black tree during in-order traversal.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29

package linkedliststack

import (
	"fmt"
	"github.com/emirpasic/gods/stacks"
	"strings"
)

func assertInterfaceImplementation() {
	var _ stacks.Interface = (*Stack)(nil)
}

type Stack struct {
	top  *element
	size int
}

type element struct {
	value interface{}
	next  *element
}

// Instantiates a new empty stack
func New() *Stack {
	return &Stack{}
}

// Pushes a value onto the top of the stack
func (stack *Stack) Push(value interface{}) {
	stack.top = &element{value, stack.top}
	stack.size += 1
}

// Pops (removes) top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *Stack) Pop() (value interface{}, ok bool) {
	if stack.size > 0 {
		value, stack.top = stack.top.value, stack.top.next
		stack.size -= 1
		return value, true
	}
	return nil, false
}

// Returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *Stack) Peek() (value interface{}, ok bool) {
	if stack.size > 0 {
		return stack.top.value, true
	}
	return nil, false
}

// Returns true if stack does not contain any elements.
func (stack *Stack) Empty() bool {
	return stack.size == 0
}

// Returns number of elements within the stack.
func (stack *Stack) Size() int {
	return stack.size
}

// Removes all elements from the stack.
func (stack *Stack) Clear() {
	stack.top = nil
	stack.size = 0
}

func (stack *Stack) String() string {
	str := "LinkedListStack\n"
	element := stack.top
	elementsValues := []string{}
	for element != nil {
		elementsValues = append(elementsValues, fmt.Sprintf("%v", element.value))
		element = element.next
	}
	for i, j := 0, len(elementsValues)-1; i < j; i, j = i+1, j-1 {
		elementsValues[i], elementsValues[j] = elementsValues[j], elementsValues[i]
	}
	str += strings.Join(elementsValues, ", ")
	return str
}
