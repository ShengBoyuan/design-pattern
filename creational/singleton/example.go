package singleton

import (
	"fmt"
	"sync"
)

// Singleton: Implementation 1
// Use sync.Mutex as a simple Lock to control cocurrency.
type Database struct {
	connection string // I use 'string' to represent real dbConn here.
	mutex      sync.Mutex
}

func (d *Database) initialize() {
	if d.connection == "" {
		d.mutex.Lock()
		defer d.mutex.Unlock()
		if d.connection == "" {
			// Ensure that the instance hasn't yet been initialized by another thread
			// while this one has been waiting for the lock's release.
			d.connection = "ready"
			fmt.Println("dbConn init finished")
		} else {
			fmt.Println("dbConn has been initialized as", d.connection)
		}
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

// Singleton: Implementation 2
// Directly use sync.Once to ensure func 'getInstance()' execute only once.
type Database2 struct {
	connection string
	once       sync.Once
}

func (d *Database2) getInstance() *Database2 {
	d.once.Do(func() {
		if d.connection == "" {
			d.connection = "ready"
			fmt.Println("dbConn initialization finished")
		}
	})
	return d
}

func (d *Database2) query(sql string) {
	if d.connection == "ready" {
		fmt.Println("sql exec OK")
	} else {
		fmt.Println("dbConn hasn't been initialized!")
	}
}

// Main Processing Flow
func GetSomeData2() {
	var wg sync.WaitGroup
	dbConn := &Database2{}

	dbConn.query("select * from users")

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			dbConn = dbConn.getInstance()
			wg.Done()
		}()
	}

	wg.Wait()
	dbConn.query("select * from users")
}
