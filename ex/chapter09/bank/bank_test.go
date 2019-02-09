// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		res := Withdraw(200)
		fmt.Printf("Withdrawal: %v\n", res)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		res := Withdraw(200)
		fmt.Printf("Withdrawal: %v\n", res)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done
	<-done

	if got, want := Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
