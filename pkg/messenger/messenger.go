package messenger

type Messenger interface {
	Send(message string)bool
	Receiver()(message string, ok bool)
}

type Telegram struct {

}

func (t *Telegram) Send(message string) bool {
	return true
}

func (t *Telegram) Reseive() (message string, ok bool) {
	return "", true
}
