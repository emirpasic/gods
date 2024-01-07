package testutils

import "testing"

func SameElements[T comparable](t *testing.T, actual, expected []T) {
	if len(actual) != len(expected) {
		t.Errorf("Got %d expected %d", len(actual), len(expected))
	}
outer:
	for _, e := range expected {
		for _, a := range actual {
			if e == a {
				continue outer
			}
		}
		t.Errorf("Did not find expected element %v in %v", e, actual)
	}
}
