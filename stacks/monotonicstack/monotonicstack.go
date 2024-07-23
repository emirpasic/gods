package monotonicstack

import (
	"cmp"

	"github.com/emirpasic/gods/v2/stacks"
	"github.com/emirpasic/gods/v2/stacks/arraystack"
	"github.com/emirpasic/gods/v2/utils"
)

// Assert Stack implementation
var _ stacks.Stack[int] = (*Stack[int])(nil)

type MonoType int

const (
	Inc MonoType = iota
	Dec          = iota
)

// Stack holds elements in an array-list (as embeded type).
type Stack[T comparable] struct {
	stacks.Stack[T]
	Type       MonoType
	Comparator utils.Comparator[T] // Key comparator
}

// New instantiates a monotonic stack with stackType (descreasing or icreasing) based on arraystack.
func New[T cmp.Ordered](stackType MonoType) *Stack[T] {
	return NewWith(stackType, arraystack.New[T]())
}

// NewWith instantiates a monotonic stack with stackType and the custom stack.
func NewWith[T cmp.Ordered](stackType MonoType, stack stacks.Stack[T]) *Stack[T] {
	return &Stack[T]{
		Stack:      stack,
		Type:       stackType,
		Comparator: cmp.Compare[T],
	}
}

// Push adds a value onto the top of the stack (if needed).
func (s *Stack[T]) Push(value T) {
	for s.Stack.Size() > 0 {
		top, _ := s.Stack.Peek()

		poped := false
		switch s.Type {
		case Dec: // remove top if top < new value
			// if a new element is greater than top of stack then we remove
			// elements from the top of the stack
			if s.Comparator(top, value) == -1 { // "-1" == <
				s.Stack.Pop()
				poped = true
			}
		case Inc: // remove top if top > new value
			// pop elements from the stack that are smaller than the new element
			if s.Comparator(top, value) == 1 { // "1" == >
				s.Stack.Pop()
				poped = true
			}
		}

		// should stop if nothing to pop
		if !poped {
			break
		}
	}
	s.Stack.Push(value)
}
