package glx

import (
	"log"

	"github.com/o5h/engine/internal/opengl/gl"
)

type FBO struct {
	W, H                int
	TextureID           uint32
	DepthRenderBufferID uint32
	ID                  uint32
}

func CreateFBO(w, h int) *FBO {
	fbo := &FBO{W: w, H: h}

	fbo.ID = gl.GenFramebuffer()
	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo.ID)

	fbo.TextureID = gl.GenTexture()
	gl.BindTexture(gl.TEXTURE_2D, fbo.TextureID)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, w, h, 0, gl.RGBA, gl.UNSIGNED_BYTE, 0)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, fbo.TextureID, 0)

	fbo.DepthRenderBufferID = gl.GenRenderbuffer()
	gl.BindRenderbuffer(gl.RENDERBUFFER, fbo.DepthRenderBufferID)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH_COMPONENT16, w, h)
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, gl.RENDERBUFFER, fbo.DepthRenderBufferID)

	status := gl.CheckFramebufferStatus(gl.FRAMEBUFFER)
	if status != gl.FRAMEBUFFER_COMPLETE {
		log.Println("FB not ready")
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.BindRenderbuffer(gl.RENDERBUFFER, 0)

	return fbo
}

func DeleteFBO(fbo *FBO) {
	gl.DeleteTexture(fbo.TextureID)
	gl.DeleteRenderbuffer(fbo.DepthRenderBufferID)
	gl.DeleteFramebuffer(fbo.ID)
}
