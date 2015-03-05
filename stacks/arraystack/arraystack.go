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

// Implementation of stack using a slice.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29

package arraystack

import (
	"fmt"
	"github.com/emirpasic/gods/stacks"
	"strings"
)

func assertInterfaceImplementation() {
	var _ stacks.Interface = (*Stack)(nil)
}

type Stack struct {
	items []interface{}
	top   int
}

// Instantiates a new empty stack
func New() *Stack {
	return &Stack{top: -1}
}

// Pushes a value onto the top of the stack
func (stack *Stack) Push(value interface{}) {
	// Increase when capacity is reached by a factor of 1.5 and add one so it grows when size is zero
	if stack.top+1 >= cap(stack.items) {
		currentSize := len(stack.items)
		sizeIncrease := int(1.5*float32(currentSize) + 1.0)
		newSize := currentSize + sizeIncrease
		newItems := make([]interface{}, newSize, newSize)
		copy(newItems, stack.items)
		stack.items = newItems
	}
	stack.top += 1
	stack.items[stack.top] = value
}

// Pops (removes) top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *Stack) Pop() (value interface{}, ok bool) {
	if stack.top >= 0 {
		value, ok = stack.items[stack.top], true
		// TODO shrink slice at some point
		stack.top -= 1
		return
	}
	return nil, false
}

// Returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *Stack) Peek() (value interface{}, ok bool) {
	if stack.top >= 0 {
		return stack.items[stack.top], true
	}
	return nil, false
}

// Returns true if stack does not contain any elements.
func (stack *Stack) Empty() bool {
	return stack.Size() == 0
}

// Returns number of elements within the stack.
func (stack *Stack) Size() int {
	return stack.top + 1
}

// Removes all elements from the stack.
func (stack *Stack) Clear() {
	stack.top = -1
	stack.items = []interface{}{}
}

func (stack *Stack) String() string {
	str := "ArrayStack\n"
	values := []string{}
	for _, value := range stack.items {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
