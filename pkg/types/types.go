package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.).
type Money int64

// Category передставляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д.).
type PaymentCategory string

//Status представляет собой статус платежа.
type PaymentStatus string

//Передопределённые статусы платежей.
const (
	PaymentStatusOk 		PaymentStatus = "OK"
	PaymentStatusFail 		PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

// Payment presents information about the payment.
type Payment struct {
	ID 			string
	AccountID	int64
	Amount 		Money
	Category 	PaymentCategory
	Status		PaymentStatus
}

type Phone string

// Account представляет информацию о счете пользователя.
type Account struct {
	ID 		int64
	Phone 	Phone
	Balance Money
}