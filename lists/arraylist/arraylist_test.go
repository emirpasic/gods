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

package arraylist

import (
	"testing"
)

func TestArrayList(t *testing.T) {

	list := New()

	list.Add("a", "b", "c", "d", "e", "f", "g")
	list.Clear()

	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	list.Add("a")
	list.Add("b", "c")

	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}

	list.Remove(2)

	if actualValue, ok := list.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}

	list.Remove(1)
	list.Remove(0)

	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	list.Add("a", "b", "c")

	if actualValue := list.Contains("a", "b", "c"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := list.Contains("a", "b", "c", "d"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	list.Clear()

	if actualValue := list.Contains("a"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue, ok := list.Get(0); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

}

func BenchmarkArrayList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := New()
		for n := 0; n < 1000; n++ {
			list.Add(i)
		}
		for !list.Empty() {
			list.Remove(0)
		}
	}

}
