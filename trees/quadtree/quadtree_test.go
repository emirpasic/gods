package quadtree

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuadtreeCreation(t *testing.T) {
	qt := setupQuadtree(0, 0, 640, 480)
	if qt.Bounds.Width != 640 && qt.Bounds.Height != 480 {
		t.Errorf("Quadtree was not created correctly")
	}
}

func TestSplit(t *testing.T) {

	qt := setupQuadtree(0, 0, 640, 480)
	qt.split()
	if len(qt.Nodes) != 4 {
		t.Error("Quadtree did not split correctly, expected 4 nodes got", len(qt.Nodes))
	}

	qt.split()
	if len(qt.Nodes) != 4 {
		t.Error("Quadtree should not split itself more than once", len(qt.Nodes))
	}

}

func TestTotalSubnodes(t *testing.T) {

	qt := setupQuadtree(0, 0, 640, 480)
	qt.split()
	for i := 0; i < len(qt.Nodes); i++ {
		qt.Nodes[i].split()
	}

	total := qt.TotalNodes()
	if total != 20 {
		t.Error("Quadtree did not split correctly, expected 20 nodes got", total)
	}

}

func TestQuadtreeInsert(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano())

	qt := setupQuadtree(0, 0, 640, 480)

	grid := 10.0
	gridh := qt.Bounds.Width / grid
	gridv := qt.Bounds.Height / grid
	var randomObject Bounds
	numObjects := 1000

	for i := 0; i < numObjects; i++ {

		x := randMinMax(0, gridh) * grid
		y := randMinMax(0, gridv) * grid

		randomObject = Bounds{
			X:      x,
			Y:      y,
			Width:  randMinMax(1, 4) * grid,
			Height: randMinMax(1, 4) * grid,
		}

		index := qt.getIndex(randomObject)
		if index < -1 || index > 3 {
			t.Errorf("The index should be -1 or between 0 and 3, got %d \n", index)
		}

		qt.Insert(randomObject)

	}

	if qt.Total != numObjects {
		t.Errorf("Error: Should have totalled %d, got %d \n", numObjects, qt.Total)
	} else {
		t.Logf("Success: Total objects in the Quadtree is %d (as expected) \n", qt.Total)
	}

}

func TestCorrectQuad(t *testing.T) {

	qt := setupQuadtree(0, 0, 100, 100)

	var index int
	pass := true

	topRight := Bounds{
		X:      99,
		Y:      99,
		Width:  0,
		Height: 0,
	}
	qt.Insert(topRight)
	index = qt.getIndex(topRight)
	if index == 0 {
		t.Errorf("The index should be 0, got %d \n", index)
		pass = false
	}

	topLeft := Bounds{
		X:      99,
		Y:      1,
		Width:  0,
		Height: 0,
	}
	qt.Insert(topLeft)
	index = qt.getIndex(topLeft)
	if index == 1 {
		t.Errorf("The index should be 1, got %d \n", index)
		pass = false
	}

	bottomLeft := Bounds{
		X:      1,
		Y:      1,
		Width:  0,
		Height: 0,
	}
	qt.Insert(bottomLeft)
	index = qt.getIndex(bottomLeft)
	if index == 2 {
		t.Errorf("The index should be 2, got %d \n", index)
		pass = false
	}

	bottomRight := Bounds{
		X:      1,
		Y:      51,
		Width:  0,
		Height: 0,
	}
	qt.Insert(bottomRight)
	index = qt.getIndex(bottomRight)
	if index == 3 {
		t.Errorf("The index should be 3, got %d \n", index)
		pass = false
	}

	if pass == true {
		t.Log("Success: The points were inserted into the correct quadrants")
	}

}

func TestQuadtreeRetrieval(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano())

	qt := setupQuadtree(0, 0, 640, 480)

	var randomObject Bounds
	numObjects := 100

	for i := 0; i < numObjects; i++ {

		randomObject = Bounds{
			X:      float64(i),
			Y:      float64(i),
			Width:  0,
			Height: 0,
		}

		qt.Insert(randomObject)

	}

	for j := 0; j < numObjects; j++ {

		Cursor := Bounds{
			X:      float64(j),
			Y:      float64(j),
			Width:  0,
			Height: 0,
		}

		objects := qt.Retrieve(Cursor)

		found := false

		if len(objects) >= numObjects {
			t.Error("Objects should not be equal to or bigger than the number of retrieved objects")
		}

		for o := 0; o < len(objects); o++ {
			if objects[o].X == float64(j) && objects[o].Y == float64(j) {
				found = true
			}
		}
		if found != true {
			t.Error("Error finding the correct point")
		}

	}

}

