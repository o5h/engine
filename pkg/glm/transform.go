package glm

var TransformIdent = Transform{
	Position: Vec3{0, 0, 0},
	Rotation: Vec3{0, 0, 0},
	Scale:    Vec3{1, 1, 1}}

type Transform struct {
	Position Vec3
	Rotation Vec3
	Scale    Vec3
}

func (t *Transform) Set(a *Transform) {
	t.Position.Set(&a.Position)
	t.Rotation.Set(&a.Rotation)
	t.Scale.Set(&a.Scale)
}

func (t *Transform) SetSum(a, b *Transform) {
	t.Position.SetSum(&a.Position, &b.Position)
	t.Rotation.SetSum(&a.Rotation, &b.Rotation)
	t.Scale.SetMul(&a.Scale, &b.Scale)
}
