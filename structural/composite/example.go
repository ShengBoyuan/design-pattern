package composite

import "fmt"

// Component
type Graphic interface {
	draw()
}

// Leaf
type Dot struct {
	x int
	y int
}

type Circle struct {
	Dot
	radius int
}

func (d *Dot) draw() {
	fmt.Println("I'm a dot at", d.x, d.y)
}

func (c *Circle) draw() {
	fmt.Println("I'm a circle at", c.x, c.y, "my radius is", c.radius)
}

// Composite
type CompoundGraphic struct {
	elements []Graphic
}

func (c *CompoundGraphic) draw() {
	for _, t := range c.elements {
		t.draw()
	}
}

func (c *CompoundGraphic) add(g Graphic) {
	c.elements = append(c.elements, g)
}

func (c *CompoundGraphic) remove() {
	c.elements = c.elements[:len(c.elements)]
}

// Client
func ImageEditor() {
	all := &CompoundGraphic{}
	all.add(&Dot{x: 1, y: 1})
	all.add(&Circle{radius: 1})
	all.draw()

	// the 4 lines below will trigger an Endless loop, cuz 'all' is a pointer,
	// when all is assigned to itself, it will make a loop call.
	// all.add(all)
	// all.add(&Dot{x: 6, y: 1})
	// all.draw()
	// all.remove()

	allForAll := &CompoundGraphic{}
	allForAll.add(all)
	allForAll.draw()
	all.remove()
	all.draw()
}
