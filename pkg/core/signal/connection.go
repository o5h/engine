package signal

type Connection[T any] struct {
	signal *Signal[T]
	fn     func(T)
}

func (con *Connection[T]) Disconnect() {
	if con.signal == nil {
		return
	}
	con.signal.disconnect(con)
	con.signal = nil
}

func (con *Connection[T]) Value() T {
	return con.signal.Value
}
