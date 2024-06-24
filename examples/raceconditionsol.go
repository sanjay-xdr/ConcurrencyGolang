// Here we are using the mutex to lock and unlock the update function so that another go routine can not access it
package examples

import (
	"fmt"
	"sync"
)

func updateMsgSol(s string, mutex *sync.Mutex) {
	defer wg.Done()

	mutex.Lock()
	msg = s
	mutex.Unlock()
}
func RaceConditionSol() {

	var mutex sync.Mutex

	msg = "Hello, World"

	wg.Add(2)
	go updateMsgSol("Go is Great", &mutex)
	go updateMsgSol("Go is Good", &mutex)
	wg.Wait()

	fmt.Println(msg)

}

//Command to check race condition
// go run -race .
//In this the above command will print the answer
