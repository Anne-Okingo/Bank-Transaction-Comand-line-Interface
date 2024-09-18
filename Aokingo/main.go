package main

import (
	"fmt"
	"os"
	"strconv"

	// "io/ioutil"
	"aokingo/functions"
)

// Function to write the balance to the file

func main() {
	UserName := "Aokingo"
	AccountNo := "7536"
	balanceFile := "balance.txt"

	// Check if the correct number of arguments is passed
	if len(os.Args) != 3 {
		fmt.Println("Error: usage should be go run . <instruction> <amount>")
		return
	}

	instructions := os.Args[1]
	str := os.Args[2]
	cash, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read the current balance from the file
	Initial_Balance, err := functions.ReadBalanceFromFile(balanceFile)
	if err != nil {
		// fmt.Print("Error reading balance:", err)
		// If the file does not exist, start with the initial balance of 1000
		Initial_Balance = 1000
	}

	// Process credit instruction
	if instructions == "credit" {
		balance := functions.Credit(Initial_Balance, cash)
		fmt.Println("UserName:", UserName)
		fmt.Println("Account Number:", AccountNo)

		Initial_Balance = balance // Update the initial balance after the transaction
		fmt.Println("Account Balance: Ksh", Initial_Balance)
	}

	// Process debit instruction
	if instructions == "debit" {
		balance := functions.Debit(Initial_Balance, cash)
		fmt.Println("UserName:", UserName)
		fmt.Println("Account Number:", AccountNo)

		if balance < 0 {
			fmt.Println("Insufficient funds for this debit transaction.")
		} else {
			Initial_Balance = balance // Update the initial balance after the transaction
		}

		fmt.Println("Account Balance: Ksh", Initial_Balance)
	}

	// Write the updated balance to the file
	functions.WriteBalanceToFile(balanceFile, Initial_Balance)
}
