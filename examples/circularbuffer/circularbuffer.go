// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import cb "github.com/emirpasic/gods/queues/circularbuffer"

// CircularBufferExample to demonstrate basic usage of CircularBuffer
func main() {
	queue := cb.New(3)     // empty (max size is 3)
	queue.Enqueue(1)       // 1
	queue.Enqueue(2)       // 1, 2
	queue.Enqueue(3)       // 1, 2, 3
	_ = queue.Values()     // 1, 2, 3
	queue.Enqueue(3)       // 4, 2, 3
	_, _ = queue.Peek()    // 4,true
	_, _ = queue.Dequeue() // 4, true
	_, _ = queue.Dequeue() // 2, true
	_, _ = queue.Dequeue() // 3, true
	_, _ = queue.Dequeue() // nil, false (nothing to deque)
	queue.Enqueue(1)       // 1
	queue.Clear()          // empty
	queue.Empty()          // true
	_ = queue.Size()       // 0
}
