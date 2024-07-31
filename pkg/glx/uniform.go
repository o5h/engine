package glx

import (
	"log"

	"github.com/o5h/engine/internal/opengl/gl"
	"github.com/o5h/engine/pkg/glm"
)

type Uniforms struct {
	Model               glm.Mat4x4
	View                glm.Mat4x4
	Projection          glm.Mat4x4
	ModelView           glm.Mat4x4
	ViewProjection      glm.Mat4x4
	ModelViewProjection glm.Mat4x4
	CameraLocation      glm.Vec3

	Color     Color
	Ambient   Color
	Difuse    Color
	Specular  Color
	Shininess float32
	Texture0  *Texture

	LightColor    Color
	LightLocation glm.Vec3
}

type uniform string
type uniformType int

const (
	ut_Mesh uniformType = iota
	ut_Material
	ut_Light
)

const (
	u_Model               = uniform("u_Model")
	u_View                = uniform("u_View")
	u_Projection          = uniform("u_Projection")
	u_ModelView           = uniform("u_ModelView")
	u_ViewProjection      = uniform("u_ViewProjection")
	u_ModelViewProjection = uniform("u_ModelViewProjection")
	u_CameraLocation      = uniform("u_CameraLocation")

	u_Color     = uniform("u_Color")
	u_Ambient   = uniform("u_Ambient")
	u_Difuse    = uniform("u_Difuse")
	u_Specular  = uniform("u_Specular")
	u_Shininess = uniform("u_Shininess")
	u_Texture0  = uniform("u_Texture0")

	u_LightColor    = uniform("u_LightColor")
	u_LightLocation = uniform("u_LightLocation")
)

func (u uniform) Type() uniformType {
	switch u {

	case u_Model,
		u_View,
		u_Projection,
		u_ModelView,
		u_ViewProjection,
		u_ModelViewProjection,
		u_CameraLocation:
		return ut_Mesh

	case u_Color,
		u_Ambient,
		u_Difuse,
		u_Specular,
		u_Shininess,
		u_Texture0:
		return ut_Material

	case u_LightColor,
		u_LightLocation:
		return ut_Light
	}
	log.Fatalln("unsupported uniform", u)
	panic("unsupported uniform")
}

var uniformBindings = map[uniform]UniformBinding{
	u_Model:               func(u *Uniforms, loc int) { gl.UniformMatrix4fv(loc, 1, false, u.Model.Ptr()) },
	u_View:                func(u *Uniforms, loc int) { gl.UniformMatrix4fv(loc, 1, false, u.View.Ptr()) },
	u_Projection:          func(u *Uniforms, loc int) { gl.UniformMatrix4fv(loc, 1, false, u.Projection.Ptr()) },
	u_ModelView:           func(u *Uniforms, loc int) { gl.UniformMatrix4fv(loc, 1, false, u.ModelView.Ptr()) },
	u_ViewProjection:      func(u *Uniforms, loc int) { gl.UniformMatrix4fv(loc, 1, false, u.ViewProjection.Ptr()) },
	u_ModelViewProjection: func(u *Uniforms, loc int) { gl.UniformMatrix4fv(loc, 1, false, u.ModelViewProjection.Ptr()) },
	u_CameraLocation: func(u *Uniforms, loc int) {
		gl.Uniform3f(loc, u.CameraLocation.X, u.CameraLocation.Y, u.CameraLocation.Z)
	},
	u_Color:     func(u *Uniforms, loc int) { gl.Uniform3f(loc, u.Color.R, u.Color.G, u.Color.B) },
	u_Ambient:   func(u *Uniforms, loc int) { gl.Uniform3f(loc, u.Ambient.R, u.Ambient.G, u.Ambient.B) },
	u_Difuse:    func(u *Uniforms, loc int) { gl.Uniform3f(loc, u.Ambient.R, u.Ambient.G, u.Ambient.B) },
	u_Specular:  func(u *Uniforms, loc int) { gl.Uniform3f(loc, u.Ambient.R, u.Ambient.G, u.Ambient.B) },
	u_Shininess: func(u *Uniforms, loc int) { gl.Uniform1f(loc, u.Shininess) },

	u_Texture0:      func(u *Uniforms, loc int) { applyTexture(u.Texture0, 0) },
	u_LightColor:    func(u *Uniforms, loc int) { gl.Uniform3f(loc, u.LightColor.R, u.LightColor.G, u.LightColor.B) },
	u_LightLocation: func(u *Uniforms, loc int) { gl.Uniform3f(loc, u.LightLocation.X, u.LightLocation.Y, u.LightLocation.Z) },
}

func applyTexture(t *Texture, i int) {
	gl.ActiveTexture(uint32(gl.TEXTURE0 + i))
	gl.BindTexture(gl.TEXTURE_2D, t.ID)
}
