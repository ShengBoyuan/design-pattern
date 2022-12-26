package singleton

import (
	"fmt"
	"sync"
)

// Singleton
type Database struct {
	connection string // I use 'string' to represent real dbConn here.
	mutex      sync.Locker
}

func (d *Database) initialize() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.connection == "" {
		d.connection = "ready"
		fmt.Println("dbConn init finished")
	} else {
		fmt.Println("dbConn has been initialized as", d.connection)
	}
}

func (d *Database) query(sql string) {
	if d.connection == "ready" {
		fmt.Println("sql exec OK")
	} else {
		fmt.Println("dbConn hasn't been initialized!")
	}
}

// Main Processing Flow
func GetSomeData() {
	var wg sync.WaitGroup
	dbConn := &Database{}

	dbConn.query("select * from users")

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			dbConn.initialize()
			wg.Done()
		}()
	}

	wg.Wait()
	dbConn.query("select * from users")
}
