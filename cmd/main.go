package main

import(
	"github.com/nazurdinov95/wallet/pkg/wallet"
	"fmt"
)

func main() {
	svc := &wallet.Service{}
	account, err := svc.RegisterAccount("+992935308254")
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(account)
	err = svc.Deposit(account.ID, 10)
	if err != nil {
		switch err {
		case wallet.ErrAmountMustBePositive:
			fmt.Println("Сумма должна быть положительной")
		case wallet.ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
		return
	}

	fmt.Println(account.Balance) //10
}