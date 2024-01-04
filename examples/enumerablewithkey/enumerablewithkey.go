// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/emirpasic/gods/v2/maps/treemap"
)

func printMap(txt string, m *treemap.Map[string, int]) {
	fmt.Print(txt, " { ")
	m.Each(func(key string, value int) {
		fmt.Print(key, ":", value, " ")
	})
	fmt.Println("}")
}

// EunumerableWithKeyExample to demonstrate basic usage of EunumerableWithKey
func main() {
	m := treemap.New[string, int]()
	m.Put("g", 7)
	m.Put("f", 6)
	m.Put("e", 5)
	m.Put("d", 4)
	m.Put("c", 3)
	m.Put("b", 2)
	m.Put("a", 1)
	printMap("Initial", m) // { a:1 b:2 c:3 d:4 e:5 f:6 g:7 }

	even := m.Select(func(key string, value int) bool {
		return value%2 == 0
	})
	printMap("Elements with even values", even) // { b:2 d:4 f:6 }

	foundKey, foundValue := m.Find(func(key string, value int) bool {
		return value%2 == 0 && value%3 == 0
	})
	if foundKey != "" {
		fmt.Println("Element with value divisible by 2 and 3 found is", foundValue, "with key", foundKey) // value: 6, index: 4
	}

	square := m.Map(func(key string, value int) (string, int) {
		return key + key, value * value
	})
	printMap("Elements' values squared and letters duplicated", square) // { aa:1 bb:4 cc:9 dd:16 ee:25 ff:36 gg:49 }

	bigger := m.Any(func(key string, value int) bool {
		return value > 5
	})
	fmt.Println("Map contains element whose value is bigger than 5 is", bigger) // true

	positive := m.All(func(key string, value int) bool {
		return value > 0
	})
	fmt.Println("All map's elements have positive values is", positive) // true

	evenNumbersSquared := m.Select(func(key string, value int) bool {
		return value%2 == 0
	}).Map(func(key string, value int) (string, int) {
		return key, value * value
	})
	printMap("Chaining", evenNumbersSquared) // { b:4 d:16 f:36 }
}
