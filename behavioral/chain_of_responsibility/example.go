package cor

import "fmt"

// Handler
type Waiter interface {
	ServeCustomer()
}

// Base Handler
type BaseWaiter struct {
	NextOne Waiter
}

func (b *BaseWaiter) ServeCustomer() {
	fmt.Println("I just provide some basic services")
}

// Concrete Handler
type WineWaiter struct {
	NextOne Waiter
}

func (w *WineWaiter) ServeCustomer() {
	fmt.Println("I serve wine to the customers")
}

// Container
// The chain relationships are established here.
type WaiterQueue struct {
	Waiters []Waiter
}

func (w *WaiterQueue) AddWaiter(newWaiter Waiter) {
	w.Waiters = append(w.Waiters, newWaiter)
}

func (w *WaiterQueue) ServeCustomer() {
	for _, t := range w.Waiters {
		t.ServeCustomer()
	}
}

// Client
type Customer struct {
	Service *WaiterQueue
}

func (c *Customer) EnjoyService() {
	c.Service.ServeCustomer()
}

// Main Processing Flow
func GoToTheRestaurant() {
	waiter1 := &BaseWaiter{}
	waiter2 := &WineWaiter{}
	waiterQueue := &WaiterQueue{}
	waiterQueue.AddWaiter(waiter1)
	waiterQueue.AddWaiter(waiter2)

	customer := &Customer{Service: waiterQueue}
	customer.EnjoyService()
}
