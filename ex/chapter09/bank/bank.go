// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package bank provides a concurrency-safe bank with one account.
package bank

type withdrawReqType struct {
	amount int
	result chan<- bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawals = make(chan withdrawReqType) // request a withdrawal

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {

	result := make(chan bool)
	withdrawals <- withdrawReqType{
		amount: amount,
		result: result,
	}

	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdrawReq := <-withdrawals:
			if balance > withdrawReq.amount {
				balance -= withdrawReq.amount
				withdrawReq.result <- true
			} else {
				withdrawReq.result <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
