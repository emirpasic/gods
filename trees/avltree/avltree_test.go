// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package avltree

import (
	"encoding/json"
	"slices"
	"strings"
	"testing"
)

func TestAVLTreeGet(t *testing.T) {
	tree := New[int, string]()

	if actualValue := tree.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	if actualValue := tree.GetNode(2).Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)
	//
	//  AVLTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      │   ┌── 3
	//      └── 2
	//          └── 1

	if actualValue := tree.Size(); actualValue != 6 {
		t.Errorf("Got %v expected %v", actualValue, 6)
	}

	if actualValue := tree.GetNode(2).Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue := tree.GetNode(4).Size(); actualValue != 6 {
		t.Errorf("Got %v expected %v", actualValue, 6)
	}

	if actualValue := tree.GetNode(7).Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestAVLTreePut(t *testing.T) {
	tree := New[int, string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite

	if actualValue := tree.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}
	if actualValue, expectedValue := tree.Keys(), []int{1, 2, 3, 4, 5, 6, 7}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := tree.Values(), []string{"a", "b", "c", "d", "e", "f", "g"}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	tests1 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, "", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := tree.Get(test[0].(int))
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}
}

func TestAVLTreeRemove(t *testing.T) {
	tree := New[int, string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite

	tree.Remove(5)
	tree.Remove(6)
	tree.Remove(7)
	tree.Remove(8)
	tree.Remove(5)

	if actualValue, expectedValue := tree.Keys(), []int{1, 2, 3, 4}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := tree.Values(), []string{"a", "b", "c", "d"}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := tree.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}

	tests2 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "", false},
		{6, "", false},
		{7, "", false},
		{8, "", false},
	}

	for _, test := range tests2 {
		actualValue, actualFound := tree.Get(test[0].(int))
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	tree.Remove(1)
	tree.Remove(4)
	tree.Remove(2)
	tree.Remove(3)
	tree.Remove(2)
	tree.Remove(2)

	if actualValue, expectedValue := tree.Keys(), []int{}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := tree.Values(), []string{}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if empty, size := tree.Empty(), tree.Size(); empty != true || size != -0 {
		t.Errorf("Got %v expected %v", empty, true)
	}

}