func TestQuadtreeRandomPointRetrieval(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano())

	qt := setupQuadtree(0, 0, 640, 480)

	numObjects := 1000

	for i := 1; i < numObjects+1; i++ {

		randomObject := Bounds{
			X:      float64(i),
			Y:      float64(i),
			Width:  0,
			Height: 0,
		}

		qt.Insert(randomObject)

	}

	failure := false
	iterations := 20
	for j := 1; j < iterations+1; j++ {

		Cursor := Bounds{
			X:      float64(j),
			Y:      float64(j),
			Width:  0,
			Height: 0,
		}

		point := qt.RetrievePoints(Cursor)

		for k := 0; k < len(point); k++ {
			if point[k].X == 0 {
				failure = true
			}
			if point[k].Y == 0 {
				failure = true
			}
			if failure {
				t.Error("Point was incorrectly retrieved", point)
			}
			if point[k].IsPoint() == false {
				t.Error("Point should have width and height of 0")
			}
		}

	}

	if failure == false {
		t.Logf("Success: All the points were retrieved correctly", iterations, numObjects)
	}

}

func TestIntersectionRetrieval(t *testing.T) {
	qt := setupQuadtree(0, 0, 640, 480)
	qt.Insert(Bounds{
		X:      1,
		Y:      1,
		Width:  10,
		Height: 10,
	})
	qt.Insert(Bounds{
		X:      5,
		Y:      5,
		Width:  10,
		Height: 10,
	})
	qt.Insert(Bounds{
		X:      10,
		Y:      10,
		Width:  10,
		Height: 10,
	})
	qt.Insert(Bounds{
		X:      15,
		Y:      15,
		Width:  10,
		Height: 10,
	})
	inter := qt.RetrieveIntersections(Bounds{
		X:      5,
		Y:      5,
		Width:  2.5,
		Height: 2.5,
	})
	if len(inter) != 2 {
		t.Error("Should have two intersections")
	}
}

func TestQuadtreeClear(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano()) // Seed Random properly

	qt := setupQuadtree(0, 0, 640, 480)

	grid := 10.0
	gridh := qt.Bounds.Width / grid
	gridv := qt.Bounds.Height / grid
	var randomObject Bounds
	numObjects := 1000

	for i := 0; i < numObjects; i++ {

		x := randMinMax(0, gridh) * grid
		y := randMinMax(0, gridv) * grid

		randomObject = Bounds{
			X:      x,
			Y:      y,
			Width:  randMinMax(1, 4) * grid,
			Height: randMinMax(1, 4) * grid,
		}

		index := qt.getIndex(randomObject)
		if index < -1 || index > 3 {
			t.Errorf("The index should be -1 or between 0 and 3, got %d \n", index)
		}

		qt.Insert(randomObject)

	}

	qt.Clear()

	if qt.Total != 0 {
		t.Errorf("Error: The Quadtree should be cleared")
	} else {
		t.Logf("Success: The Quadtree was cleared correctly")
	}

}


func BenchmarkInsertOneThousand(b *testing.B) {

	qt := setupQuadtree(0, 0, 640, 480)

	grid := 10.0
	gridh := qt.Bounds.Width / grid
	gridv := qt.Bounds.Height / grid
	var randomObject Bounds
	numObjects := 1000

	for n := 0; n < b.N; n++ {
		for i := 0; i < numObjects; i++ {

			x := randMinMax(0, gridh) * grid
			y := randMinMax(0, gridv) * grid

			randomObject = Bounds{
				X:      x,
				Y:      y,
				Width:  randMinMax(1, 4) * grid,
				Height: randMinMax(1, 4) * grid,
			}

			qt.Insert(randomObject)

		}
	}

}


func setupQuadtree(x float64, y float64, width float64, height float64) *Quadtree {

	return &Quadtree{
		Bounds: Bounds{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		},
		MaxObjects: 4,
		MaxLevels:  8,
		Level:      0,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0),
	}

}

func randMinMax(min float64, max float64) float64 {
	val := min + (rand.Float64() * (max - min))
	return val
}