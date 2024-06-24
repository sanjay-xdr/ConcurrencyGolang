// A WaitGroup in Go is a simple way to wait for multiple goroutines to finish their execution. It’s like a counter: when you launch a goroutine, you increase the counter, and when a goroutine finishes, it decreases the counter. The Wait() function is used to block the execution of the program until the counter is zero, meaning all the goroutines have finished.
package examples

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func WaitGroupExample() {
	var wg sync.WaitGroup

	numbers := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}

	wg.Add(len(numbers))

	for i, x := range numbers {
		go printSomething(fmt.Sprintf("%d: %d", i+1, x), &wg)
	}

	wg.Wait()

	// time.Sleep(1 * time.Second)
	wg.Add(1)
	printSomething("This is second thing to be printed", &wg)
}
