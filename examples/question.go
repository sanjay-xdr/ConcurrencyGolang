package examples

import (
	"fmt"
	"sync"
)

/*


variable for bank balance
print out starting values
define weekly revenue

loop through 52 weeks and print out how much is made

print the final balance



*/

type Income struct {
	Source string
	Amount int
}

// var wg sync.WaitGroup

func FindbankBalance() {

	var bankbalance int

	var mutex sync.Mutex

	fmt.Printf("Initial Account Balance: %d.00", bankbalance)
	fmt.Println()

	var weeklyRevenue = []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Second Job", Amount: 100},
		{Source: "Investment ", Amount: 80},
		{Source: "Saas Tool", Amount: 500},
	}

	wg.Add(len(weeklyRevenue))

	for i, income := range weeklyRevenue {

		go func(i int, income Income, mutext *sync.Mutex) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				mutext.Lock()
				temp := bankbalance
				temp += income.Amount
				bankbalance = temp
				mutext.Unlock()
				fmt.Printf("On week %d ,You earned %d.00 from %s\n", week, income.Amount, income.Source)
			}

		}(i, income, &mutex)

	}
	wg.Wait()
	fmt.Print("FInal Bank balance")
	println()
	fmt.Print(bankbalance)

}
