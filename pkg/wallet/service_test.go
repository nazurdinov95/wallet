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

func TestService_Reject_fail(t *testing.T) {
	svc := Service{}

	svc.RegisterAccount("+992000000000")

	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Error(err)
	}

	err = svc.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Error(err)
	}

	payment, err := svc.Pay(account.ID, 100_00, "auto")
	if err != nil {
		t.Error(err)
	}

	pay, err := svc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Error(pay)
	}

	editPayID := "4"

	err = svc.Reject(editPayID)
	if err != ErrPaymentNotFound {
		t.Error(err)
	}
}



func TestService_Repeat_success(t *testing.T) {
	svc := &Service{}

	phone := types.Phone("+992000000000")

	account, err := svc.RegisterAccount(phone)
	if err != nil {
		t.Error(err)
		return
	}

	err = svc.Deposit(account.ID, 1000)
	if err != nil {
		t.Error(err)
		return
	}

	pay, err := svc.Pay(account.ID, 500, "auto")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = svc.Repeat(pay.ID)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestService_FavoritePayment_success(t *testing.T) {
	svc := &Service{}

	phone := types.Phone("+992000000000")

	account, err := svc.RegisterAccount(phone)
	if err != nil {
		t.Error(err)
		return
	}

	err = svc.Deposit(account.ID, 1000)
	if err != nil {
		t.Error(err)
		return
	}

	pay, err := svc.Pay(account.ID, 500, "auto")
	if err != nil {
		t.Error(err)
		return
	}

	favorite, err := svc.FavoritePayment(pay.ID, "pay")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(favorite)
}

func TestService_PayFromFavorite_success(t *testing.T) {
	svc := &Service{}

	phone := types.Phone("+992000000000")

	account, err := svc.RegisterAccount(phone)
	if err != nil {
		t.Error(err)
		return
	}

	err = svc.Deposit(account.ID, 1000)
	if err != nil {
		t.Error(err)
		return
	}

	pay, err := svc.Pay(account.ID, 500, "auto")
	if err != nil {
		t.Error(err)
		return
	}

	favorite, err := svc.FavoritePayment(pay.ID, "pay")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = svc.PayFromFavorite(favorite.ID)
	if err != nil {
		t.Error(err)
		return
	}

}