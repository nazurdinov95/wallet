package main

import(
	"github.com/nazurdinov95/wallet/pkg/wallet"
)

func main() {
	svc := &wallet.Service{}
	svc.RegisterAccount("+992935308254")

	svc.Deposit(1, 10)
	svc.RegisterAccount("+992935338254")
}