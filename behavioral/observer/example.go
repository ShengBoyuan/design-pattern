package observer

import "fmt"

// Base Publisher
type EventManager struct {
	Listeners map[string]Listener
}

func (e *EventManager) Subscribe(event string, listener Listener) {
	e.Listeners[event] = listener
}

func (e *EventManager) Unsubscribe(event string, listener Listener) {
	delete(e.Listeners, event)
}

func (e *EventManager) Notify(event string, data string) {
	l, ok := e.Listeners[event]
	if ok {
		l.Update(data)
	}
}

// Concrete Publisher
type Editor struct {
	Events  EventManager
	myDiary File
}

type File struct {
	Name string
}

func (e *Editor) Init(name string) *Editor {
	e.Events.Listeners = make(map[string]Listener)
	e.myDiary = File{Name: name}
	return e
}

func (e *Editor) OpenFile(path string) {
	e.myDiary = File{Name: path}
	e.Events.Notify("open", e.myDiary.Name)
}

func (e *Editor) SaveFile() {
	e.Events.Notify("save", e.myDiary.Name)
}

// Abstract Subscriber
type Listener interface {
	Update(string)
}

// Concrete Subscribers
type LoggingListener struct {
	Log     File
	Message string
}

func (l *LoggingListener) Update(message string) {
	l.Message = message
	fmt.Println("this is a log", l.Message)
}

type EmailAlertListener struct {
	Email   string
	Message string
}

func (e *EmailAlertListener) Update(message string) {
	e.Message = message
	fmt.Println("this is an alert email", e.Message)
}

// Main Processing Flow
func CreateSomeObservers() {
	editor := new(Editor).Init("null")

	logger := &LoggingListener{
		Log:     File{Name: "1"},
		Message: "null",
	}
	editor.Events.Subscribe("open", logger)

	emailAlert := &EmailAlertListener{
		Email:   "null",
		Message: "null",
	}
	editor.Events.Subscribe("save", emailAlert)

	editor.OpenFile("The Go Programming Language.")
	editor.SaveFile()
}
