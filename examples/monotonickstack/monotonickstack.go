package main

import (
	ms "github.com/emirpasic/gods/v2/stacks/monotonicstack"
)

func main() {
	stack := ms.New[int](ms.Inc) // empty
	stack.Push(3)                // 3
	stack.Push(1)                // 1
	stack.Push(4)                // 1, 4
	stack.Values()               // 4, 1 (LIFO order)
	stack.Push(1)                // 1, 1
	stack.Push(5)                // 1, 1, 5
	_, _ = stack.Peek()          // 5, true
	_, _ = stack.Pop()           // 5, true
	_, _ = stack.Peek()          // 1, true
	stack.Push(9)                // 1, 1, 5, 9
	stack.Push(2)                // 1, 1, 2
	stack.Push(6)                // 1, 1, 2, 6
	_ = stack.Size()             // 4
}
