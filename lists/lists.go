// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lists provides an abstract List interface.
//
// In computer science, a list or sequence is an abstract data type that represents an ordered sequence of values, where the same value may occur more than once. An instance of a list is a computer representation of the mathematical concept of a finite sequence; the (potentially) infinite analog of a list is a stream.  Lists are a basic example of containers, as they contain other values. If the same value occurs multiple times, each occurrence is considered a distinct item.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package lists

import (
	"github.com/emirpasic/gods/v2/containers"
	"github.com/emirpasic/gods/v2/utils"
)

// List interface that all lists implement
type List[T comparable] interface {
	Get(index int) (T, bool)
	Remove(index int)
	Add(values ...T)
	Contains(values ...T) bool
	Sort(comparator utils.Comparator[T])
	Swap(index1, index2 int)
	Insert(index int, values ...T)
	Set(index int, value T)

	containers.Container[T]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
	// String() string
}
