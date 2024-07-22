package monotonicstack

import (
	"slices"
	"testing"

	"github.com/emirpasic/gods/v2/stacks/linkedliststack"
)

func reverse[T any](xs []T) {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
}

func TestMonotonicStack(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack[int]
		in    []int
		want  []int
	}{
		// arraystack
		{"ArrayStack, Inc", New[int](Inc), []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 6}},
		{"ArrayStack, Inc Stable", New[int](Inc), []int{1, 3, 10, 15, 17}, []int{1, 3, 10, 15, 17}},
		{"ArrayStack, Dec", New[int](Dec), []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{9, 6}},
		{"ArrayStack, Dec Stable", New[int](Dec), []int{17, 14, 10, 5, 1}, []int{17, 14, 10, 5, 1}},

		// linkedliststack
		{"ListStack, Inc", NewWith[int](Inc, linkedliststack.New[int]()), []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 6}},
		{"ListStack, Dec", NewWith[int](Dec, linkedliststack.New[int]()), []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{9, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for _, n := range tt.in {
				tt.stack.Push(n)
			}
			got := tt.stack.Values()
			reverse(got) // fix LIFO order

			if !slices.Equal(got, tt.want) {
				t.Errorf("Got %v expected %v", got, tt.want)
			}
		})
	}
}
