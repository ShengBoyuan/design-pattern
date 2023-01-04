package visitor

import "fmt"

// Element Interface
type Shape interface {
	move(int, int)
	draw()
	accept(Visitor)
}

// Concrete Element
type Dot struct{}
type Circle struct{}
type Rectangle struct{}

var _ Shape = Dot{}
var _ Shape = Circle{}
var _ Shape = Rectangle{}

func (Dot) move(x int, y int) {
	fmt.Println("Dot at", x, y)
}

func (Dot) draw() {
	fmt.Println("draw a dot")
}

func (d Dot) accept(v Visitor) {
	v.visitDot(d)
}

func (Circle) move(x int, y int) {
	fmt.Println("Circle at", x, y)
}

func (Circle) draw() {
	fmt.Println("draw a circle")
}

func (c Circle) accept(v Visitor) {
	v.visitCircle(c)
}

func (Rectangle) move(x int, y int) {
	fmt.Println("Rectangle at", x, y)
}

func (Rectangle) draw() {
	fmt.Println("draw an rectangle")
}

func (r Rectangle) accept(v Visitor) {
	v.visitRectangle(r)
}

// Visitor
type Visitor interface {
	visitDot(Dot)
	visitCircle(Circle)
	visitRectangle(Rectangle)
}

// Concrete Visitors
type XMLVisitor struct{}
type HTMLVisitor struct{}

var _ Visitor = XMLVisitor{}
var _ Visitor = HTMLVisitor{}

func (XMLVisitor) visitDot(d Dot) {
	fmt.Println("XML Dot")
	d.move(0, 0)
	d.draw()
}

func (XMLVisitor) visitCircle(c Circle) {
	fmt.Println("XML Circle")
	c.move(1, 1)
	c.draw()
}

func (XMLVisitor) visitRectangle(r Rectangle) {
	fmt.Println("XML Rectangle")
	r.move(2, 6)
	r.draw()
}

func (HTMLVisitor) visitDot(d Dot) {
	fmt.Println("HTML Dot")
	d.move(-4, -4)
	d.draw()
}

func (HTMLVisitor) visitCircle(c Circle) {
	fmt.Println("HTML Circle")
	c.move(-1, -1)
	c.draw()
}

func (HTMLVisitor) visitRectangle(r Rectangle) {
	fmt.Println("HTML Rectangle")
	r.move(-2, -6)
	r.draw()
}

// Main Processing Flow
func LetShapesDrawThemselves() {
	xml := XMLVisitor{}
	html := HTMLVisitor{}
	shapes := make([]Shape, 0)
	shapes = append(shapes, Dot{})
	shapes = append(shapes, Circle{})
	shapes = append(shapes, Dot{})
	shapes = append(shapes, Rectangle{})
	shapes = append(shapes, Circle{})

	for _, s := range shapes {
		s.accept(xml)
		s.accept(html)
	}
}
