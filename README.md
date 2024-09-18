# Bank Transaction CLI

This Go program simulates a simple banking system where users can perform credit and debit transactions. The account balance is stored in a file (balance.txt) and is updated after each transaction, ensuring that the balance persists across multiple program executions.

## Features

    Credit: Add money to your account balance.
    Debit: Withdraw money from your account, provided you have sufficient balance.
    File Persistence: The balance is saved in a file (balance.txt), allowing it to be maintained between runs of the program.

Requirements

    Go 1.16 or later

Getting Started

1. Clone the repository:
```bash
git clone https://github.com/Anne-Okingo/Bank-Transaction-Comand-line-Interface.git
```


2. Set up your Go environment: Make sure you have Go installed and properly set up in your environment.

3. Run the program: The program accepts two arguments: the instruction (credit or debit) and the amount of money.
Example Commands:

    To Credit the account by 500:
```bash
go run . credit 500
```

    To Debit the account by 200:
```bash
        go run . debit 200
```

## How It Works

    The program reads the current balance from a file named balance.txt. If the file doesn't exist, it initializes the balance to 1000.
    Depending on the instruction (credit or debit), the program either:
        Credit: Adds the specified amount to the balance.
        Debit: Subtracts the specified amount from the balance, ensuring that the balance doesn’t go negative.
    After each transaction, the updated balance is written back to balance.txt.

### File-based Persistence

The account balance is stored in balance.txt. After every transaction, the balance is updated in the file so that the next time the program runs, it starts with the correct balance.

## Functions

    Credit: Adds the amount to the current balance.

    Debit: Deducts the amount from the current balance. If the balance after the transaction is negative, the programs shows you how have an inssufficient balance.

### Error Handling

    If the balance is insufficient for a debit transaction, the program will print an error message indicating insufficient funds.
    If the input amount is not a valid integer, the program will display an error.

### Example Output
Credit Example:

```bash
go run . credit 500
```
* Output

UserName: Aokingo
Account Number: 7536
Account Balance: Ksh 1500

Debit Example:
```bash
go run . debit 200
```

* Output

UserName: Aokingo
Account Number: 7536
Account Balance: Ksh 1300

Insufficient Funds Example:

bash

$ go run . debit 5000
UserName: Aokingo
Account Number: 7536
Insufficient funds for this debit transaction.
Account Balance: Ksh 1300

Customization

    Initial Balance: The initial balance is set to 1000 if balance.txt does not exist.
    Account Details: The UserName and AccountNo values are hardcoded for now. You can change these in the main.go file.

License

This project is open-source and free to use. Although Licenced under [aokingo](https://github.com/Anne-Okingo)