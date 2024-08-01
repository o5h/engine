package glx

import (
	"unsafe"

	"github.com/o5h/opengles/gl"
)

type Component interface {
	int8 | uint8 | int16 | uint16 | uint32 | float32
}

type BufferComponent interface {
	IBOComponent | VBOComponent
}

type IBOComponent interface{ uint8 | uint16 }
type VBOComponent interface {
	VertexP | VertexPN | VertexPUV | VertexPNUV
}

func AsComponentBuffer[T Component](buf []byte, len uint32) []T {
	return unsafe.Slice((*T)(unsafe.Pointer(&buf[0])), len)
}

func AsByteBuffer[T BufferComponent](buf []T) []byte {
	size := len(buf) * BufferComponentSizeInBytes[T]()
	if size == 0 {
		return []byte{}
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(&buf[0])), size)
}

func VBOComponentAttributes[T VBOComponent]() []Attribute {
	var t T
	switch any(t).(type) {
	case VertexP:
		return []Attribute{Attribute_Position}
	case VertexPN:
		return []Attribute{Attribute_Position, Attribute_Normal}
	case VertexPUV:
		return []Attribute{Attribute_Position, Attribute_UV_0}
	case VertexPNUV:
		return []Attribute{Attribute_Position, Attribute_Normal, Attribute_UV_0}
	}
	return []Attribute{}
}

func BufferComponentSizeInBytes[T BufferComponent]() int {
	var t T
	switch any(t).(type) {
	default:
		return 1
	case int16, uint16:
		return 2
	case uint32, float32:
		return 4
	case VertexP:
		return VBOComponentSizeInBytes[VertexP]()
	case VertexPN:
		return VBOComponentSizeInBytes[VertexPN]()
	case VertexPUV:
		return VBOComponentSizeInBytes[VertexPUV]()
	case VertexPNUV:
		return VBOComponentSizeInBytes[VertexPNUV]()
	}
}

func ComponentFormat[T Component]() Format {
	var t T
	switch any(t).(type) {
	default:
		return gl.BYTE
	case uint8:
		return gl.UNSIGNED_BYTE
	case int16:
		return gl.SHORT
	case uint16:
		return gl.UNSIGNED_SHORT
	case uint32:
		return gl.UNSIGNED_INT
	case float32:
		return gl.FLOAT
	}
}

type Format uint32

const (
	Format_Int8    Format = gl.BYTE
	Format_UInt8   Format = gl.UNSIGNED_BYTE
	Format_Int16   Format = gl.SHORT
	Format_UInt16  Format = gl.UNSIGNED_SHORT
	Format_UInt32  Format = gl.UNSIGNED_INT
	Format_Float32 Format = gl.FLOAT
)

func (f Format) SizeInBytes() int {
	switch f {
	case Format_Int8:
		return 1
	case Format_UInt8:
		return 1
	case Format_Int16:
		return 2
	case Format_UInt16:
		return 2
	case Format_UInt32:
		return 4
	case Format_Float32:
		return 4
	}
	panic("Unsupported format ")
}