func TestAVLTreeLeftAndRight(t *testing.T) {
	tree := New[int, string]()

	if actualValue := tree.Left(); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := tree.Right(); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}

	tree.Put(1, "a")
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x") // overwrite
	tree.Put(2, "b")

	if actualValue, expectedValue := tree.Left().Key, 1; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := tree.Left().Value, "x"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, expectedValue := tree.Right().Key, 7; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := tree.Right().Value, "g"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeCeilingAndFloor(t *testing.T) {
	tree := New[int, string]()

	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
	if node, found := tree.Ceiling(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")

	if node, found := tree.Floor(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	if node, found := tree.Ceiling(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Ceiling(8); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
}

func TestAVLTreeIteratorNextOnEmpty(t *testing.T) {
	tree := New[int, string]()
	it := tree.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestAVLTreeIteratorPrevOnEmpty(t *testing.T) {
	tree := New[int, string]()
	it := tree.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestAVLTreeIterator1Next(t *testing.T) {
	tree := New[int, string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()

	count := 0
	for it.Next() {
		count++
		key := it.Key()
		if actualValue, expectedValue := key, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator1Prev(t *testing.T) {
	tree := New[int, string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator2Next(t *testing.T) {
	tree := New[int, string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		if actualValue, expectedValue := key, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator2Prev(t *testing.T) {
	tree := New[int, string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator3Next(t *testing.T) {
	tree := New[int, string]()
	tree.Put(1, "a")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		if actualValue, expectedValue := key, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator3Prev(t *testing.T) {
	tree := New[int, string]()
	tree.Put(1, "a")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator4Next(t *testing.T) {
	tree := New[int, int]()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		value := it.Value()
		if actualValue, expectedValue := value, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator4Prev(t *testing.T) {
	tree := New[int, int]()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := tree.Size()
	for it.Next() {
	}
	for it.Prev() {
		value := it.Value()
		if actualValue, expectedValue := value, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		count--
	}
	if actualValue, expectedValue := count, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIteratorBegin(t *testing.T) {
	tree := New[int, string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()

	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	it.Begin()

	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	for it.Next() {
	}

	it.Begin()

	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	it.Next()
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestAVLTreeIteratorEnd(t *testing.T) {
	tree := New[int, string]()
	it := tree.Iterator()

	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	it.End()
	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it.End()
	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	it.Prev()
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestAVLTreeIteratorFirst(t *testing.T) {
	tree := New[int, string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestAVLTreeIteratorLast(t *testing.T) {
	tree := New[int, string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestAVLTreeIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// NextTo (empty)
	{
		tree := New[int, string]()
		it := tree.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
	}

	// NextTo (not found)
	{
		tree := New[int, string]()
		tree.Put(0, "xx")
		tree.Put(1, "yy")
		it := tree.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
	}

	// NextTo (found)
	{
		tree := New[int, string]()
		tree.Put(2, "cc")
		tree.Put(0, "aa")
		tree.Put(1, "bb")
		it := tree.Iterator()
		it.Begin()
		if !it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
		if index, value := it.Key(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Next() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Key(), it.Value(); index != 2 || value != "cc" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "cc")
		}
		if it.Next() {
			t.Errorf("Should not go past last element")
		}
	}
}

func TestAVLTreeIteratorPrevTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// PrevTo (empty)
	{
		tree := New[int, string]()
		it := tree.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
	}

	// PrevTo (not found)
	{
		tree := New[int, string]()
		tree.Put(0, "xx")
		tree.Put(1, "yy")
		it := tree.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
	}

	// PrevTo (found)
	{
		tree := New[int, string]()
		tree.Put(2, "cc")
		tree.Put(0, "aa")
		tree.Put(1, "bb")
		it := tree.Iterator()
		it.End()
		if !it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
		if index, value := it.Key(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Prev() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Key(), it.Value(); index != 0 || value != "aa" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "aa")
		}
		if it.Prev() {
			t.Errorf("Should not go before first element")
		}
	}
}

func TestAVLTreeSerialization(t *testing.T) {
	tree := New[string, string]()
	tree.Put("c", "3")
	tree.Put("b", "2")
	tree.Put("a", "1")

	var err error
	assert := func() {
		if actualValue, expectedValue := tree.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue, expectedValue := tree.Keys(), []string{"a", "b", "c"}; !slices.Equal(actualValue, expectedValue) {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue, expectedValue := tree.Values(), []string{"1", "2", "3"}; !slices.Equal(actualValue, expectedValue) {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	bytes, err := tree.ToJSON()
	assert()

	err = tree.FromJSON(bytes)
	assert()

	bytes, err = json.Marshal([]interface{}{"a", "b", "c", tree})
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	intTree := New[string, int]()
	err = json.Unmarshal([]byte(`{"a":1,"b":2}`), intTree)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	if actualValue, expectedValue := intTree.Size(), 2; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := intTree.Keys(), []string{"a", "b"}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := intTree.Values(), []int{1, 2}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeString(t *testing.T) {
	c := New[int, int]()
	c.Put(1, 1)
	c.Put(2, 1)
	c.Put(3, 1)
	c.Put(4, 1)
	c.Put(5, 1)
	c.Put(6, 1)
	c.Put(7, 1)
	c.Put(8, 1)

	if !strings.HasPrefix(c.String(), "AVLTree") {
		t.Errorf("String should start with container name")
	}
}

func benchmarkGet(b *testing.B, tree *Tree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, tree *Tree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, tree *Tree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Remove(n)
		}
	}
}

func BenchmarkAVLTreeGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkAVLTreeGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkAVLTreeGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkAVLTreeGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkAVLTreePut100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := New[int, struct{}]()
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkAVLTreePut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkAVLTreePut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkAVLTreePut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkAVLTreeRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkAVLTreeRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkAVLTreeRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkAVLTreeRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}
