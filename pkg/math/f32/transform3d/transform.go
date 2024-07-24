package transform3d

import "github.com/o5h/engine/pkg/math/f32/vec3"

var Ident = Transform3d{
	Position: vec3.Vec3{X: 0, Y: 0, Z: 0},
	Rotation: vec3.Vec3{X: 0, Y: 0, Z: 0},
	Scale:    vec3.Vec3{X: 1, Y: 1, Z: 1}}

type Transform3d struct {
	Position vec3.Vec3
	Rotation vec3.Vec3
	Scale    vec3.Vec3
}

func (t *Transform3d) GetPosition() vec3.Vec3         { return t.Position }
func (t *Transform3d) GetRotation() vec3.Vec3         { return t.Rotation }
func (t *Transform3d) GetScale() vec3.Vec3            { return t.Scale }
func (t *Transform3d) SetPosition(position vec3.Vec3) { t.Position = position }
func (t *Transform3d) SetRotation(rotation vec3.Vec3) { t.Rotation = rotation }
func (t *Transform3d) SetScale(scale vec3.Vec3)       { t.Scale = scale }

func (t *Transform3d) Set(a *Transform3d) {
	t.Position = a.Position
	t.Rotation = a.Rotation
	t.Scale = a.Scale
}

func (t *Transform3d) SetSum(a, b *Transform3d) {
	t.Position.SetSum(&a.Position, &b.Position)
	t.Rotation.SetSum(&a.Rotation, &b.Rotation)
	t.Scale.SetMul(&a.Scale, &b.Scale)
}
