package eventhandler

import (
	"log"
	"os"

	"golang.org/x/term"
)

type handler struct {
	active bool
}

func NewHandler() *handler {
	return &handler{active: true}
}

func (h *handler) Run() {
	h.active = true

	for h.active {
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			log.Fatal("Unable to handle input: ", err)
		}

		defer term.Restore(int(os.Stdin.Fd()), oldState)

		event := make([]byte, 1)
		_, err = os.Stdin.Read(event)
		if err != nil {
			log.Fatal("Unable to return input: ", err)
		}
		log.Printf("This is the character: %q", string(event[0]))
		term.Restore(0, oldState)
	}
}

func (h *handler) Kill() {
	h.active = false
	h = nil
}
