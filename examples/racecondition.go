//A race condition occurs when multiple threads try to access and modify the same data (memory address). E.g., if one thread tries to increase an integer and another thread tries to read it, this will cause a race condition

package examples

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMsg(s string) {
	defer wg.Done()
	msg = s
}
func RaceCondition() {

	msg = "Hello, World"

	wg.Add(2)
	go updateMsg("Go is Great") //this trying to update the msg
	go updateMsg("Go is Good")  // this as well
	wg.Wait()

	fmt.Println(msg)

}

//Command to check race condition
// go run -race .
//In this the above command will print the warning regarding the race condition
