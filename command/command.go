package main

import "fmt"

var overDraftLimit = -500

type BankAccount struct {
	balance int
}

func NewBankAccount(balance int) *BankAccount {
	return &BankAccount{balance: balance}
}

func (b *BankAccount) Print() {
	fmt.Println("Account balance is", b.balance)
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overDraftLimit {
		b.balance -= amount
		fmt.Println("Witdrew", amount, "\b, balance is now", b.balance)
		return true
	}
	return false
}

type Command interface {
	Call()
	Undo()
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account  *BankAccount
	action   Action
	amount   int
	succeded bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (cmd *BankAccountCommand) Call() {
	switch cmd.action {
	case Deposit:
		cmd.account.Deposit(cmd.amount)
		cmd.succeded = true
	case Withdraw:
		cmd.succeded = cmd.account.Withdraw(cmd.amount)
	}
}

func (cmd *BankAccountCommand) Undo() {
	if !cmd.succeded {
		return
	}

	switch cmd.action {
	case Withdraw:
		cmd.account.Deposit(cmd.amount)
	case Deposit:
		cmd.succeded = cmd.account.Withdraw(cmd.amount)
	}
}

func main() {
	account := NewBankAccount(1000)

	account.Print()

	cmd := NewBankAccountCommand(account, Withdraw, 500)
	cmd.Call()

	account.Print()

	cmd2 := NewBankAccountCommand(account, Deposit, 100)
	cmd2.Call()

	account.Print()

	cmd2.Undo()
	account.Print()
}
