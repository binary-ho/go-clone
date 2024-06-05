package banking

import "errors"

type account struct {
	owner   string
	balance int
}

func CreateAccount(owner string) *account {
	newAccount := &account{owner: owner, balance: 0}
	updateAccount(newAccount)
	return newAccount
}

func (account account) Owner() string {
	return account.owner
}

func (account account) Balance() int {
	return account.balance
}

func (account *account) Deposit(amount int) {
	account.balance += amount
}

func (account *account) Withdraw(amount int) error {
	if account.balance < amount {
		return errors.New("돈이 없어요..")
	}
	account.balance -= amount
	return nil
}
