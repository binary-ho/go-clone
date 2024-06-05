package banking

import (
	banking "clone-project/banking/account"
	"fmt"
)

func Run() {
	account := banking.CreateAccount("Jinho")
	fmt.Println(account)
}
