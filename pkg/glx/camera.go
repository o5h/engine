package glx

import "github.com/o5h/engine/pkg/glm"

type Camera struct {
	Viewport glm.RectI

	Pos glm.Vec3
	Dir glm.Vec3
	Up  glm.Vec3

	FieldOfView float32
	Near        float32
	Far         float32

	View       glm.Mat4x4
	Projection glm.Mat4x4
}

func NewCamera() *Camera {
	camera := &Camera{}
	camera.init()
	return camera
}

func (camera *Camera) init() {
	camera.Pos = glm.Vec3{X: 0, Y: 0, Z: 0}
	camera.Dir = glm.Vec3{X: 0, Y: 1, Z: 0}
	camera.Up = glm.Vec3{X: 0, Y: 0, Z: 1}
	camera.FieldOfView = 60
	camera.Near = 0.001
	camera.Far = 10000
	camera.Viewport = glm.RectI{X: 0, Y: 0, W: 100, H: 100}
}

func (camera *Camera) GetView() glm.Mat4x4 { return camera.View }

func (camera *Camera) GetProjection() glm.Mat4x4 { return camera.Projection }

func (camera *Camera) UpdateView() {
	lookAt := camera.Pos
	lookAt.Add(camera.Dir)
	camera.View.LookAt(camera.Pos, lookAt, camera.Up)
}

func (camera *Camera) MoveForward(d float32) {
	tmp := camera.Dir
	tmp.Scale(d)
	camera.Pos.Add(tmp)
}

func (camera *Camera) MoveBackward(d float32) {
	tmp := camera.Dir
	tmp.Scale(d)
	camera.Pos.Sub(tmp)
}

func (camera *Camera) MoveLeft(d float32) {
	tmp := glm.CrossProduct(camera.Dir, camera.Up)
	tmp.Normalize()
	tmp.Scale(d)
	camera.Pos.Sub(tmp)
}

func (camera *Camera) MoveRight(d float32) {
	tmp := glm.CrossProduct(camera.Dir, camera.Up)
	tmp.Normalize()
	tmp.Scale(d)
	camera.Pos.Add(tmp)
}

func (camera *Camera) MoveUp(d float32) {
	tmp := camera.Up
	tmp.Scale(d)
	camera.Pos.Add(tmp)
}

func (camera *Camera) MoveDown(d float32) {
	tmp := camera.Up
	tmp.Scale(d)
	camera.Pos.Sub(tmp)
}

func (camera *Camera) UpdateProjection() {
	camera.Projection.Perspective(camera.Near, camera.Far, camera.FieldOfView, camera.Aspect())
}

func (camera *Camera) Aspect() float32 {
	return float32(camera.Viewport.W) / float32(camera.Viewport.H)
}

func (camera *Camera) RayCast(touch glm.Vec2) *glm.Ray {
	return glm.RayCast(
		glm.Vec2{X: float32(camera.Viewport.W), Y: float32(camera.Viewport.H)},
		touch,
		camera.Pos,
		&camera.Projection,
		&camera.View)
}
