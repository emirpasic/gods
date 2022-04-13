// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import llq "github.com/emirpasic/gods/queues/linkedlistqueue"

// LinkedListQueueExample to demonstrate basic usage of LinkedListQueue
func main() {
	queue := llq.New()     // empty
	queue.Enqueue(1)       // 1
	queue.Enqueue(2)       // 1, 2
	_ = queue.Values()     // 1, 2 (FIFO order)
	_, _ = queue.Peek()    // 1,true
	_, _ = queue.Dequeue() // 1, true
	_, _ = queue.Dequeue() // 2, true
	_, _ = queue.Dequeue() // nil, false (nothing to deque)
	queue.Enqueue(1)       // 1
	queue.Clear()          // empty
	queue.Empty()          // true
	_ = queue.Size()       // 0
}
