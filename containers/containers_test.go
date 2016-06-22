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

// All data structures must implement the container structure

package containers

import (
	"github.com/emirpasic/gods/utils"
	"testing"
)

// For testing purposes
type ContainerTest struct {
	values []interface{}
}

func (container ContainerTest) Empty() bool {
	return len(container.values) == 0
}

func (container ContainerTest) Size() int {
	return len(container.values)
}

func (container ContainerTest) Clear() {
	container.values = []interface{}{}
}

func (container ContainerTest) Values() []interface{} {
	return container.values
}

func TestGetSortedValuesInts(t *testing.T) {
	container := ContainerTest{}
	container.values = []interface{}{5, 1, 3, 2, 4}
	values := GetSortedValues(container, utils.IntComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1].(int) > values[i].(int) {
			t.Errorf("Not sorted!")
		}
	}
}

func TestGetSortedValuesStrings(t *testing.T) {
	container := ContainerTest{}
	container.values = []interface{}{"g", "a", "d", "e", "f", "c", "b"}
	values := GetSortedValues(container, utils.StringComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1].(string) > values[i].(string) {
			t.Errorf("Not sorted!")
		}
	}
}
