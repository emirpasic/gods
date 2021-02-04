// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trie

import (
	"testing"
)

func TestTrieInsertAndContains(t *testing.T) {
	tree := New()
	tree.Insert("zoo")
	tree.Insert("zebra")
	tree.Insert("zip")
	tree.Insert("voo")
	tree.Insert("foo")
	tree.Insert("voodoo")
	tree.Insert("foo") //overwrite

	if actualValue := tree.Contains("zoo"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := tree.Contains("zebra"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := tree.Contains("zip"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := tree.Contains("voo"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := tree.Contains("foo"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := tree.Contains("voodoo"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}
