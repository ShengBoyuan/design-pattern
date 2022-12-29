package flyweight

import "fmt"

// Flyweight Class
type TreeType struct {
	Name    string
	Color   string
	Texture string
}

func (t *TreeType) Init(name string, color string, texture string) {
	t = &TreeType{
		Name:    name,
		Color:   color,
		Texture: texture,
	}
}

// Flyweight Factory
type TreeFactory struct {
	Trees map[int]*TreeType
}

func (t *TreeFactory) GetTreeType(id int) *TreeType {
	if tree, ok := t.Trees[id]; ok {
		return tree
	} else {
		tree := &TreeType{
			Name:    "demo",
			Color:   "white",
			Texture: "",
		}
		t.Trees[id] = tree
		return tree
	}
}

// Contextual Class / Object
type Tree struct {
	X    int
	Y    int
	Type *TreeType
}

func (t *Tree) Draw() {
	fmt.Println("I'm a", t.Type.Texture, "Tree at", t.X, t.Y)
}

// Clients
func PlantSomeTree() {
	factory := &TreeFactory{
		Trees: make(map[int]*TreeType),
	}
	factory.Trees[0] = &TreeType{
		Name:    "tree0",
		Color:   "black",
		Texture: "wooden",
	}

	tree0 := &Tree{
		X:    0,
		Y:    0,
		Type: factory.GetTreeType(0),
	}
	tree0.Draw()

	treeNil := &Tree{
		X:    1,
		Y:    1,
		Type: factory.GetTreeType(1),
	}
	treeNil.Draw()
}
