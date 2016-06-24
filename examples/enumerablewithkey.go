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

package examples

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
)

func printMap(txt string, m *treemap.Map) {
	fmt.Print(txt, " { ")
	m.Each(func(key interface{}, value interface{}) {
		fmt.Print(key, ":", value, " ")
	})
	fmt.Println("}")
}

func EunumerableWithKey() {
	m := treemap.NewWithStringComparator()
	m.Put("g", 7)
	m.Put("f", 6)
	m.Put("e", 5)
	m.Put("d", 4)
	m.Put("c", 3)
	m.Put("b", 2)
	m.Put("a", 1)
	printMap("Initial", m) // { a:1 b:2 c:3 d:4 e:5 f:6 g:7 }

	even := m.Select(func(key interface{}, value interface{}) bool {
		return value.(int)%2 == 0
	})
	printMap("Elements with even values", even) // { b:2 d:4 f:6 }

	foundKey, foundValue := m.Find(func(key interface{}, value interface{}) bool {
		return value.(int)%2 == 0 && value.(int)%3 == 0
	})
	if foundKey != nil {
		fmt.Println("Element with value divisible by 2 and 3 found is", foundValue, "with key", foundKey) // value: 6, index: 4
	}

	square := m.Map(func(key interface{}, value interface{}) (interface{}, interface{}) {
		return key.(string) + key.(string), value.(int) * value.(int)
	})
	printMap("Elements' values squared and letters duplicated", square) // { aa:1 bb:4 cc:9 dd:16 ee:25 ff:36 gg:49 }

	bigger := m.Any(func(key interface{}, value interface{}) bool {
		return value.(int) > 5
	})
	fmt.Println("Map contains element whose value is bigger than 5 is", bigger) // true

	positive := m.All(func(key interface{}, value interface{}) bool {
		return value.(int) > 0
	})
	fmt.Println("All map's elements have positive values is", positive) // true

	evenNumbersSquared := m.Select(func(key interface{}, value interface{}) bool {
		return value.(int)%2 == 0
	}).Map(func(key interface{}, value interface{}) (interface{}, interface{}) {
		return key, value.(int) * value.(int)
	})
	printMap("Chaining", evenNumbersSquared) // { b:4 d:16 f:36 }
}
