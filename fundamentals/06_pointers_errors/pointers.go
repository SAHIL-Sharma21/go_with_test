package main

import "fmt"

func main() {
	fmt.Println("testing with pointers and errors")

	// wallet := Wallet{
	// 	balance: 20,
	// }

	// fmt.Println(wallet.Balance())
	// wallet.Deposit(30)
	// fmt.Println(wallet.Balance())

}

//learning pointers and error wioth test case

/*
-------------------------------------------
Fintech loves Go and uhhh bitcoins? So let's show what an amazing banking system we can make.

Let's make a Wallet struct which lets us deposit Bitcoin.
-------------------------------------------
*/

// creating our type
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// when you call a function or a method the arguments are copied.
// to solve this we will use pointer in this as it will give give the original value and change the original not the copy one
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("amount of balance is %p\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// implementing stringer
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
