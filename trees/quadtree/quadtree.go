
package quadtree

type Quadtree struct {
	Bounds     Bounds
	MaxObjects int
	MaxLevels  int
	Level      int
	Objects    []Bounds
	Nodes      []Quadtree
	Total      int
}

type Bounds struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (b *Bounds) IsPoint() bool {

	if b.Width == 0 && b.Height == 0 {
		return true
	}

	return false

}

func (b *Bounds) Intersects(a Bounds) bool {

	aMaxX := a.X + a.Width
	aMaxY := a.Y + a.Height
	bMaxX := b.X + b.Width
	bMaxY := b.Y + b.Height

	if aMaxX < b.X {
		return false
	}

	if a.X > bMaxX {
		return false
	}

	if aMaxY < b.Y {
		return false
	}

	if a.Y > bMaxY {
		return false
	}

	return true

}

func (qt *Quadtree) TotalNodes() int {

	total := 0

	if len(qt.Nodes) > 0 {
		for i := 0; i < len(qt.Nodes); i++ {
			total += 1
			total += qt.Nodes[i].TotalNodes()
		}
	}

	return total

}

func (qt *Quadtree) split() {

	if len(qt.Nodes) == 4 {
		return
	}

	nextLevel := qt.Level + 1
	subWidth := qt.Bounds.Width / 2
	subHeight := qt.Bounds.Height / 2
	x := qt.Bounds.X
	y := qt.Bounds.Y

	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x + subWidth,
			Y:      y,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x,
			Y:      y,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x,
			Y:      y + subHeight,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x + subWidth,
			Y:      y + subHeight,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

}

func (qt *Quadtree) getIndex(pRect Bounds) int {

	index := -1

	verticalMidpoint := qt.Bounds.X + (qt.Bounds.Width / 2)
	horizontalMidpoint := qt.Bounds.Y + (qt.Bounds.Height / 2)
	topQuadrant := (pRect.Y < horizontalMidpoint) && (pRect.Y+pRect.Height < horizontalMidpoint)
	bottomQuadrant := (pRect.Y > horizontalMidpoint)
	if (pRect.X < verticalMidpoint) && (pRect.X+pRect.Width < verticalMidpoint) {

		if topQuadrant {
			index = 1
		} else if bottomQuadrant {
			index = 2
		}

	} else if pRect.X > verticalMidpoint {

		if topQuadrant {
			index = 0
		} else if bottomQuadrant {
			index = 3
		}

	}

	return index

}

func (qt *Quadtree) Insert(pRect Bounds) {

	qt.Total++

	i := 0
	var index int

	if len(qt.Nodes) > 0 == true {

		index = qt.getIndex(pRect)

		if index != -1 {
			qt.Nodes[index].Insert(pRect)
			return
		}
	}

	qt.Objects = append(qt.Objects, pRect)

	if (len(qt.Objects) > qt.MaxObjects) && (qt.Level < qt.MaxLevels) {

		if len(qt.Nodes) > 0 == false {
			qt.split()
		}

		for i < len(qt.Objects) {

			index = qt.getIndex(qt.Objects[i])

			if index != -1 {

				splice := qt.Objects[i]
				qt.Objects = append(qt.Objects[:i], qt.Objects[i+1:]...)

				qt.Nodes[index].Insert(splice)

			} else {

				i++

			}

		}

	}

}

func (qt *Quadtree) Retrieve(pRect Bounds) []Bounds {

	index := qt.getIndex(pRect)

	returnObjects := qt.Objects

	if len(qt.Nodes) > 0 {
		if index != -1 {

			returnObjects = append(returnObjects, qt.Nodes[index].Retrieve(pRect)...)

		} else {

			for i := 0; i < len(qt.Nodes); i++ {
				returnObjects = append(returnObjects, qt.Nodes[i].Retrieve(pRect)...)
			}

		}
	}

	return returnObjects

}

func (qt *Quadtree) RetrievePoints(find Bounds) []Bounds {

	var foundPoints []Bounds
	potentials := qt.Retrieve(find)
	for o := 0; o < len(potentials); o++ {

		xyMatch := potentials[o].X == float64(find.X) && potentials[o].Y == float64(find.Y)
		if xyMatch && potentials[o].IsPoint() {
			foundPoints = append(foundPoints, find)
		}
	}

	return foundPoints

}

func (qt *Quadtree) RetrieveIntersections(find Bounds) []Bounds {

	var foundIntersections []Bounds

	potentials := qt.Retrieve(find)
	for o := 0; o < len(potentials); o++ {
		if potentials[o].Intersects(find) {
			foundIntersections = append(foundIntersections, potentials[o])
		}
	}

	return foundIntersections

}

func (qt *Quadtree) Clear() {

	qt.Objects = []Bounds{}

	if len(qt.Nodes)-1 > 0 {
		for i := 0; i < len(qt.Nodes); i++ {
			qt.Nodes[i].Clear()
		}
	}

	qt.Nodes = []Quadtree{}
	qt.Total = 0

}
