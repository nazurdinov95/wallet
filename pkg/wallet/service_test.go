package wallet

import (
	"reflect"
	"testing"
	"github.com/nazurdinov95/wallet/pkg/types"
	//"github.com/google/uuid"
)

func TestService_Reject_success(t *testing.T)  {
	// создаём сервис
	s := &Service{}

	// регистрируем там пользователя
	phone := types.Phone("+992000000001")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("Reject(): cant register account, error = %v", err)
		return
	}

	//пополняем его счёт
	err = s.Deposit(account.ID, 10_000_00)
	if err != nil {
		t.Errorf("Reject(): cant deposit account, error = %v", err)
		return
	}

	// осуществляем платёж на его счёт
	payment, err := s.Pay(account.ID, 1000_00, "auto")
	if err != nil {
		t.Errorf("Reject(): cant create payment, error = %v", err)
		return
	}

	// пробуем отменить платёж
	err = s.Reject(payment.ID)
	if err != nil {
		t.Errorf("Reject(): error = %v", err)
		return
	}
}

func TestService_FindPaymentByID_success(t *testing.T)  {
	// создаём сервис
	s := &Service{}

	// регистрируем там пользователя
	phone := types.Phone("+992000000001")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("FindPaymentByID(): cant register account, error = %v", err)
		return
	}

	//пополняем его счёт
	err = s.Deposit(account.ID, 10_000_00)
	if err != nil {
		t.Errorf("FindPaymentByID(): cant deposit account, error = %v", err)
		return
	}

	// осуществляем платёж на его счёт
	payment, err := s.Pay(account.ID, 1000_00, "auto")
	if err != nil {
		t.Errorf("FindPaymentByID(): cant create payment, error = %v", err)
		return
	}

	// пробуем отменить платёж
	got, err := s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("FindPaymentByID(): error = %v", err)
		return
	}

	// сравниваем платежи
	if !reflect.DeepEqual(payment, got) {
		t.Errorf("FindPaymentByID(): wrong payment returned = %v", err)
		return
	}
}