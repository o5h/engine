package app

import (
	"github.com/o5h/engine/pkg/app/input/keyboard"
	"github.com/o5h/engine/pkg/app/input/mouse"
	"github.com/o5h/engine/pkg/core/signal"
)

type Context interface {
	Done()
	MouseEvents() *signal.Signal[mouse.Event]
	KeyboardEvents() *signal.Signal[keyboard.Event]
}
