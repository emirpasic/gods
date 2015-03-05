[![Build Status](https://travis-ci.org/emirpasic/gods.svg)](https://travis-ci.org/emirpasic/gods) 

# GoDS (Go Data Structures)

Implementation of various data structures in Go. 

## Implementations

- Sets:
  - HashSet (unordered)
  - TreeSet (ordered)
- Stacks:
  - LinkedListStack
  - ArrayStack
- Maps:
  - HashMap (unordered)
  - TreeMap (ordered)
- Tree:
  - RedBlackTree

## Motivations

Java Collections, C++ Standard Template Library (STL) containers, Qt Containers missing in Go. 

## Goals

- Fast algorithms: 
  - Based on decades of knowledge and experiences of other libraries mentioned below.
- Memory efficient algorithms: 
  - Avoiding to keep consume memory by using optimal algorithms and data structures for the given set of problems.
- Easy to use library: 
  - Well-structued library with minimalistic set of atomic operations from which more complex operations can be crafted.
- Stable library: 
  - Only additions are permitted keeping the library backward compatible.
- Solid documentation and examples: 
  - TODO
- Production ready: 
  - TODO 

There is often a tug of war between speed and memory when crafting algorithms. We choose to optimize for speed in most cases within reasonable limits on memory consumption.

## Testing and Benchmarking

`go test -v -bench . -benchmem  -benchtime 1s ./...`

## License

Copyright (c) Emir Pasic, All rights reserved.

GNU Lesser General Public License Version 3, see [LICENSE](https://github.com/emirpasic/gods/blob/master/LICENSE) for more details.