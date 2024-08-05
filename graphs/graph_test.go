package graphs

import (
	"testing"
)

func TestGrsphAddVertice(t *testing.T) {
	g1 := NewDirectedGraph()
	g1.AddVertice("a")
	g1.AddVertice("b")
	g1.AddVertice("c")
	g1.AddVertice("d")
	actualValue1 := g1.GetVertices()
	expectedValue1 := []string{"a", "b", "c", "d"}
	for i, v := range expectedValue1 {
		if v != actualValue1[i] {
			t.Errorf("Got %v expected %v", actualValue1, expectedValue1)
			return
		}
	}

	g2 := NewDirectedGraph()
	g2.AddVertice(1)
	g2.AddVertice(2)
	g2.AddVertice(3)
	g2.AddVertice(4)
	actualValue2 := g2.GetVertices()

	expectedValue2 := []int{1, 2, 3, 4}
	for i, v := range expectedValue2 {
		if v != actualValue2[i] {
			t.Errorf("Got %v expected %v", actualValue2, expectedValue2)
			return
		}
	}

	g3 := NewUndirectedGraph()
	g3.AddVertice("a")
	g3.AddVertice("b")
	g3.AddVertice("c")
	g3.AddVertice("d")
	actualValue3 := g3.GetVertices()
	expectedValue3 := []string{"a", "b", "c", "d"}
	for i, v := range expectedValue3 {
		if v != actualValue3[i] {
			t.Errorf("Got %v expected %v", actualValue3, expectedValue3)
			return
		}
	}
}

func TestGrsphAddEdge(t *testing.T) {
	g1 := NewDirectedGraph()
	g1.AddVertice("a")
	g1.AddVertice("b")
	g1.AddEdge("a", "b")
	g1.AddEdge("c", "a")
	actualValue1 := g1.GetEdge("a")
	expectedValue1 := []string{"b"}
	for i, v := range expectedValue1 {
		if v != actualValue1[i] {
			t.Errorf("Got %v expected %v", actualValue1, expectedValue1)
			return
		}
	}

	g2 := NewUndirectedGraph()
	g2.AddVertice("a")
	g2.AddVertice("b")
	g2.AddEdge("a", "b")
	g2.AddEdge("c", "a")
	actualValue2 := g2.GetEdge("a")
	expectedValue2 := []string{"b", "c"}
	for i, v := range expectedValue2 {
		if v != actualValue2[i] {
			t.Errorf("Got %v expected %v", actualValue2, expectedValue2)
			return
		}
	}
}

func TestShortestPath(t *testing.T) {
	g := NewDirectedGraph()
	g.addEdge("a", "b")
	g.addEdge("b", "c")
	g.addEdge("c", "a")
	g.addEdge("c", "d")

	tests := [][]interface{}{
		{1, "a", 0},
		{2, "b", 1},
		{3, "c", 2},
		{4, "d", 3},
	}

	for _, test := range tests {
		actualValue, expectedValue := g.shortestPath("a", "b")[test[1]], test[2]
		if actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
}
