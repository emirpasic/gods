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

package utils

import (
	"testing"
)

func TestIntComparator(t *testing.T) {

	// i1,i2,expected
	tests := [][]interface{}{
		{1, 1, 0},
		{1, 2, -1},
		{2, 1, 1},
		{11, 22, -1},
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
	}

	for _, test := range tests {
		actual := IntComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestStringComparator(t *testing.T) {

	// s1,s2,expected
	tests := [][]interface{}{
		{"a", "a", 0},
		{"a", "b", -1},
		{"b", "a", 1},
		{"aa", "aab", -1},
		{"", "", 0},
		{"a", "", 1},
		{"", "a", -1},
		{"", "aaaaaaa", -1},
	}

	for _, test := range tests {
		actual := StringComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestCustomComparator(t *testing.T) {

	type Custom struct {
		id   int
		name string
	}

	byID := func(a, b interface{}) int {
		c1 := a.(Custom)
		c2 := b.(Custom)
		switch {
		case c1.id > c2.id:
			return 1
		case c1.id < c2.id:
			return -1
		default:
			return 0
		}
	}

	// o1,o2,expected
	tests := [][]interface{}{
		{Custom{1, "a"}, Custom{1, "a"}, 0},
		{Custom{1, "a"}, Custom{2, "b"}, -1},
		{Custom{2, "b"}, Custom{1, "a"}, 1},
		{Custom{1, "a"}, Custom{1, "b"}, 0},
	}

	for _, test := range tests {
		actual := byID(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}
