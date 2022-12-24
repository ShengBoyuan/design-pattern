package factory

import "fmt"

// Creator
type Drinker interface {
	GetAlcohol() Alcohol
	Drink()
}

// Concrete Creators
type BeerDrinker struct{}
type WineDrinker struct{}

// Product
type Alcohol interface {
	PreProcess()
}

// Concrete Products
type Beer struct{}
type Wine struct{}

// Main Class
type Bar struct {
	Customer Drinker
}

// Implementations
func (BeerDrinker) GetAlcohol() Alcohol {
	fmt.Println("I wanna a bottle of IPA")
	return &Beer{}
}

func (BeerDrinker) Drink() {
	fmt.Println("e.., I think I prefer Trappist ones.")
}

func (WineDrinker) GetAlcohol() Alcohol {
	fmt.Println("Please give me a cup of Whisky")
	return &Wine{}
}

func (WineDrinker) Drink() {
	fmt.Println("Another drink, thanks~")
}

func (Beer) PreProcess() {
	fmt.Println("added some butter")
}

func (Wine) PreProcess() {
	fmt.Println("mix with twice the amount of Coke")
}

func (b *Bar) ApplyService(drinkerType string) {
	if drinkerType == "Beer" {
		b.Customer = BeerDrinker{}
	} else if drinkerType == "Wine" {
		b.Customer = WineDrinker{}
	} else {
		b.Customer = nil
	}
}

// Main Processing Flow
func OpenMyOwnBar() {
	myBar := Bar{}

	myBar.ApplyService("Beer")
	myBar.Customer.GetAlcohol()
	myBar.Customer.Drink()

	myBar.ApplyService("Wine")
	myBar.Customer.GetAlcohol()
	myBar.Customer.Drink()
}
