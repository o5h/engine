package mouse

type Button int

type Event struct {
	Action Action
	X, Y   int
	Button Button
}

const (
	ButtonNone Button = iota
	ButtonLeft
	ButtonMiddle
	ButtonRight
)

type Action int

const (
	ActionNone Action = iota
	ActionPress
	ActionRelease
	ActionMove
)
