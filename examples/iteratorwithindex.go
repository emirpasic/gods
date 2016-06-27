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

// IteratorWithIndexExample to demonstrate basic usage of IteratorWithIndex
func IteratorWithIndexExample() {
	set := treeset.NewWithStringComparator()
	set.Add("a", "b", "c")
	it := set.Iterator()

	fmt.Print("\nForward iteration\n")
	for it.Next() {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]") // [0:a][1:b][2:c]
	}

	fmt.Print("\nForward iteration (again)\n")
	for it.Begin(); it.Next(); {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]") // [0:a][1:b][2:c]
	}

	fmt.Print("\nBackward iteration\n")
	for it.Prev() {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]") // [2:c][1:b][0:a]
	}

	fmt.Print("\nBackward iteration (again)\n")
	for it.End(); it.Prev(); {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]") // [2:c][1:b][0:a]
	}

	if it.First() {
		fmt.Print("\nFirst index: ", it.Index()) // First index: 0
		fmt.Print("\nFirst value: ", it.Value()) // First value: a
	}

	if it.Last() {
		fmt.Print("\nLast index: ", it.Index()) // Last index: 3
		fmt.Print("\nLast value: ", it.Value()) // Last value: c
	}
}
