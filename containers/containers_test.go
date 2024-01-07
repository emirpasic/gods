// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// All data structures must implement the container structure

package containers

import (
	"cmp"
	"fmt"
	"strings"
	"testing"
)

// For testing purposes
type ContainerTest[T any] struct {
	values []T
}

func (container ContainerTest[T]) Empty() bool {
	return len(container.values) == 0
}

func (container ContainerTest[T]) Size() int {
	return len(container.values)
}

func (container ContainerTest[T]) Clear() {
	container.values = []T{}
}

func (container ContainerTest[T]) Values() []T {
	return container.values
}

func (container ContainerTest[T]) String() string {
	str := "ContainerTest\n"
	var values []string
	for _, value := range container.values {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func TestGetSortedValuesInts(t *testing.T) {
	container := ContainerTest[int]{}
	GetSortedValues(container)
	container.values = []int{5, 1, 3, 2, 4}
	values := GetSortedValues(container)
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}

type NotInt struct {
	i int
}

func TestGetSortedValuesNotInts(t *testing.T) {
	container := ContainerTest[NotInt]{}
	GetSortedValuesFunc(container, func(x, y NotInt) int {
		return cmp.Compare(x.i, y.i)
	})
	container.values = []NotInt{{5}, {1}, {3}, {2}, {4}}
	values := GetSortedValuesFunc(container, func(x, y NotInt) int {
		return cmp.Compare(x.i, y.i)
	})
	for i := 1; i < container.Size(); i++ {
		if values[i-1].i > values[i].i {
			t.Errorf("Not sorted!")
		}
	}
}
