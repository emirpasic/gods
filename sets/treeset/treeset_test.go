/*
Copyright (c) Emir Pasic, All rights reserved.

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 3.0 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library. See the file LICENSE included
with this distribution for more information.
*/

package treeset

import (
	"fmt"
	"testing"
)

func TestHashSet(t *testing.T) {

	set := NewWithIntComparator()

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

	// repeat 10 time since map in golang has a random iteration order each time and we want to make sure that the set is ordered
	for i := 1; i <= 10; i++ {
		if actualValue, expactedValue := fmt.Sprintf("%d%d%d", set.Values()...), "123"; actualValue != expactedValue {
			t.Errorf("Got %v expected %v", actualValue, expactedValue)
		}
	}

	set.Remove()
	set.Remove(1)

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
