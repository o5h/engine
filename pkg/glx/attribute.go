package glx

import "fmt"

type Attribute string

var (
	Attribute_Position = Attribute("a_Position")
	Attribute_Normal   = Attribute("a_Normal")
	Attribute_Color    = Attribute("a_Color")
	Attribute_UV_0     = Attribute("a_UV_0")
	Attribute_Pos2     = Attribute("a_Pos2")
)

func (a Attribute) Location() uint32 {
	switch a {
	case Attribute_Position:
		return 0
	case Attribute_Normal:
		return 1
	case Attribute_Color:
		return 2
	case Attribute_UV_0:
		return 3
	case Attribute_Pos2:
		return 4
	}
	panic(fmt.Sprintf("unsupported attribute %v", a))
}

func (a Attribute) Format() Format {
	switch a {
	case Attribute_Position:
		return Format_Float32
	case Attribute_Normal:
		return Format_Float32
	case Attribute_Color:
		return Format_Float32
	case Attribute_UV_0:
		return Format_Float32
	case Attribute_Pos2:
		return Format_Float32
	}
	panic(fmt.Sprintf("unsupported attribute %v", a))
}

func (a Attribute) Size() int {
	switch a {
	case Attribute_Position:
		return 3
	case Attribute_Normal:
		return 3
	case Attribute_Color:
		return 3
	case Attribute_UV_0:
		return 2
	case Attribute_Pos2:
		return 2
	}
	panic(fmt.Sprintf("unsupported attribute %v", a))
}

func (a Attribute) SizeInBytes() int { return a.Format().SizeInBytes() * a.Size() }

var Attributes = []Attribute{
	Attribute_Position,
	Attribute_Normal,
	Attribute_Color,
	Attribute_UV_0,
	Attribute_Pos2}
