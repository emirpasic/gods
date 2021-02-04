package graphs

// Graph holds vertices and edges of a graph
type Graph struct {
	isDirected bool
	vertices   []node
}

type node struct {
	name  interface{}
	edges []interface{}
}

const infinity = int(^uint(0) >> 1)

// NewDirectedGraph is for creating a directed graph
func NewDirectedGraph() *Graph {
	g := Graph{}
	g.vertices = make([]node, 0)
	g.isDirected = true
	return &g
}

// NewUndirectedGraph is for creating an undirected graph
func NewUndirectedGraph() *Graph {
	g := Graph{}
	g.vertices = make([]node, 0)
	g.isDirected = false
	return &g
}

// AddVertice is for adding a new vertice. id vertice exists do nothing
func (g *Graph) AddVertice(name interface{}) {
	g.addVertice(name)
}

// AddEdge if for adding a new edge.
func (g *Graph) AddEdge(from, to interface{}) {
	g.addEdge(from, to)
}

// GetEdge is for getting all the edges of vertice n
func (g *Graph) GetEdge(n interface{}) []interface{} {
	return g.getEdge(n)
}

// GetVertices is for getting all vertices info
func (g *Graph) GetVertices() []interface{} {
	return g.getVertices()
}

// ShortestPath is for finding shortest path from edge s to t by using belman-ford algorithm
func (g *Graph) ShortestPath(s, t interface{}) map[interface{}]int {
	return g.shortestPath(s, t)
}

func (g *Graph) addVertice(name interface{}) int {
	for i, v := range g.vertices {
		if v.name == name {
			return i
		}
	}
	n := node{}
	n.name = name
	n.edges = make([]interface{}, 0)
	g.vertices = append(g.vertices, n)

	return len(g.vertices) - 1
}

func (g *Graph) addEdge(from, to interface{}) {
	f := g.addVertice(from)
	t := g.addVertice(to)
	g.vertices[f].edges = append(g.vertices[f].edges, to)
	if !g.isDirected {
		g.vertices[t].edges = append(g.vertices[t].edges, from)
	}
}

func (g *Graph) getEdge(n interface{}) []interface{} {
	for _, v := range g.vertices {
		if v.name == n {
			return v.edges
		}
	}
	// todo throw exception
	return nil
}

func (g *Graph) getVertices() []interface{} {
	vertices := make([]interface{}, 0)
	for i := 0; i < len(g.vertices); i++ {
		vertices = append(vertices, g.vertices[i].name)
	}
	return vertices
}

func (g *Graph) shortestPath(sName, tName interface{}) map[interface{}]int {
	n := len(g.vertices)
	s := g.addVertice(sName)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = infinity
	}
	dist[s] = 0
	for u := 0; u < n-1; u++ {
		for _, vName := range g.vertices[u].edges {
			v := g.addVertice(vName)
			if dist[u] != infinity && dist[u]+1 < dist[v] {
				dist[v] = dist[u] + 1
			}
		}
	}
	result := make(map[interface{}]int)
	for i, d := range dist {
		result[g.vertices[i].name] = d
	}
	return result
}
