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

package hashset

import (
	"fmt"
	"testing"
)

func TestHashSet(t *testing.T) {

	set := New()

	// insertions
	set.Add()
	set.Add(1)
	set.Add(2)
	set.Add(2, 3)
	set.Add()

	if actualValue := set.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	// Asking if a set is superset of nothing, thus it's true
	if actualValue := set.Contains(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := set.Contains(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := set.Contains(1, 2, 3); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := set.Contains(1, 2, 3, 4); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	set.Remove()
	set.Remove(1)

	if actualValue, expectedValues := fmt.Sprintf("%d%d", set.Values()...), [2]string{"23", "32"}; actualValue != expectedValues[0] && actualValue != expectedValues[1] {
		t.Errorf("Got %v expected %v", actualValue, expectedValues)
	}

	if actualValue := set.Contains(1); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	set.Remove(1, 2, 3)

	if actualValue := set.Contains(3); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := set.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

}

func BenchmarkHashSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := New()
		for n := 0; n < 1000; n++ {
			set.Add(i)
		}
		for n := 0; n < 1000; n++ {
			set.Remove(n)
		}
	}
}
