package signal

import "slices"

type Signal[T any] struct {
	Name        string
	Value       T
	connections []*Connection[T]
}

func (s *Signal[T]) Next(value T) {
	s.Value = value
	for _, con := range s.connections {
		con.fn(value)
	}
}

func (s *Signal[T]) Connect(fn func(T)) *Connection[T] {
	con := &Connection[T]{signal: s, fn: fn}
	s.connections = append(s.connections, con)
	return con
}

func (s *Signal[T]) disconnect(con *Connection[T]) {
	s.connections = slices.DeleteFunc(s.connections, func(c *Connection[T]) bool {
		return con == c
	})
}
