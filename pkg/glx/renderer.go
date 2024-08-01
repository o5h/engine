package glx

import (
	"errors"
	"image"
	"image/draw"
	"log"

	"github.com/o5h/engine/pkg/app"
	"github.com/o5h/glm"
	"github.com/o5h/opengles/gl"
)

type Renderer struct {
	C             *app.Context
	U             Uniforms
	ActiveProgram *Program
}

func (rnd *Renderer) StartFrame(cam *Camera, c Color) {
	//rnd.U.Viewport = cam.Viewport
	rnd.U.CameraLocation = cam.Pos
	cam.UpdateView()
	cam.UpdateProjection()
	rnd.SetViewProjection(cam.View, cam.Projection)
	gl.Viewport(cam.Viewport.X, cam.Viewport.Y, cam.Viewport.W, cam.Viewport.H)
	gl.Enable(gl.DEPTH_TEST)
	gl.Disable(gl.CULL_FACE)
	gl.ClearColor(c.R, c.G, c.B, c.A)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (rnd *Renderer) EndFrame() {
	//TODO:swap buffers
}

func (rnd *Renderer) SetViewProjection(view, proj glm.Mat4x4) {
	rnd.U.View = view
	rnd.U.Projection = proj
	rnd.U.ViewProjection.SetMul(proj, view)
}

func (rnd *Renderer) SetModel(model glm.Mat4x4) {
	rnd.U.Model = model
	rnd.U.ModelView.SetMul(model, rnd.U.View)
	rnd.U.ModelViewProjection.SetMul(rnd.U.ViewProjection, model)
}

func (rnd *Renderer) BeginFrameBuffer(fbo *FBO, bgColor Color) {
	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo.ID)
	gl.ClearColor(bgColor.R, bgColor.G, bgColor.B, bgColor.A)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) // we're not using the stencil buffer now
	gl.Enable(gl.DEPTH_TEST)
}

func (rnd *Renderer) EndFrameBuffer() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (rnd *Renderer) ReadImage(w, h int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	gl.ReadPixels(0, 0, w, h, gl.RGBA, gl.UNSIGNED_BYTE, img.Pix)
	return img
}

// func (rnd *Renderer) Draw(ibo *IBO, vbo *VBO, mode DrawMode) {
// 	rnd.EnableVBO(vbo, rnd.activeProgram.Attribs)
// 	for _, uniform := range rnd.activeProgram.ModelUniforms {
// 		uniform.Bind(&rnd.U, uniform.Location)
// 	}
// 	ibo.Draw(mode)
// }

// func (rnd *Renderer) EnableVBOAttribute(a VBOAttribute) {
// 	gl.EnableVertexAttribArray(a.Location)
// 	gl.VertexAttribPointer(a.Location, a.Size, uint32(a.Format), a.Normalized, a.Stride, uintptr(a.Offset))
// }

func (rnd *Renderer) UseProgram(p *Program) {
	rnd.ActiveProgram = p
	gl.UseProgram(p.ID)
	for _, uniform := range p.MaterialUniforms {
		uniform.Bind(&rnd.U, uniform.Location)
	}
}

func (rnd *Renderer) DeleteProgram(p *Program) {
	gl.DeleteProgram(p.ID)
	gl.DeleteShader(p.VertexShaderId)
	gl.DeleteShader(p.FragmentShaderId)
}

func compileShader(src string, typ uint32) (uint32, error) {
	id := gl.CreateShader(uint32(typ))
	if !gl.IsShader(id) {
		log.Fatal(id)
		panic("nor shader")
	}
	gl.ShaderSource(id, src)
	gl.CompileShader(id)
	if gl.GetShaderiv(id, gl.COMPILE_STATUS) == gl.FALSE {
		msg := gl.GetShaderInfoLog(id)
		log.Println("ERROR:", msg)
		return 0, errors.New(msg)
	}
	return id, nil
}

func ConvertImage(f TextureFormat, img image.Image) (b []byte) {
	switch f {
	case TextureFormat_RGBA:
		b = convertRGBA(img)
	default:
		panic("Unsupported")
	}
	return
}

func convertRGBA(img image.Image) []byte {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	w := bounds.Dx()
	h := bounds.Dy()
	bb := make([]byte, 0, w*h*4)
	for dy := h - 1; dy >= 0; dy-- {
		s := dy * w * 4
		bb = append(bb, rgba.Pix[s:s+w*4]...)
	}
	return bb
}

func check() {
	err := gl.GetError()
	if err != nil {
		panic(err)
	}
}
