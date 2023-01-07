package mediator

import "reflect"

// Mediator
type Mediator interface {
	notify(Component, string)
}

// Concrete Mediator
type AuthenticationDialog struct {
	Title         string
	LoginChkBx    Checkbox
	LoginUsername Textbox
	OkBtn         Button
	CancelBtn     Button
}

func (a *AuthenticationDialog) Init(login Checkbox, loginName Textbox, okB Button, cancelB Button) {
	a.LoginChkBx = login
	a.LoginUsername = loginName
	a.OkBtn = okB
	a.CancelBtn = cancelB
}

func (a *AuthenticationDialog) notify(sender Component, event string) {
	if _, ok := sender.(*Button); ok {
		// judge and do
		a.Title = "it's a Button Click"
	}
	if reflect.TypeOf(sender) == reflect.TypeOf(&Button{}) && event == "click" {
		// Another ugly judgement
		a.Title = "it's a Button Click"
	}
}

// Component
type Component interface {
	click()
	keypress()
}

// Concrete Components
type Button struct {
	dialog Mediator
}

type Textbox struct {
	dialog Mediator
}

type Checkbox struct {
	dialog Mediator
}

func (b *Button) Init(m Mediator) {
	b.dialog = m
}

func (b *Button) click() {
	b.dialog.notify(b, "click")
}

func (b *Button) keypress() {
	b.dialog.notify(b, "keypress")
}

func (t *Textbox) Init(m Mediator) {
	t.dialog = m
}

func (t *Textbox) click() {
	t.dialog.notify(t, "click")
}

func (t *Textbox) keypress() {
	t.dialog.notify(t, "keypress")
}

func (c *Checkbox) Init(m Mediator) {
	c.dialog = m
}

func (c *Checkbox) click() {
	c.dialog.notify(c, "click")
}

func (c *Checkbox) keypress() {
	c.dialog.notify(c, "keypress")
}
