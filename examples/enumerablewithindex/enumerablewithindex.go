// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/emirpasic/gods/v2/sets/treeset"
)

func printSet(txt string, set *treeset.Set[int]) {
	fmt.Print(txt, "[ ")
	set.Each(func(index int, value int) {
		fmt.Print(value, " ")
	})
	fmt.Println("]")
}

// EnumerableWithIndexExample to demonstrate basic usage of EnumerableWithIndex
func main() {
	set := treeset.New[int]()
	set.Add(2, 3, 4, 2, 5, 6, 7, 8)
	printSet("Initial", set) // [ 2 3 4 5 6 7 8 ]

	even := set.Select(func(index int, value int) bool {
		return value%2 == 0
	})
	printSet("Even numbers", even) // [ 2 4 6 8 ]

	foundIndex, foundValue := set.Find(func(index int, value int) bool {
		return value%2 == 0 && value%3 == 0
	})
	if foundIndex != -1 {
		fmt.Println("Number divisible by 2 and 3 found is", foundValue, "at index", foundIndex) // value: 6, index: 4
	}

	square := set.Map(func(index int, value int) int {
		return value * value
	})
	printSet("Numbers squared", square) // [ 4 9 16 25 36 49 64 ]

	bigger := set.Any(func(index int, value int) bool {
		return value > 5
	})
	fmt.Println("Set contains a number bigger than 5 is ", bigger) // true

	positive := set.All(func(index int, value int) bool {
		return value > 0
	})
	fmt.Println("All numbers are positive is", positive) // true

	evenNumbersSquared := set.Select(func(index int, value int) bool {
		return value%2 == 0
	}).Map(func(index int, value int) int {
		return value * value
	})
	printSet("Chaining", evenNumbersSquared) // [ 4 16 36 64 ]
}
