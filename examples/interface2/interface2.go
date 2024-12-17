package main

import "fmt"

type Account struct {
	name string
	mail string
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

type AccountNotifier interface {
	NotifyAccountCreated(Account) error
}

func NotifyAccountCreated(account Account) error {
	fmt.Println("account created")
	return nil
}

func (h *AccountHandler) handleCreateAccount() {
	account := Account{name: "bob", mail: "bob@hotail.com"}

	if err := h.AccountNotifier.NotifyAccountCreated(account); err != nil {
		panic("aa")
	}
}

func main() {
}
