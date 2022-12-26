package prototype

import "fmt"

// Base Prototype
type Shape interface {
	Clone() Shape
}

// Concrete Prototype
type Circle struct {
	radius float64
}
type Rectangle struct {
	width  int
	height int
}

// Client
func makeSomeShapes() []Shape {
	shapes := make([]Shape, 0, 5)
	shapes = append(shapes, Circle{radius: 1})
	shapes = append(shapes, Circle{radius: 2})
	shapes = append(shapes, Rectangle{width: 3, height: 3})
	shapes = append(shapes, Rectangle{width: 4, height: 4})
	shapes = append(shapes, Rectangle{width: 5, height: 5})
	return shapes
}

// Implementations
func (c Circle) Clone() Shape {
	return Circle{radius: c.radius}
}

func (r Rectangle) Clone() Shape {
	return Rectangle{width: r.width, height: r.height}
}

// Main Processing Flow
func CloneSomeShapes() {
	src := makeSomeShapes()
	res := make([]Shape, 0, len(src))

	for _, t := range src {
		res = append(res, t.Clone())
	}
	fmt.Println(res)
}
