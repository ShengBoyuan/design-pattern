package abstract_factory

import "fmt"

// Abstract Factory
type GUIFactory interface {
	createButton() Button
	createCheckbox() Checkbox
}

// Concrete Factory
type WinFactory struct{}
type MacFactory struct{}

// Abstract Products
type Button interface {
	Click()
}
type Checkbox interface {
	Select()
}

// Concrete Products
type WinButton struct{}
type MacButton struct{}
type WinCheckbox struct{}
type MacCheckbox struct{}

// Client
type Application struct {
	oneButton   Button
	oneCheckbox Checkbox
	oneFactory  GUIFactory
}

// Implementations
func (WinFactory) createButton() Button {
	return WinButton{}
}

func (WinFactory) createCheckbox() Checkbox {
	return WinCheckbox{}
}

func (MacFactory) createButton() Button {
	return MacButton{}
}

func (MacFactory) createCheckbox() Checkbox {
	return MacCheckbox{}
}

func (WinButton) Click() {
	fmt.Println("I'm a windows button")
}

func (WinCheckbox) Select() {
	fmt.Println("I'm a windows checkbox")
}

func (MacButton) Click() {
	fmt.Println("I'm a macOS button")
}

func (MacCheckbox) Select() {
	fmt.Println("I'm a macOS checkbox")
}

func (a *Application) Init(platform string) {
	if platform == "macOS" {
		a.oneFactory = MacFactory{}
	} else if platform == "windows" {
		a.oneFactory = WinFactory{}
	} else {
		return
	}

	a.oneButton = a.oneFactory.createButton()
	a.oneCheckbox = a.oneFactory.createCheckbox()
}

// Main Processing Flow
func MyOwnApplication() {
	myApp := &Application{}

	myApp.Init("windows")
	myApp.oneButton.Click()
	myApp.oneCheckbox.Select()

	myApp.Init("macOS")
	myApp.oneButton.Click()
	myApp.oneCheckbox.Select()
}
