package handler

import "CZ3004-RPi/src/message"

// Handler is an adapter to wrap receieve methods
type Handler func(r message.Request)

// Handle is an adapter method
func (h Handler) Handle(r message.Request) {
	h(r)
}

// Receiver ...
type Receiver func(m message.Message) (n int, e error)

// Receive ...
func (r Receiver) Receive(m message.Message) (n int, e error) {
	n, e = r(m)
	if e != nil {
		return 0, e
	}
	return n, nil
}
