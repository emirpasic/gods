// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package containers provides core interfaces and functions for data structures.
//
// Container is the base interface for all data structures to implement.
//
// Iterators provide stateful iterators.
//
// Enumerable provides Ruby inspired (each, select, map, find, any?, etc.) container functions.
//
// Serialization provides serializers (marshalers) and deserializers (unmarshalers).
package containers

import (
	"cmp"
	"slices"

	"github.com/emirpasic/gods/v2/utils"
)

// Container is base interface that all data structures implement.
type Container[T any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}

// GetSortedValues returns sorted container's elements with respect to the passed comparator.
// Does not affect the ordering of elements within the container.
func GetSortedValues[T cmp.Ordered](container Container[T]) []T {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	slices.Sort(values)
	return values
}

// GetSortedValuesFunc is the equivalent of GetSortedValues for containers of values that are not ordered.
func GetSortedValuesFunc[T any](container Container[T], comparator utils.Comparator[T]) []T {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	slices.SortFunc(values, comparator)
	return values
}
