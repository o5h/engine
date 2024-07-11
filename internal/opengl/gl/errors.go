package gl

import "errors"

var (
	ErrInvalidEnum                 = errors.New("invalid enum")
	ErrInvalidValue                = errors.New("invalid value")
	ErrInvalidOperation            = errors.New("invalid operation")
	ErrInvalidFrameBufferOperation = errors.New("invalid frame buffer operation")
	ErrOutOfMemory                 = errors.New("out of memory")
	ErrUnknown                     = errors.New("unknown error")
)
