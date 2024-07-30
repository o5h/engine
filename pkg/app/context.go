package app

import (
	"github.com/o5h/engine/pkg/app/input/keyboard"
	"github.com/o5h/engine/pkg/app/input/mouse"
	"github.com/o5h/engine/pkg/core/rx"
)

type Context interface {
	Done()
	MouseEvents() *rx.Subject[mouse.Event]
	KeyboardEvents() *rx.Subject[keyboard.Event]
}
