package glx

import "github.com/o5h/glm"

type Vertex struct {
	Position glm.Vec3
	Normal   glm.Vec3
	Color    glm.Vec3
	UV       glm.Vec2
}

type IVertex interface {
	Attributes() []Attribute
}

type VertexP2UV struct {
	Pos2 glm.Vec2
	UV   glm.Vec2
}

func (VertexP2UV) Attributes() []Attribute { return []Attribute{Attribute_Pos2, Attribute_UV_0} }

type VertexP struct {
	Position glm.Vec3
}

type VertexPN struct {
	Position glm.Vec3
	Normal   glm.Vec3
}
type VertexPUV struct {
	Position glm.Vec3
	UV       glm.Vec2
}

type VertexPNUV struct {
	Position glm.Vec3
	Normal   glm.Vec3
	UV       glm.Vec2
}

type VertexPNCUV struct {
	Position glm.Vec3
	Normal   glm.Vec3
	Color    glm.Vec3
	UV       glm.Vec2
}

func (VertexP) Attributes() []Attribute   { return []Attribute{Attribute_Position} }
func (VertexPN) Attributes() []Attribute  { return []Attribute{Attribute_Position, Attribute_Normal} }
func (VertexPUV) Attributes() []Attribute { return []Attribute{Attribute_Position, Attribute_UV_0} }
func (VertexPNUV) Attributes() []Attribute {
	return []Attribute{Attribute_Position, Attribute_Normal, Attribute_UV_0}
}

func (VertexPNUV) Compact(v Vertex) VertexPNUV {
	return VertexPNUV{Position: v.Position, Normal: v.Normal, UV: v.UV}
}

func (v *VertexP) Set(vx Vertex) {
	v.Position = vx.Position
}

func (v *VertexPNUV) Set(vx Vertex) {
	v.Position = vx.Position
	v.Normal = vx.Normal
	v.UV = vx.UV
}

func (v *VertexPNCUV) Set(vx Vertex) {
	v.Position = vx.Position
	v.Normal = vx.Normal
	v.Color = vx.Color
	v.UV = vx.UV
}

func (VertexPNCUV) Attributes() []Attribute {
	return []Attribute{Attribute_Position, Attribute_Normal, Attribute_Color, Attribute_UV_0}
}

func VertexPNEqEpsilon(a, b *VertexPN) bool {
	if a.Position.EqEpsilon(b.Position) && a.Normal.EqEpsilon(b.Normal) {
		return true
	}
	return false
}

// func (v *VertexPNUV) EqEpsilon(o *VertexPNUV) bool {
// 	if v.Position.EqEpsilon(&o.Position) &&
// 		v.Normal.EqEpsilon(&o.Normal) &&
// 		v.UV.EqEpsilon(&o.UV) {
// 		return true
// 	}
// 	return false
// }

// func (v *VertexPNCUV) EqEpsilon(o *VertexPNCUV) bool {
// 	if v.Position.EqEpsilon(&o.Position) &&
// 		v.Normal.EqEpsilon(&o.Normal) &&
// 		v.Color.EqEpsilon(&o.Color) &&
// 		v.UV.EqEpsilon(&o.UV) {
// 		return true
// 	}
// 	return false
// }

func ConvertVBO[V VBOComponent](vertices []Vertex) []V {
	data := make([]V, 0, len(vertices))
	for _, vx := range vertices {
		var v V
		switch vv := ((any)(&v)).(type) {
		case *VertexP:
			vv.Set(vx)
		case *VertexPNUV:
			vv.Set(vx)
		case *VertexPNCUV:
			vv.Set(vx)
		}
		data = append(data, v)
	}
	return data
}
