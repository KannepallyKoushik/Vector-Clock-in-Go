package main

import (
	"fmt"
	"os"
	"sync"
)

//waitgroups to let the main function wait until all goroutines are finished
var wg = sync.WaitGroup{}

//time vectors of each account
var timings1, receiver1 []int
var timings2, receiver2 []int
var timings3, receiver3 []int

//Initial balances of each account
var balance1 int = 500
var balance2 int = 700
var balance3 int = 1000

func main() {
	//channel to send and receive messages
	ch := make(chan []int)
	//Intializing the vector clock 1
	timings1 = append(timings1, 0)
	timings1 = append(timings1, 0)
	timings1 = append(timings1, 0)
	//Intializing the vector clock 2
	timings2 = append(timings2, 0)
	timings2 = append(timings2, 0)
	timings2 = append(timings2, 0)
	//Intializing the vector clock 3
	timings3 = append(timings3, 0)
	timings3 = append(timings3, 0)
	timings3 = append(timings3, 0)

	fmt.Println("Starting The Banking Application")
	for true {
		//Transfering the money or checking balance is considered to be an event
		fmt.Println("Enter your option")
		fmt.Println("1.Transfer Money from Accounts")
		fmt.Println("2.Check Balance")
		fmt.Println("3.Show the Synchronized timers in each process")
		fmt.Println("4.Exit Application")
		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			var sender, receiver string
			fmt.Println("Enter sender Account and Reciever Account names")
			fmt.Scanln(&sender, &receiver)
			//based on sender and receiver, appropriate goroutine is called.
			if sender == "A" {

				fmt.Println("Enter the amount you want to transfer into ", receiver, "'s Account")
				var money int
				fmt.Scanln(&money)
				if balance1-money > 0 {
					balance1 = balance1 - money
					//adding 2 waitgroups
					wg.Add(2)
					//calling sender and receivers respective goroutines
					go process1(ch, "transfer", 1)
					if receiver == "B" {
						balance2 = balance2 + money
						go process2(ch, "transfer", 0)
					} else if receiver == "C" {
						balance3 = balance3 + money
						go process3(ch, "transfer", 0)
					}
				} else {
					fmt.Println("Sorry you dont have enough Balance")
				}

			} else if sender == "B" {

				fmt.Println("Enter the amount you want to transfer into ", receiver, "'s Account")
				var money int
				fmt.Scanln(&money)
				if balance2-money > 0 {

					balance2 = balance2 - money
					//adding 2 waitgroups
					wg.Add(2)
					//calling sender and receivers respective goroutines
					go process2(ch, "transfer", 1)
					if receiver == "A" {
						balance1 = balance1 + money
						go process1(ch, "transfer", 0)
					} else if receiver == "C" {
						balance3 = balance3 + money
						go process3(ch, "transfer", 0)
					}
				} else {
					fmt.Println("Sorry you dont have enough Balance")
				}

			} else if sender == "C" {

				fmt.Println("Enter the amount you want to transfer into ", receiver, "'s Account")
				var money int
				fmt.Scanln(&money)

				if balance3-money > 0 {

					balance3 = balance3 - money
					//adding 2 waitgroups
					wg.Add(2)
					//calling sender and receivers respective goroutines
					go process3(ch, "transfer", 1)
					if receiver == "A" {
						balance1 = balance1 + money
						go process1(ch, "transfer", 0)
					} else if receiver == "B" {
						balance2 = balance2 + money
						go process2(ch, "transfer", 0)
					}
				} else {
					fmt.Println("Sorry you dont have enough Balance")
				}

			}
			//waiting for the goroutines to finish
			wg.Wait()
		case 2:
			var account string
			fmt.Println("Enter a Account name to check balance for that account")
			fmt.Scanln(&account)
			if account == "A" {
				wg.Add(1)
				go process1(ch, "Balance", 1)

			} else if account == "B" {
				wg.Add(1)
				go process2(ch, "Balance", 1)

			} else if account == "C" {
				wg.Add(1)
				go process3(ch, "Balance", 1)

			}
			//waiting for the goroutines to finish
			wg.Wait()
		case 3:
			//showing the synchorinsed clocks at a paricular instant
			wg.Add(3)
			go process1(ch, "Print", 1)
			go process2(ch, "Print", 1)
			go process3(ch, "Print", 1)
			//waiting for the goroutines to finish
			wg.Wait()
		case 4:
			os.Exit(3)
		}
		fmt.Println("------------------------------------------------")
	}
}

func process1(ch chan []int, whattodo string, senderflag int) {

	//if the event is to send or receive a message
	if whattodo == "transfer" {
		//if message sending
		if senderflag == 1 {
			timings1[0]++
			ch <- timings1

		} else { //if receiving message
			receiver1 = <-ch
			for i, val := range receiver1 {
				if val >= timings1[i] {
					timings1[i] = receiver1[i]
				}
			}

			timings1[0]++
		}

	} else if whattodo == "Balance" { //if the event is to check balance
		timings1[0]++
		fmt.Println("Your Balance is ", balance1)
	} else { //to check time at an instant
		fmt.Println("Timings from Process-1 is ", timings1)
	}
	//Indiacating a goroutine is completed
	wg.Done()
}

func process2(ch chan []int, whattodo string, senderflag int) {

	//if the event is to send or receive a message
	if whattodo == "transfer" {
		//if message sending
		if senderflag == 1 {
			timings2[1]++
			ch <- timings2
		} else { //if receiving message
			receiver2 = <-ch
			for i, val := range receiver2 {
				if val >= timings2[i] {
					timings2[i] = receiver2[i]
				}
			}
			timings2[1]++
		}

	} else if whattodo == "Balance" { //if the event is to check balance
		timings2[1]++
		fmt.Println("Your Balance is ", balance2)
	} else {
		fmt.Println("Timings from Process-2 is", timings2)
	}
	wg.Done()
}

func process3(ch chan []int, whattodo string, senderflag int) {
	//code regarding the sending or receiving a message
	if whattodo == "transfer" {
		//if message sending
		if senderflag == 1 {
			timings3[2]++
			ch <- timings3
		} else { //if receiving message
			receiver3 = <-ch
			for i, val := range receiver3 {
				if val >= timings3[i] {
					timings3[i] = receiver3[i]
				}
			}
			timings3[2]++
		}

	} else if whattodo == "Balance" {
		timings3[2]++
		fmt.Println("Your Balance is ", balance3)
	} else {
		fmt.Println("Timings from Process-3 is ", timings3)
	}
