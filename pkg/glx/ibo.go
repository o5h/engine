package glx

import (
	"fmt"
	"unsafe"

	"github.com/o5h/opengles/gl"
)

type IBO struct {
	UsageMode BufferUsage
	Format    Format
	Size      int
	Data      []byte
	ID        uint32
}

func BindIBO(usage BufferUsage, indices []int) *IBO {
	format := defineIBOFormat(indices)
	switch format {
	case Format_UInt8:
		data := make([]uint8, 0, len(indices))
		for _, i := range indices {
			data = append(data, uint8(i))
		}
		return CreateIBO(usage, ComponentFormat[uint8](), len(indices), AsByteBuffer(data))
	case Format_UInt16:
		data := make([]uint16, 0, len(indices))
		for _, i := range indices {
			data = append(data, uint16(i))
		}

		return CreateIBO(usage, ComponentFormat[uint16](), len(indices), AsByteBuffer(data))
	}
	panic(fmt.Sprint("unsupported index format ", format))
}

func CreateIBO(usage BufferUsage, format Format, size int, data []byte) *IBO {
	ibo := &IBO{
		UsageMode: usage,
		Format:    format,
		Size:      size,
		Data:      data,
		ID:        gl.GenBuffer()}
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo.ID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(ibo.Data), unsafe.Pointer(&ibo.Data[0]), uint32(ibo.UsageMode))
	return ibo
}

func (b *IBO) Draw(mode DrawMode) {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, b.ID)
	gl.DrawElements(uint32(mode), int(b.Size), uint32(b.Format), 0)
}

func (ibo *IBO) Delete() { gl.DeleteBuffer(ibo.ID) }

func defineIBOFormat(indices []int) Format {
	max := 0
	for _, i := range indices {
		if max < i {
			max = i
		}
	}
	if max < 255 {
		return Format_UInt8
	} else if max < 65535 {
		return Format_UInt16
	} else {
		panic(fmt.Sprint("index is too high", max))
	}
}
