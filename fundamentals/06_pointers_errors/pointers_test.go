package main

import (
	"testing"
)

// testing for banking system
// func TestWallet(t *testing.T) {
// 	wallet := Wallet{}

// 	wallet.Deposit(Bitcoin(10))
// 	got := wallet.Balance()
// 	fmt.Printf("Address of balance in test is %p\n", &wallet.balance)
// 	want := Bitcoin(10)

// 	if got != want {
// 		t.Errorf("got %s want %s", got, want)
// 	}
// }

/*
Aftet using pointer we get the same address
amount of balance is 0xc0001020e8
Address of balance in test is 0xc0001020e8
*/

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s and want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s and want %s", got, want)
		}
	})
}
