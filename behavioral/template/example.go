package template

import "fmt"

// Abstract Class
type GameAI interface {
	buildStructures()
	buildUnits()
	doSomething()
}

// Concrete Classes
type Monster struct{}
type Soilder struct{}

var _ GameAI = Monster{}
var _ GameAI = Soilder{}

func (Monster) buildStructures() {
	fmt.Println("I've built my own lair")
}

func (Monster) buildUnits() {
	fmt.Println("Come here my fellows!")
}

func (Monster) doSomething() {
	fmt.Println("Destroy all their castles!")
}

func (Soilder) buildStructures() {
	fmt.Println("Constructions complete")
}

func (Soilder) buildUnits() {
	fmt.Println("Brothers and sisters, join me!")
}

func (Soilder) doSomething() {
	fmt.Println("Defend our homeland!")
}

// Template
type AITemplate struct {
	HP   int
	MP   int
	Race GameAI
}

func (a *AITemplate) Init() { // Step1() ~ StepN()
	a.HP = 100
	a.MP = 100
}

// Main Processing Flow
func PlayGame() {
	soilder := &AITemplate{}
	soilder.Init()
	soilder.Race = Soilder{}

	monster := &AITemplate{}
	monster.Init()
	monster.Race = Monster{}

	soilder.Race.buildStructures()
	soilder.Race.buildUnits()
	soilder.Race.doSomething()

	monster.Race.buildStructures()
	monster.Race.buildUnits()
	monster.Race.doSomething()
}
