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
	"github.com/emirpasic/gods/sets/treeset"
)

func printSet(txt string, set *treeset.Set) {
	fmt.Print(txt, "[ ")
	set.Each(func(index int, value interface{}) {
		fmt.Print(value, " ")
	})
	fmt.Println("]")
}

func EnumerableWithIndexTest() {
	set := treeset.NewWithIntComparator()
	set.Add(2, 3, 4, 2, 5, 6, 7, 8)
	printSet("Initial", set) // [ 2 3 4 5 6 7 8 ]

	even := set.Select(func(index int, value interface{}) bool {
		return value.(int)%2 == 0
	})
	printSet("Even numbers", even) // [ 2 4 6 8 ]

	foundIndex, foundValue := set.Find(func(index int, value interface{}) bool {
		return value.(int)%2 == 0 && value.(int)%3 == 0
	})
	if foundIndex != -1 {
		fmt.Println("Number divisible by 2 and 3 found is", foundValue, "at index", foundIndex) // value: 6, index: 4
	}

	square := set.Map(func(index int, value interface{}) interface{} {
		return value.(int) * value.(int)
	})
	printSet("Numbers squared", square) // [ 4 9 16 25 36 49 64 ]

	bigger := set.Any(func(index int, value interface{}) bool {
		return value.(int) > 5
	})
	fmt.Println("Set contains a number bigger than 5 is ", bigger) // true

	positive := set.All(func(index int, value interface{}) bool {
		return value.(int) > 0
	})
	fmt.Println("All numbers are positive is", positive) // true

	evenNumbersSquared := set.Select(func(index int, value interface{}) bool {
		return value.(int)%2 == 0
	}).Map(func(index int, value interface{}) interface{} {
		return value.(int) * value.(int)
	})
	printSet("Chaining", evenNumbersSquared) // [ 4 16 36 64 ]
}
