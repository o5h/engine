package glx

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/o5h/engine/internal/opengl/gl"
)

type VBOAttribute struct {
	Name       Attribute
	Size       int
	Format     Format
	Stride     int
	Offset     int
	Location   uint32
	Normalized bool //TODO: ??
}

func (a *VBOAttribute) Enable() {
	gl.EnableVertexAttribArray(a.Location)
	check()
	gl.VertexAttribPointer(a.Location, a.Size, uint32(a.Format), a.Normalized, a.Stride, uintptr(a.Offset))
	check()
}

type TVBO struct {
	SizeInBytes          int
	Usage                BufferUsage
	Attrs                []VBOAttribute
	Size                 int
	Data                 []byte
	AttributesByLocation map[uint32]VBOAttribute
	ID                   uint32
}

func BindVBO(usageMode BufferUsage, attrs []Attribute, vertices []Vertex) *TVBO {
	if sliceHasExact(attrs, Attribute_Position) {
		return createVBO[VertexP](usageMode, vertices)

	} else if sliceHasExact(attrs, Attribute_Position, Attribute_Normal) {
		return createVBO[VertexPN](usageMode, vertices)

	} else if sliceHasExact(attrs, Attribute_Position, Attribute_Normal) {
		return createVBO[VertexPUV](usageMode, vertices)

	} else if sliceHasExact(attrs, Attribute_Position, Attribute_Normal, Attribute_UV_0) {
		return createVBO[VertexPNUV](usageMode, vertices)
	}
	panic(fmt.Sprint("unsupported attribute combination", attrs))
}

func createVBO[T VBOComponent](usageMode BufferUsage, vertices []Vertex) *TVBO {
	size := len(vertices)
	data := AsByteBuffer(ConvertVBO[T](vertices))
	attrs, sizeInBytes := NewVBOAttributes[T]()
	return CreateVBO(usageMode, attrs, sizeInBytes, size, data)
}

func CreateVBO(usageMode BufferUsage, attrs []VBOAttribute, sizeInBytes, size int, data []byte) *TVBO {
	vbo := &TVBO{
		SizeInBytes: sizeInBytes,
		Usage:       usageMode,
		Attrs:       attrs,
		Size:        size,
		Data:        data}

	vbo.AttributesByLocation = make(map[uint32]VBOAttribute)
	for _, attr := range attrs {
		vbo.AttributesByLocation[attr.Location] = attr
	}

	vbo.ID = gl.GenBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.ID)
	switch usageMode {
	case StaticDraw:
		gl.BufferData(gl.ARRAY_BUFFER, len(vbo.Data), unsafe.Pointer(&vbo.Data[0]), uint32(vbo.Usage))
	case DynamicDraw:
		gl.BufferData(gl.ARRAY_BUFFER, cap(vbo.Data), nil, uint32(vbo.Usage))
	}
	return vbo
}

func (vbo *TVBO) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.ID)
	check()
}

func (vbo *TVBO) Delete() {
	gl.DeleteBuffer(vbo.ID)
	check()
}

func (vbo *TVBO) Enable(attrs []Attribute) {
	for _, attr := range attrs {
		a, ok := vbo.AttributesByLocation[attr.Location()]
		if !ok {
			log.Fatalln("Incompatible material, VBO should have  ", attr, " information")
		}
		a.Enable()
		check()
	}
}

func (vbo *TVBO) Draw(mode DrawMode, first, count int) {
	gl.DrawArrays(int(mode), first, count)
	check()

}

func (vbo *TVBO) UpdateData(offset, size int, v []byte) {
	gl.BufferSubData(gl.ARRAY_BUFFER, offset*vbo.SizeInBytes, size*vbo.SizeInBytes, unsafe.Pointer(&v[0]))
	check()
}

func VBOComponentSizeInBytes[T VBOComponent]() int {
	attrs := VBOComponentAttributes[T]()
	sizeInBytes := 0
	for _, attr := range attrs {
		sizeInBytes = sizeInBytes + attr.SizeInBytes()
	}
	return sizeInBytes
}

func NewVBOAttributes[T VBOComponent]() ([]VBOAttribute, int) {
	attrs := VBOComponentAttributes[T]()
	sizeInBytes := VBOComponentSizeInBytes[T]()
	vboAttributes := []VBOAttribute{}
	offset := 0
	for _, attr := range attrs {
		vboAttributes = append(vboAttributes, VBOAttribute{
			Name:       attr,
			Size:       attr.Size(),
			Format:     attr.Format(),
			Stride:     sizeInBytes,
			Offset:     offset,
			Location:   attr.Location(),
			Normalized: false})
		offset += attr.SizeInBytes()
	}
	return vboAttributes, sizeInBytes
}
