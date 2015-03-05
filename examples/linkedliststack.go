package examples

import lls "github.com/emirpasic/gods/stacks/linkedliststack"

func linkedliststackExample() {
	stack := lls.New()  // empty
	stack.Push(1)       // 1
	stack.Push(2)       // 1, 2
	_, _ = stack.Peek() // 2,true
	_, _ = stack.Pop()  // 2, true
	_, _ = stack.Pop()  // 1, true
	_, _ = stack.Pop()  // nil, false (nothing to pop)
	stack.Push(1)       // 1
	stack.Clear()       // empty
	stack.Empty()       // true
	stack.Size()        // 0
}
