package builder

import (
	"fmt"
)

// Abstract Builder
type Builder interface {
	cleanCup()
	getBaseWine(string)
	getIngredients([]string)
	mix()
	shakeShake(bool)
}

// Products & Concrete Builder
type Cocktail struct {
	name string
}
type Beer struct {
	name string
}

type CocktailBuilder struct {
	base            string
	ingredients     []string
	finishedProduct Cocktail
}
type BeerBuilder struct {
	base            string
	finishedProduct Beer
}

// Director
type Bartender struct {
	MixingGlass Builder
}

func (b *Bartender) OrderWine(builder Builder) {
	switch builder.(type) {
	case *CocktailBuilder:
		b.MixingGlass = builder
		b.MakeCocktail()
	case *BeerBuilder:
		b.MixingGlass = builder
		b.MakeBeer()
	default:
		b.MixingGlass = nil
		fmt.Println("Sorry, I don't know how to make it")
	}
}

func (b *Bartender) MakeCocktail() {
	b.MixingGlass.cleanCup()
	b.MixingGlass.getBaseWine("Vodka")
	b.MixingGlass.getIngredients([]string{"tomato juice", "chili sauce"})
	b.MixingGlass.mix()
	b.MixingGlass.shakeShake(true)
}

func (b *Bartender) MakeBeer() {
	b.MixingGlass.cleanCup()
	b.MixingGlass.getBaseWine("Stout")
	b.MixingGlass.getIngredients(nil)
	b.MixingGlass.mix()
	b.MixingGlass.shakeShake(false)
}

// Client
type Bar struct {
	name      string
	bartender Bartender
}

func (b *Bar) SayHello() {
	fmt.Println("Hey! Welcome to", b.name)
}

func (b *Bar) StartBusiness() {
	b.bartender = Bartender{}
}

// Implementations
func (c *CocktailBuilder) cleanCup() {
	c.finishedProduct = Cocktail{}
	c.base = ""
	c.ingredients = make([]string, 0)
	fmt.Println("cocktail cup ready")
}

func (c *CocktailBuilder) getBaseWine(base string) {
	c.base = base
	fmt.Println(base, "added")
}

func (c *CocktailBuilder) getIngredients(ingredients []string) {
	c.ingredients = ingredients
	for _, t := range c.ingredients {
		fmt.Println(t, "added")
	}
}

func (c *CocktailBuilder) mix() {
	c.finishedProduct.name = "Bloody Mary"
	fmt.Println(c.finishedProduct.name, "making ~")
}

func (c *CocktailBuilder) shakeShake(t bool) {
	if t {
		fmt.Println("shake ~ shake ~")
	} else {
		fmt.Println("I suppose this wine should not be shaken")
	}
}

func (c *CocktailBuilder) serveWine() Cocktail {
	fmt.Println("Here you are, your", c.finishedProduct.name, ", Enjoy ~")
	return c.finishedProduct
}

func (b *BeerBuilder) cleanCup() {
	b.finishedProduct = Beer{}
	b.base = ""
	fmt.Println("beer cup ready")
}

func (b *BeerBuilder) getBaseWine(base string) {
	b.base = base
	fmt.Println(base, "added")
}

func (b *BeerBuilder) getIngredients([]string) {
	fmt.Println("beer don't need any decoration")
}

func (b *BeerBuilder) mix() {
	b.finishedProduct.name = b.base
	fmt.Println("nothing to mix in")
}

func (b *BeerBuilder) shakeShake(t bool) {
	if t {
		fmt.Println("shake ~ shake ~")
	} else {
		fmt.Println("I suppose this wine should not be shaken")
	}
}

func (b *BeerBuilder) serveWine() Beer {
	fmt.Println("Here you are, your", b.finishedProduct.name, ", Enjoy ~")
	return b.finishedProduct
}

// Test two concrete builders whether implement Abstract Builder correctly
var TestCocktail Builder = &CocktailBuilder{}
var TestBeer Builder = &BeerBuilder{}

// Main processing flow
func OpenMyOwnBar() {
	// init
	myBar := Bar{name: "SBY's little bar"}
	myBar.SayHello()
	myBar.StartBusiness()

	// serve differentcustomer
	cocktailBuilder := &CocktailBuilder{}
	myBar.bartender.OrderWine(cocktailBuilder)
	cocktail := cocktailBuilder.serveWine()
	fmt.Println(cocktail.name, "is so spicy")

	beerBuilder := &BeerBuilder{}
	myBar.bartender.OrderWine(beerBuilder)
	beer := beerBuilder.serveWine()
	fmt.Println(beer.name, "tastes too bitter")

	myBar.bartender.OrderWine(nil)
}
