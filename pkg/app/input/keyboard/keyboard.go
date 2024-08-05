package keyboard

import "github.com/o5h/engine/signal"

type Event struct {
	Code      Code
	Direction Direction
	Rune      rune
}

type Direction uint8

const (
	None    Direction = 0
	Press   Direction = 1
	Release Direction = 2
)

var Events = signal.Signal[Event]{Name: "keyboard"}
