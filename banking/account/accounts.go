package banking

import "errors"

type accounts map[string]account

var accountsInstance *accounts

func getAccounts() *accounts {
	if accountsInstance == nil {
		accountsInstance = &accounts{}
	}
	return accountsInstance
}

func updateAccount(account *account) {
	(*getAccounts())[account.owner] = *account
}

func removeAccount(account *account) {
	delete(*getAccounts(), account.owner)
}

func GetAccount(owner string) (*account, error) {
	result, exsist := (*getAccounts())[owner]
	if !exsist {
		return nil, errors.New("계좌가 없습니다.")
	}
	return &result, nil
}
