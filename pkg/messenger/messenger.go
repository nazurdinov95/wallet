package messenger

type Messenger interface {
	Send(message string)bool
	Receiver()(message string, ok bool)
}