//go:build windows

package gl

// #include <stdlib.h>
import "C"

import (
	"math"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	dll                                     *windows.DLL
	procActiveTexture                       *windows.Proc
	procAttachShader                        *windows.Proc
	procBindAttribLocation                  *windows.Proc
	procBindBuffer                          *windows.Proc
	procBindFramebuffer                     *windows.Proc
	procBindRenderbuffer                    *windows.Proc
	procBindTexture                         *windows.Proc
	procBlendColor                          *windows.Proc
	procBlendEquation                       *windows.Proc
	procBlendEquationSeparate               *windows.Proc
	procBlendFunc                           *windows.Proc
	procBlendFuncSeparate                   *windows.Proc
	procBufferData                          *windows.Proc
	procBufferSubData                       *windows.Proc
	procCheckFramebufferStatus              *windows.Proc
	procClear                               *windows.Proc
	procClearColor                          *windows.Proc
	procClearDepthf                         *windows.Proc
	procClearStencil                        *windows.Proc
	procColorMask                           *windows.Proc
	procCompileShader                       *windows.Proc
	procCompressedTexImage2D                *windows.Proc
	procCompressedTexSubImage2D             *windows.Proc
	procCopyTexImage2D                      *windows.Proc
	procCopyTexSubImage2D                   *windows.Proc
	procCreateProgram                       *windows.Proc
	procCreateShader                        *windows.Proc
	procCullFace                            *windows.Proc
	procDeleteBuffers                       *windows.Proc
	procDeleteFramebuffers                  *windows.Proc
	procDeleteProgram                       *windows.Proc
	procDeleteRenderbuffers                 *windows.Proc
	procDeleteShader                        *windows.Proc
	procDeleteTextures                      *windows.Proc
	procDepthFunc                           *windows.Proc
	procDepthMask                           *windows.Proc
	procDepthRangef                         *windows.Proc
	procDetachShader                        *windows.Proc
	procDisable                             *windows.Proc
	procDisableVertexAttribArray            *windows.Proc
	procDrawArrays                          *windows.Proc
	procDrawElements                        *windows.Proc
	procEnable                              *windows.Proc
	procEnableVertexAttribArray             *windows.Proc
	procFinish                              *windows.Proc
	procFlush                               *windows.Proc
	procFramebufferRenderbuffer             *windows.Proc
	procFramebufferTexture2D                *windows.Proc
	procFrontFace                           *windows.Proc
	procGenBuffers                          *windows.Proc
	procGenerateMipmap                      *windows.Proc
	procGenFramebuffers                     *windows.Proc
	procGenRenderbuffers                    *windows.Proc
	procGenTextures                         *windows.Proc
	procGetActiveAttrib                     *windows.Proc
	procGetActiveUniform                    *windows.Proc
	procGetAttachedShaders                  *windows.Proc
	procGetAttribLocation                   *windows.Proc
	procGetBooleanv                         *windows.Proc
	procGetBufferParameteriv                *windows.Proc
	procGetError                            *windows.Proc
	procGetFloatv                           *windows.Proc
	procGetFramebufferAttachmentParameteriv *windows.Proc
	procGetIntegerv                         *windows.Proc
	procGetProgramiv                        *windows.Proc
	procGetProgramInfoLog                   *windows.Proc
	procGetRenderbufferParameteriv          *windows.Proc
	procGetShaderiv                         *windows.Proc
	procGetShaderInfoLog                    *windows.Proc
	procGetShaderPrecisionFormat            *windows.Proc
	procGetShaderSource                     *windows.Proc
	procGetString                           *windows.Proc
	procGetTexParameterfv                   *windows.Proc
	procGetTexParameteriv                   *windows.Proc
	procGetUniformfv                        *windows.Proc
	procGetUniformiv                        *windows.Proc
	procGetUniformLocation                  *windows.Proc
	procGetVertexAttribfv                   *windows.Proc
	procGetVertexAttribiv                   *windows.Proc
	procGetVertexAttribPointerv             *windows.Proc
	procHint                                *windows.Proc
	procIsBuffer                            *windows.Proc
	procIsEnabled                           *windows.Proc
	procIsFramebuffer                       *windows.Proc
	procIsProgram                           *windows.Proc
	procIsRenderbuffer                      *windows.Proc
	procIsShader                            *windows.Proc
	procIsTexture                           *windows.Proc
	procLineWidth                           *windows.Proc
	procLinkProgram                         *windows.Proc
	procPixelStorei                         *windows.Proc
	procPolygonOffset                       *windows.Proc
	procReadPixels                          *windows.Proc
	procReleaseShaderCompiler               *windows.Proc
	procRenderbufferStorage                 *windows.Proc
	procSampleCoverage                      *windows.Proc
	procScissor                             *windows.Proc
	procShaderBinary                        *windows.Proc
	procShaderSource                        *windows.Proc
	procStencilFunc                         *windows.Proc
	procStencilFuncSeparate                 *windows.Proc
	procStencilMask                         *windows.Proc
	procStencilMaskSeparate                 *windows.Proc
	procStencilOp                           *windows.Proc
	procStencilOpSeparate                   *windows.Proc
	procTexImage2D                          *windows.Proc
	procTexParameterf                       *windows.Proc
	procTexParameterfv                      *windows.Proc
	procTexParameteri                       *windows.Proc
	procTexParameteriv                      *windows.Proc
	procTexSubImage2D                       *windows.Proc
	procUniform1f                           *windows.Proc
	procUniform1fv                          *windows.Proc
	procUniform1i                           *windows.Proc
	procUniform1iv                          *windows.Proc
	procUniform2f                           *windows.Proc
	procUniform2fv                          *windows.Proc
	procUniform2i                           *windows.Proc
	procUniform2iv                          *windows.Proc
	procUniform3f                           *windows.Proc
	procUniform3fv                          *windows.Proc
	procUniform3i                           *windows.Proc
	procUniform3iv                          *windows.Proc
	procUniform4f                           *windows.Proc
	procUniform4fv                          *windows.Proc
	procUniform4i                           *windows.Proc
	procUniform4iv                          *windows.Proc
	procUniformMatrix2fv                    *windows.Proc
	procUniformMatrix3fv                    *windows.Proc
	procUniformMatrix4fv                    *windows.Proc
	procUseProgram                          *windows.Proc
	procValidateProgram                     *windows.Proc
	procVertexAttrib1f                      *windows.Proc
	procVertexAttrib1fv                     *windows.Proc
	procVertexAttrib2f                      *windows.Proc
	procVertexAttrib2fv                     *windows.Proc
	procVertexAttrib3f                      *windows.Proc
	procVertexAttrib3fv                     *windows.Proc
	procVertexAttrib4f                      *windows.Proc
	procVertexAttrib4fv                     *windows.Proc
	procVertexAttribPointer                 *windows.Proc
	procViewport                            *windows.Proc
)

func init() {
	dll = windows.MustLoadDLL("libGLESv2.dll")
	procActiveTexture = dll.MustFindProc("glActiveTexture")
	procAttachShader = dll.MustFindProc("glAttachShader")
	procBindAttribLocation = dll.MustFindProc("glBindAttribLocation")
	procBindBuffer = dll.MustFindProc("glBindBuffer")
	procBindFramebuffer = dll.MustFindProc("glBindFramebuffer")
	procBindRenderbuffer = dll.MustFindProc("glBindRenderbuffer")
	procBindTexture = dll.MustFindProc("glBindTexture")
	procBlendColor = dll.MustFindProc("glBlendColor")
	procBlendEquation = dll.MustFindProc("glBlendEquation")
	procBlendEquationSeparate = dll.MustFindProc("glBlendEquationSeparate")
	procBlendFunc = dll.MustFindProc("glBlendFunc")
	procBlendFuncSeparate = dll.MustFindProc("glBlendFuncSeparate")
	procBufferData = dll.MustFindProc("glBufferData")
	procBufferSubData = dll.MustFindProc("glBufferSubData")
	procCheckFramebufferStatus = dll.MustFindProc("glCheckFramebufferStatus")
	procClear = dll.MustFindProc("glClear")
	procClearColor = dll.MustFindProc("glClearColor")
	procClearDepthf = dll.MustFindProc("glClearDepthf")
	procClearStencil = dll.MustFindProc("glClearStencil")
	procColorMask = dll.MustFindProc("glColorMask")
	procCompileShader = dll.MustFindProc("glCompileShader")
	procCompressedTexImage2D = dll.MustFindProc("glCompressedTexImage2D")
	procCompressedTexSubImage2D = dll.MustFindProc("glCompressedTexSubImage2D")
	procCopyTexImage2D = dll.MustFindProc("glCopyTexImage2D")
	procCopyTexSubImage2D = dll.MustFindProc("glCopyTexSubImage2D")
	procCreateProgram = dll.MustFindProc("glCreateProgram")
	procCreateShader = dll.MustFindProc("glCreateShader")
	procCullFace = dll.MustFindProc("glCullFace")
	procDeleteBuffers = dll.MustFindProc("glDeleteBuffers")
	procDeleteFramebuffers = dll.MustFindProc("glDeleteFramebuffers")
	procDeleteProgram = dll.MustFindProc("glDeleteProgram")
	procDeleteRenderbuffers = dll.MustFindProc("glDeleteRenderbuffers")
	procDeleteShader = dll.MustFindProc("glDeleteShader")
	procDeleteTextures = dll.MustFindProc("glDeleteTextures")
	procDepthFunc = dll.MustFindProc("glDepthFunc")
	procDepthMask = dll.MustFindProc("glDepthMask")
	procDepthRangef = dll.MustFindProc("glDepthRangef")
	procDetachShader = dll.MustFindProc("glDetachShader")
	procDisable = dll.MustFindProc("glDisable")
	procDisableVertexAttribArray = dll.MustFindProc("glDisableVertexAttribArray")
	procDrawArrays = dll.MustFindProc("glDrawArrays")
	procDrawElements = dll.MustFindProc("glDrawElements")
	procEnable = dll.MustFindProc("glEnable")
	procEnableVertexAttribArray = dll.MustFindProc("glEnableVertexAttribArray")
	procFinish = dll.MustFindProc("glFinish")
	procFlush = dll.MustFindProc("glFlush")
	procFramebufferRenderbuffer = dll.MustFindProc("glFramebufferRenderbuffer")
	procFramebufferTexture2D = dll.MustFindProc("glFramebufferTexture2D")
	procFrontFace = dll.MustFindProc("glFrontFace")
	procGenBuffers = dll.MustFindProc("glGenBuffers")
	procGenerateMipmap = dll.MustFindProc("glGenerateMipmap")
	procGenFramebuffers = dll.MustFindProc("glGenFramebuffers")
	procGenRenderbuffers = dll.MustFindProc("glGenRenderbuffers")
	procGenTextures = dll.MustFindProc("glGenTextures")
	procGetActiveAttrib = dll.MustFindProc("glGetActiveAttrib")
	procGetActiveUniform = dll.MustFindProc("glGetActiveUniform")
	procGetAttachedShaders = dll.MustFindProc("glGetAttachedShaders")
	procGetAttribLocation = dll.MustFindProc("glGetAttribLocation")
	procGetBooleanv = dll.MustFindProc("glGetBooleanv")
	procGetBufferParameteriv = dll.MustFindProc("glGetBufferParameteriv")
	procGetError = dll.MustFindProc("glGetError")
	procGetFloatv = dll.MustFindProc("glGetFloatv")
	procGetFramebufferAttachmentParameteriv = dll.MustFindProc("glGetFramebufferAttachmentParameteriv")
	procGetIntegerv = dll.MustFindProc("glGetIntegerv")
	procGetProgramiv = dll.MustFindProc("glGetProgramiv")
	procGetProgramInfoLog = dll.MustFindProc("glGetProgramInfoLog")
	procGetRenderbufferParameteriv = dll.MustFindProc("glGetRenderbufferParameteriv")
	procGetShaderiv = dll.MustFindProc("glGetShaderiv")
	procGetShaderInfoLog = dll.MustFindProc("glGetShaderInfoLog")
	procGetShaderPrecisionFormat = dll.MustFindProc("glGetShaderPrecisionFormat")
	procGetShaderSource = dll.MustFindProc("glGetShaderSource")
	procGetString = dll.MustFindProc("glGetString")
	procGetTexParameterfv = dll.MustFindProc("glGetTexParameterfv")
	procGetTexParameteriv = dll.MustFindProc("glGetTexParameteriv")
	procGetUniformfv = dll.MustFindProc("glGetUniformfv")
	procGetUniformiv = dll.MustFindProc("glGetUniformiv")
	procGetUniformLocation = dll.MustFindProc("glGetUniformLocation")
	procGetVertexAttribfv = dll.MustFindProc("glGetVertexAttribfv")
	procGetVertexAttribiv = dll.MustFindProc("glGetVertexAttribiv")
	procGetVertexAttribPointerv = dll.MustFindProc("glGetVertexAttribPointerv")
	procHint = dll.MustFindProc("glHint")
	procIsBuffer = dll.MustFindProc("glIsBuffer")
	procIsEnabled = dll.MustFindProc("glIsEnabled")
	procIsFramebuffer = dll.MustFindProc("glIsFramebuffer")
	procIsProgram = dll.MustFindProc("glIsProgram")
	procIsRenderbuffer = dll.MustFindProc("glIsRenderbuffer")
	procIsShader = dll.MustFindProc("glIsShader")
	procIsTexture = dll.MustFindProc("glIsTexture")
	procLineWidth = dll.MustFindProc("glLineWidth")
	procLinkProgram = dll.MustFindProc("glLinkProgram")
	procPixelStorei = dll.MustFindProc("glPixelStorei")
	procPolygonOffset = dll.MustFindProc("glPolygonOffset")
	procReadPixels = dll.MustFindProc("glReadPixels")
	procReleaseShaderCompiler = dll.MustFindProc("glReleaseShaderCompiler")
	procRenderbufferStorage = dll.MustFindProc("glRenderbufferStorage")
	procSampleCoverage = dll.MustFindProc("glSampleCoverage")
	procScissor = dll.MustFindProc("glScissor")
	procShaderBinary = dll.MustFindProc("glShaderBinary")
	procShaderSource = dll.MustFindProc("glShaderSource")
	procStencilFunc = dll.MustFindProc("glStencilFunc")
	procStencilFuncSeparate = dll.MustFindProc("glStencilFuncSeparate")
	procStencilMask = dll.MustFindProc("glStencilMask")
	procStencilMaskSeparate = dll.MustFindProc("glStencilMaskSeparate")
	procStencilOp = dll.MustFindProc("glStencilOp")
	procStencilOpSeparate = dll.MustFindProc("glStencilOpSeparate")
	procTexImage2D = dll.MustFindProc("glTexImage2D")
	procTexParameterf = dll.MustFindProc("glTexParameterf")
	procTexParameterfv = dll.MustFindProc("glTexParameterfv")
	procTexParameteri = dll.MustFindProc("glTexParameteri")
	procTexParameteriv = dll.MustFindProc("glTexParameteriv")
	procTexSubImage2D = dll.MustFindProc("glTexSubImage2D")
	procUniform1f = dll.MustFindProc("glUniform1f")
	procUniform1fv = dll.MustFindProc("glUniform1fv")
	procUniform1i = dll.MustFindProc("glUniform1i")
	procUniform1iv = dll.MustFindProc("glUniform1iv")
	procUniform2f = dll.MustFindProc("glUniform2f")
	procUniform2fv = dll.MustFindProc("glUniform2fv")
	procUniform2i = dll.MustFindProc("glUniform2i")
	procUniform2iv = dll.MustFindProc("glUniform2iv")
	procUniform3f = dll.MustFindProc("glUniform3f")
	procUniform3fv = dll.MustFindProc("glUniform3fv")
	procUniform3i = dll.MustFindProc("glUniform3i")
	procUniform3iv = dll.MustFindProc("glUniform3iv")
	procUniform4f = dll.MustFindProc("glUniform4f")
	procUniform4fv = dll.MustFindProc("glUniform4fv")
	procUniform4i = dll.MustFindProc("glUniform4i")
	procUniform4iv = dll.MustFindProc("glUniform4iv")
	procUniformMatrix2fv = dll.MustFindProc("glUniformMatrix2fv")
	procUniformMatrix3fv = dll.MustFindProc("glUniformMatrix3fv")
	procUniformMatrix4fv = dll.MustFindProc("glUniformMatrix4fv")
	procUseProgram = dll.MustFindProc("glUseProgram")
	procValidateProgram = dll.MustFindProc("glValidateProgram")
	procVertexAttrib1f = dll.MustFindProc("glVertexAttrib1f")
	procVertexAttrib1fv = dll.MustFindProc("glVertexAttrib1fv")
	procVertexAttrib2f = dll.MustFindProc("glVertexAttrib2f")
	procVertexAttrib2fv = dll.MustFindProc("glVertexAttrib2fv")
	procVertexAttrib3f = dll.MustFindProc("glVertexAttrib3f")
	procVertexAttrib3fv = dll.MustFindProc("glVertexAttrib3fv")
	procVertexAttrib4f = dll.MustFindProc("glVertexAttrib4f")
	procVertexAttrib4fv = dll.MustFindProc("glVertexAttrib4fv")
	procVertexAttribPointer = dll.MustFindProc("glVertexAttribPointer")
	procViewport = dll.MustFindProc("glViewport")

}

func ActiveTexture(texture uint32) {
	procActiveTexture.Call(uintptr(texture))
}

func AttachShader(program uint32, shader uint32) {
	procAttachShader.Call(uintptr(program), uintptr(shader))
}

func BindAttribLocation(program uint32, index uint32, name string) {
	cstr := unsafe.Pointer(C.CString(name))
	defer C.free(cstr)
	procBindAttribLocation.Call(uintptr(program), uintptr(index), uintptr(cstr))
}

func BindBuffer(target int, b uint32) {
	procBindBuffer.Call(uintptr(target), uintptr(b))
}

func BindFramebuffer(target int, b uint32) {
	procBindFramebuffer.Call(uintptr(target), uintptr(b))
}

func BindRenderbuffer(target int, b uint32) {
	procBindRenderbuffer.Call(uintptr(target), uintptr(b))
}

func BindTexture(target int, texture uint32) {
	procBindTexture.Call(uintptr(target), uintptr(texture))
}

func BlendColor(
	red float32,
	green float32,
	blue float32,
	alpha float32) {
	procBlendColor.Call(
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha))
}

func BlendEquation(mode int) {
	procBlendEquation.Call(uintptr(mode))
}

func BlendEquationSeparate(
	modeRGB int,
	modeAlpha int) {
	procBlendEquationSeparate.Call(
		uintptr(modeRGB),
		uintptr(modeAlpha))
}

func BlendFunc(
	sfactor int,
	dfactor int) {
	procBlendFunc.Call(
		uintptr(sfactor),
		uintptr(dfactor))
}

func BlendFuncSeparate(
	sfactorRGB int,
	dfactorRGB int,
	sfactorAlpha int,
	dfactorAlpha int) {
	procBlendFuncSeparate.Call(
		uintptr(sfactorRGB),
		uintptr(dfactorRGB),
		uintptr(sfactorAlpha),
		uintptr(dfactorAlpha))
}

func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	procBufferData.Call(
		uintptr(target),
		uintptr(size),
		uintptr(data),
		uintptr(usage))
}

func BufferSubData(target int, offset int, size int, data unsafe.Pointer) {
	procBufferSubData.Call(
		uintptr(target),
		uintptr(offset),
		uintptr(size),
		uintptr(data))
}

func CheckFramebufferStatus(target int) int {
	r1, _, _ := procCheckFramebufferStatus.Call(uintptr(target))
	return int(r1)
}

func Clear(mask uint32) {
	procClear.Call(uintptr(mask))
}

func ClearColor(
	red float32,
	green float32,
	blue float32,
	alpha float32) {
	procClearColor.Call(
		f2ptr(red),
		f2ptr(green),
		f2ptr(blue),
		f2ptr(alpha))
}

func ClearDepthf(d float32) {
	procClearDepthf.Call(f2ptr(d))
}

func ClearStencil(s int) {
	procClearStencil.Call(uintptr(s))
}

func ColorMask(
	red int8,
	green int8,
	blue int8,
	alpha int8) {
	procColorMask.Call(
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha))
}

func CompileShader(shader uint32) {
	procCompileShader.Call(uintptr(shader))
}

func CompressedTexImage2D(
	target int,
	level int,
	internalformat int,
	width int,
	height int,
	border int,
	imageSize int,
	data unsafe.Pointer /*const void **/) {
	procCompressedTexImage2D.Call(
		uintptr(target),
		uintptr(level),
		uintptr(internalformat),
		uintptr(width),
		uintptr(height),
		uintptr(border),
		uintptr(imageSize),
		uintptr(data))
}

func CompressedTexSubImage2D(
	target int,
	level int,
	xoffset int,
	yoffset int,
	width int,
	height int,
	format int,
	imageSize int,
	data unsafe.Pointer /*const void **/) {
	procCompressedTexSubImage2D.Call(
		uintptr(target),
		uintptr(level),
		uintptr(xoffset),
		uintptr(yoffset),
		uintptr(width),
		uintptr(height),
		uintptr(format),
		uintptr(imageSize),
		uintptr(data))
}

func CopyTexImage2D(
	target int,
	level int,
	internalformat int,
	x int,
	y int,
	width int,
	height int,
	border int) {
	procCopyTexImage2D.Call(
		uintptr(target),
		uintptr(level),
		uintptr(internalformat),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(border))
}

func CopyTexSubImage2D(
	target int,
	level int,
	xoffset int,
	yoffset int,
	x int,
	y int,
	width int,
	height int) {
	procCopyTexSubImage2D.Call(
		uintptr(target),
		uintptr(level),
		uintptr(xoffset),
		uintptr(yoffset),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
}

func CreateProgram() uint32 {
	r1, _, _ := procCreateProgram.Call()
	return uint32(r1)
}

func CreateShader(ty uint32) uint32 {
	r1, _, _ := procCreateShader.Call(uintptr(ty))
	return uint32(r1)
}

func CullFace(mode int) {
	procCullFace.Call(uintptr(mode))
}

func DeleteBuffers(b []uint32) {
	procDeleteBuffers.Call(uintptr(len(b)), uintptr(unsafe.Pointer(&b[0])))
}

func DeleteBuffer(b uint32) {
	procDeleteBuffers.Call(uintptr(1), uintptr(unsafe.Pointer(&b)))
}

func DeleteFramebuffers(n int, framebuffers *uint32) {
	procDeleteFramebuffers.Call(uintptr(n), uintptr(unsafe.Pointer(framebuffers)))
}

func DeleteFramebuffer(b uint32) {
	DeleteFramebuffers(1, &b)
}

func DeleteProgram(program uint32) {
	procDeleteProgram.Call(uintptr(program))
}

func DeleteRenderbuffers(n int, renderbuffers *uint32) {
	procDeleteRenderbuffers.Call(uintptr(n), uintptr(unsafe.Pointer(renderbuffers)))
}

func DeleteRenderbuffer(b uint32) {
	DeleteRenderbuffers(1, &b)
}

func DeleteShader(shader uint32) {
	procDeleteShader.Call(uintptr(shader))
}

func DeleteTextures(n int, t unsafe.Pointer) { procDeleteTextures.Call(uintptr(n), uintptr(t)) }
func DeleteTexture(t uint32)                 { procDeleteTextures.Call(uintptr(1), uintptr(unsafe.Pointer(&t))) }

func DepthFunc(fn int)                 { procDepthFunc.Call(uintptr(fn)) }
func DepthMask(flag bool)              { procDepthMask.Call(uintptr(bool2uint8(flag))) }
func DepthRangef(n float32, f float32) { procDepthRangef.Call(uintptr(n), uintptr(f)) }

func DetachShader(
	program uint32,
	shader uint32) {
	procDetachShader.Call(
		uintptr(program),
		uintptr(shader))
}

func Disable(cap int) {
	procDisable.Call(uintptr(cap))
}

func DisableVertexAttribArray(index uint32) {
	procDisableVertexAttribArray.Call(uintptr(index))
}

func DrawArrays(
	mode int,
	first int,
	count int) {
	procDrawArrays.Call(
		uintptr(mode),
		uintptr(first),
		uintptr(count))
}

func DrawElements(mode uint32, count int, ty uint32, indices uintptr) {
	procDrawElements.Call(
		uintptr(mode),
		uintptr(count),
		uintptr(ty),
		uintptr(indices))
}

func Enable(cap int) {
	procEnable.Call(uintptr(cap))
}

func EnableVertexAttribArray(index uint32) {
	procEnableVertexAttribArray.Call(uintptr(index))
}

func Finish() {
	procFinish.Call()
}

func Flush() {
	procFlush.Call()
}

func FramebufferRenderbuffer(target, attachment, renderbuffertarget int, renderbuffer uint32) {
	procFramebufferRenderbuffer.Call(
		uintptr(target),
		uintptr(attachment),
		uintptr(renderbuffertarget),
		uintptr(renderbuffer))
}

func FramebufferTexture2D(
	target int,
	attachment int,
	textarget int,
	texture uint32,
	level int) {
	procFramebufferTexture2D.Call(
		uintptr(target),
		uintptr(attachment),
		uintptr(textarget),
		uintptr(texture),
		uintptr(level))
}

func FrontFace(mode int) {
	procFrontFace.Call(uintptr(mode))
}

func GenBuffers(b []uint32) {
	procGenBuffers.Call(uintptr(len(b)), uintptr(unsafe.Pointer(&b[0])))
}

func GenBuffer() (b uint32) {
	procGenBuffers.Call(uintptr(1), uintptr(unsafe.Pointer(&b)))
	return
}

func GenerateMipmap(target int) {
	procGenerateMipmap.Call(uintptr(target))
}

func GenFramebuffers(b []uint32) {
	procGenFramebuffers.Call(uintptr(len(b)), uintptr(unsafe.Pointer(&b[0])))
}

func GenFramebuffer() (b uint32) {
	procGenFramebuffers.Call(uintptr(1), uintptr(unsafe.Pointer(&b)))
	return
}

func GenRenderbuffers(b []uint32) {
	procGenRenderbuffers.Call(uintptr(len(b)), uintptr(unsafe.Pointer(&b[0])))
}

func GenRenderbuffer() (b uint32) {
	procGenRenderbuffers.Call(1, uintptr(unsafe.Pointer(&b)))
	return
}

func GenTextures(n int, textures *uint32) {
	procGenTextures.Call(uintptr(n), uintptr(unsafe.Pointer(textures)))
}

func GenTexture() uint32 {
	var t uint32
	procGenTextures.Call(uintptr(1), uintptr(unsafe.Pointer(&t)))
	return t
}

func GetActiveAttrib(program uint32, index uint32) (string, int32, int) {
	buf := make([]byte, 1024)
	var length int
	var size int32
	var ty int
	procGetActiveAttrib.Call(
		uintptr(program),
		uintptr(index),
		uintptr(1024),
		uintptr(unsafe.Pointer(&length)),
		uintptr(unsafe.Pointer(&size)),
		uintptr(unsafe.Pointer(&ty)),
		uintptr(unsafe.Pointer(&buf[0])))
	return goString(buf), int32(size), int(ty)
}

func GetActiveUniform(program uint32, i uint32) (string, int32, int) {
	buf := make([]byte, 1024)
	var length int
	var size int32
	var ty int
	procGetActiveUniform.Call(
		uintptr(program),
		uintptr(i),
		uintptr(1024),
		uintptr(unsafe.Pointer(&length)),
		uintptr(unsafe.Pointer(&size)),
		uintptr(unsafe.Pointer(&ty)),
		uintptr(unsafe.Pointer(&buf[0])))
	return goString(buf), int32(size), int(ty)
}

func GetAttachedShaders(
	program uint32,
	maxCount int,
	count unsafe.Pointer, /*GLsizei **/
	shaders unsafe.Pointer /*GLuint **/) {
	procGetAttachedShaders.Call(
		uintptr(program),
		uintptr(maxCount),
		uintptr(count),
		uintptr(shaders))
}

func GetAttribLocation(program uint32, name string) int32 {
	cstr := unsafe.Pointer(C.CString(name))
	defer C.free(cstr)
	r1, _, _ := procGetAttribLocation.Call(uintptr(program), uintptr(cstr))
	return int32(r1)
}

func GetBooleanv(
	pname int,
	data unsafe.Pointer /*GLboolean **/) {
	procGetBooleanv.Call(
		uintptr(pname),
		uintptr(data))
}

func GetBufferParameteriv(
	target int,
	pname int,
	params unsafe.Pointer /*GLint **/) {
	procGetBufferParameteriv.Call(
		uintptr(target),
		uintptr(pname),
		uintptr(params))
}

func GetError() error {
	r1, _, _ := procGetError.Call()
	switch int(r1) {
	case NO_ERROR:
		return nil
	case INVALID_ENUM:
		return ErrInvalidEnum
	case INVALID_VALUE:
		return ErrInvalidValue
	case INVALID_OPERATION:
		return ErrInvalidOperation
	}
	return ErrUnknown
}

func GetFloatv(
	pname int,
	data unsafe.Pointer /*GLfloat **/) {
	procGetFloatv.Call(
		uintptr(pname),
		uintptr(data))
}

func GetFramebufferAttachmentParameteriv(
	target int,
	attachment int,
	pname int,
	params unsafe.Pointer /*GLint **/) {
	procGetFramebufferAttachmentParameteriv.Call(
		uintptr(target),
		uintptr(attachment),
		uintptr(pname),
		uintptr(params))
}

func GetIntegerv(
	pname int,
	data unsafe.Pointer /*GLint **/) {
	procGetIntegerv.Call(
		uintptr(pname),
		uintptr(data))
}

func GetProgramiv(program uint32, pname int) (i int32) {
	procGetProgramiv.Call(uintptr(program), uintptr(pname), uintptr(unsafe.Pointer(&i)))
	return
}

func GetProgramInfoLog(program uint32) string {
	maxLength := GetProgramiv(program, INFO_LOG_LENGTH)
	if maxLength == 0 {
		return ""
	}
	buf := make([]byte, maxLength)
	var len int

	procGetProgramInfoLog.Call(
		uintptr(program),
		uintptr(maxLength),
		uintptr(unsafe.Pointer(&len)),
		uintptr(unsafe.Pointer(&buf[0])))
	return goString(buf)
}

func GetRenderbufferParameteriv(
	target int,
	pname int,
	params unsafe.Pointer /*GLint **/) {
	procGetRenderbufferParameteriv.Call(
		uintptr(target),
		uintptr(pname),
		uintptr(params))
}

func GetShaderiv(shader uint32, pname int) (params int32) {
	procGetShaderiv.Call(
		uintptr(shader),
		uintptr(pname),
		uintptr(unsafe.Pointer(&params)))
	return
}

func GetShaderInfoLog(shader uint32) string {
	maxLength := GetShaderiv(shader, INFO_LOG_LENGTH)

	buf := make([]byte, maxLength)
	var len int
	procGetShaderInfoLog.Call(
		uintptr(shader),
		uintptr(maxLength),
		uintptr(unsafe.Pointer(&len)),
		uintptr(unsafe.Pointer(&buf[0])))
	return goString(buf)
}

func GetShaderPrecisionFormat(
	shadertype int,
	precisiontype int,
	rng unsafe.Pointer, /*GLint **/
	precision unsafe.Pointer /*GLint **/) {
	procGetShaderPrecisionFormat.Call(
		uintptr(shadertype),
		uintptr(precisiontype),
		uintptr(rng),
		uintptr(precision))
}

func GetShaderSource(
	shader uint32,
	bufSize int,
	length unsafe.Pointer, /*GLsizei **/
	source unsafe.Pointer /*GLchar **/) {
	procGetShaderSource.Call(
		uintptr(shader),
		uintptr(bufSize),
		uintptr(length),
		uintptr(source))
}

func GetString(name int) unsafe.Pointer /*const GLubyte **/ {
	r1, _, _ := procGetString.Call(uintptr(name))
	return unsafe.Pointer /*const GLubyte **/ (r1)
}

func GetTexParameterfv(
	target int,
	pname int,
	params unsafe.Pointer /*GLfloat **/) {
	procGetTexParameterfv.Call(
		uintptr(target),
		uintptr(pname),
		uintptr(params))
}

func GetTexParameteriv(
	target int,
	pname int,
	params unsafe.Pointer /*GLint **/) {
	procGetTexParameteriv.Call(
		uintptr(target),
		uintptr(pname),
		uintptr(params))
}

func GetUniformfv(
	program uint32,
	location int,
	params unsafe.Pointer /*GLfloat **/) {
	procGetUniformfv.Call(
		uintptr(program),
		uintptr(location),
		uintptr(params))
}

func GetUniformiv(
	program uint32,
	location int,
	params unsafe.Pointer /*GLint **/) {
	procGetUniformiv.Call(
		uintptr(program),
		uintptr(location),
		uintptr(params))
}

func GetUniformLocation(program uint32, name string) int {
	cstr := unsafe.Pointer(C.CString(name))
	defer C.free(cstr)
	r1, _, _ := procGetUniformLocation.Call(uintptr(program), uintptr(cstr))
	return int(r1)
}

func GetVertexAttribfv(
	index uint32,
	pname int,
	params unsafe.Pointer /*GLfloat **/) {
	procGetVertexAttribfv.Call(
		uintptr(index),
		uintptr(pname),
		uintptr(params))
}

func GetVertexAttribiv(
	index uint32,
	pname int,
	params unsafe.Pointer /*GLint **/) {
	procGetVertexAttribiv.Call(
		uintptr(index),
		uintptr(pname),
		uintptr(params))
}

func GetVertexAttribPointerv(
	index uint32,
	pname int,
	pointer unsafe.Pointer /*void ***/) {
	procGetVertexAttribPointerv.Call(
		uintptr(index),
		uintptr(pname),
		uintptr(pointer))
}

func Hint(target int, mode int) {
	procHint.Call(uintptr(target), uintptr(mode))
}

func IsBuffer(buffer uint32) bool {
	r1, _, _ := procIsBuffer.Call(uintptr(buffer))
	return r1 == TRUE
}

func IsEnabled(cap int) bool {
	r1, _, _ := procIsEnabled.Call(uintptr(cap))
	return r1 == TRUE
}

func IsFramebuffer(framebuffer uint32) bool {
	r1, _, _ := procIsFramebuffer.Call(uintptr(framebuffer))
	return r1 == TRUE
}

func IsProgram(program uint32) bool {
	r1, _, _ := procIsProgram.Call(uintptr(program))
	return r1 == TRUE
}

func IsRenderbuffer(renderbuffer uint32) int8 {
	r1, _, _ := procIsRenderbuffer.Call(uintptr(renderbuffer))
	return int8(r1)
}

func IsShader(shader uint32) bool {
	r1, _, _ := procIsShader.Call(uintptr(shader))
	return r1 != FALSE
}

func IsTexture(texture uint32) bool {
	r1, _, _ := procIsTexture.Call(uintptr(texture))
	return r1 == TRUE
}

func LineWidth(width float32) {
	procLineWidth.Call(uintptr(width))
}

func LinkProgram(program uint32) {
	procLinkProgram.Call(uintptr(program))
}

func PixelStorei(
	pname int,
	param int) {
	procPixelStorei.Call(
		uintptr(pname),
		uintptr(param))
}

func PolygonOffset(
	factor float32,
	units float32) {
	procPolygonOffset.Call(
		uintptr(factor),
		uintptr(units))
}

func ReadPixels(x, y, width, height, format, ty int, data []byte) {
	procReadPixels.Call(
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(format),
		uintptr(ty),
		uintptr(unsafe.Pointer(&data[0])))
}

func ReleaseShaderCompiler() {
	procReleaseShaderCompiler.Call()
}

func RenderbufferStorage(
	target int,
	internalformat int,
	width int,
	height int) {
	procRenderbufferStorage.Call(
		uintptr(target),
		uintptr(internalformat),
		uintptr(width),
		uintptr(height))
}

func SampleCoverage(
	value float32,
	invert int8) {
	procSampleCoverage.Call(
		uintptr(value),
		uintptr(invert))
}

func Scissor(
	x int,
	y int,
	width int,
	height int) {
	procScissor.Call(
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
}

func ShaderBinary(
	count int,
	shaders unsafe.Pointer, /*const GLuint **/
	binaryformat int,
	binary unsafe.Pointer, /*const void **/
	length int) {
	procShaderBinary.Call(
		uintptr(count),
		uintptr(shaders),
		uintptr(binaryformat),
		uintptr(binary),
		uintptr(length))
}

func ShaderSource(shader uint32, src string) {
	cstr := unsafe.Pointer(C.CString(src))
	ptr := C.malloc(C.size_t(unsafe.Sizeof((*int)(nil))))
	*(*uintptr)(ptr) = uintptr(cstr)
	defer C.free(ptr)
	defer C.free(cstr)
	procShaderSource.Call(uintptr(shader), uintptr(1), uintptr(ptr), uintptr(0))
}

func StencilFunc(fn int, ref int, mask uint32) {
	procStencilFunc.Call(uintptr(fn), uintptr(ref), uintptr(mask))
}

func StencilFuncSeparate(
	face int,
	fn int,
	ref int,
	mask uint32) {
	procStencilFuncSeparate.Call(
		uintptr(face),
		uintptr(fn),
		uintptr(ref),
		uintptr(mask))
}

func StencilMask(mask uint32) {
	procStencilMask.Call(uintptr(mask))
}

func StencilMaskSeparate(
	face int,
	mask uint32) {
	procStencilMaskSeparate.Call(
		uintptr(face),
		uintptr(mask))
}

func StencilOp(
	fail int,
	zfail int,
	zpass int) {
	procStencilOp.Call(
		uintptr(fail),
		uintptr(zfail),
		uintptr(zpass))
}

func StencilOpSeparate(
	face int,
	sfail int,
	dpfail int,
	dppass int) {
	procStencilOpSeparate.Call(
		uintptr(face),
		uintptr(sfail),
		uintptr(dpfail),
		uintptr(dppass))
}

func TexImage2D(
	target uint32,
	level int,
	internalFormat uint32,
	width int,
	height int,
	border int,
	format uint32,
	ty uint32,
	data uintptr) {
	procTexImage2D.Call(
		uintptr(target),
		uintptr(level),
		uintptr(internalFormat),
		uintptr(width),
		uintptr(height),
		uintptr(border),
		uintptr(format),
		uintptr(ty),
		data)

}

func TexParameterf(
	target int,
	pname int,
	param float32) {
	procTexParameterf.Call(
		uintptr(target),
		uintptr(pname),
		uintptr(param))
}

func TexParameterfv(
	target int,
	pname int,
	params unsafe.Pointer /*const GLfloat **/) {
	procTexParameterfv.Call(
		uintptr(target),
		uintptr(pname),
		uintptr(params))
}

func TexParameteri(target int, pname int, param int) {
	procTexParameteri.Call(uintptr(target), uintptr(pname), uintptr(param))
}

func TexParameteriv(
	target int,
	pname int,
	params unsafe.Pointer /*const GLint **/) {
	procTexParameteriv.Call(
		uintptr(target),
		uintptr(pname),
		uintptr(params))
}

func TexSubImage2D(
	target int,
	level int,
	xoffset int,
	yoffset int,
	width int,
	height int,
	format int,
	ty int,
	pixels unsafe.Pointer /*const void **/) {
	procTexSubImage2D.Call(
		uintptr(target),
		uintptr(level),
		uintptr(xoffset),
		uintptr(yoffset),
		uintptr(width),
		uintptr(height),
		uintptr(format),
		uintptr(ty),
		uintptr(pixels))
}

func Uniform1f(location int, v0 float32) {
	procUniform1f.Call(uintptr(location), f2ptr(v0))
}

func Uniform1fv(
	location int,
	count int,
	value unsafe.Pointer /*const GLfloat **/) {
	procUniform1fv.Call(
		uintptr(location),
		uintptr(count),
		uintptr(value))
}

func Uniform1i(
	location int,
	v0 int) {
	procUniform1i.Call(
		uintptr(location),
		uintptr(v0))
}

func Uniform1iv(
	location int,
	count int,
	value unsafe.Pointer /*const GLint **/) {
	procUniform1iv.Call(
		uintptr(location),
		uintptr(count),
		uintptr(value))
}

func Uniform2f(
	location int,
	v0 float32,
	v1 float32) {
	procUniform2f.Call(
		uintptr(location),
		uintptr(v0),
		uintptr(v1))
}

func Uniform2fv(
	location int,
	count int,
	value unsafe.Pointer /*const GLfloat **/) {
	procUniform2fv.Call(
		uintptr(location),
		uintptr(count),
		uintptr(value))
}

func Uniform2i(
	location int,
	v0 int,
	v1 int) {
	procUniform2i.Call(
		uintptr(location),
		uintptr(v0),
		uintptr(v1))
}

func Uniform2iv(
	location int,
	count int,
	value unsafe.Pointer /*const GLint **/) {
	procUniform2iv.Call(
		uintptr(location),
		uintptr(count),
		uintptr(value))
}

func Uniform3f(
	location int, v0, v1, v2 float32) {
	procUniform3f.Call(uintptr(location), f2ptr(v0), f2ptr(v1), f2ptr(v2))
}

func Uniform3fv(
	location int,
	count int,
	value unsafe.Pointer /*const GLfloat **/) {
	procUniform3fv.Call(
		uintptr(location),
		uintptr(count),
		uintptr(value))
}

func Uniform3i(
	location int,
	v0 int,
	v1 int,
	v2 int) {
	procUniform3i.Call(
		uintptr(location),
		uintptr(v0),
		uintptr(v1),
		uintptr(v2))
}

func Uniform3iv(
	location int,
	count int,
	value unsafe.Pointer /*const GLint **/) {
	procUniform3iv.Call(
		uintptr(location),
		uintptr(count),
		uintptr(value))
}

func Uniform4f(
	location int,
	v0 float32,
	v1 float32,
	v2 float32,
	v3 float32) {
	procUniform4f.Call(
		uintptr(location),
		uintptr(v0),
		uintptr(v1),
		uintptr(v2),
		uintptr(v3))
}

func Uniform4fv(
	location int,
	count int,
	value unsafe.Pointer /*const GLfloat **/) {
	procUniform4fv.Call(
		uintptr(location),
		uintptr(count),
		uintptr(value))
}

func Uniform4i(
	location int,
	v0 int,
	v1 int,
	v2 int,
	v3 int) {
	procUniform4i.Call(
		uintptr(location),
		uintptr(v0),
		uintptr(v1),
		uintptr(v2),
		uintptr(v3))
}

func Uniform4iv(
	location int,
	count int,
	value unsafe.Pointer /*const GLint **/) {
	procUniform4iv.Call(
		uintptr(location),
		uintptr(count),
		uintptr(value))
}

func UniformMatrix2fv(
	location int,
	count int,
	transpose int8,
	value unsafe.Pointer /*const GLfloat **/) {
	procUniformMatrix2fv.Call(
		uintptr(location),
		uintptr(count),
		uintptr(transpose),
		uintptr(value))
}

func UniformMatrix3fv(location int, value []float32) {
	procUniformMatrix3fv.Call(
		uintptr(location),
		uintptr(1),
		uintptr(0),
		uintptr(unsafe.Pointer(&value[0])))
}

func UniformMatrix4fv(location int, count int, transpose bool, value uintptr) {
	procUniformMatrix4fv.Call(
		uintptr(location),
		uintptr(count),
		cGLboolean(transpose),
		value)
}

func UseProgram(p uint32) {
	procUseProgram.Call(uintptr(p))
}

func ValidateProgram(p uint32) {
	procValidateProgram.Call(uintptr(p))
}

func VertexAttrib1f(
	index uint32,
	x float32) {
	procVertexAttrib1f.Call(
		uintptr(index),
		uintptr(x))
}

func VertexAttrib1fv(
	index uint32,
	v unsafe.Pointer /*const GLfloat **/) {
	procVertexAttrib1fv.Call(
		uintptr(index),
		uintptr(v))
}

func VertexAttrib2f(
	index uint32,
	x float32,
	y float32) {
	procVertexAttrib2f.Call(
		uintptr(index),
		uintptr(x),
		uintptr(y))
}

func VertexAttrib2fv(
	index uint32,
	v unsafe.Pointer /*const GLfloat **/) {
	procVertexAttrib2fv.Call(
		uintptr(index),
		uintptr(v))
}

func VertexAttrib3f(
	index uint32,
	x float32,
	y float32,
	z float32) {
	procVertexAttrib3f.Call(
		uintptr(index),
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func VertexAttrib3fv(
	index uint32,
	v unsafe.Pointer /*const GLfloat **/) {
	procVertexAttrib3fv.Call(
		uintptr(index),
		uintptr(v))
}

func VertexAttrib4f(
	index uint32,
	x float32,
	y float32,
	z float32,
	w float32) {
	procVertexAttrib4f.Call(
		uintptr(index),
		uintptr(x),
		uintptr(y),
		uintptr(z),
		uintptr(w))
}

func VertexAttrib4fv(
	index uint32,
	v unsafe.Pointer /*const GLfloat **/) {
	procVertexAttrib4fv.Call(
		uintptr(index),
		uintptr(v))
}

func VertexAttribPointer(
	index uint32,
	size int,
	ty uint32,
	normalized bool,
	stride int,
	pointer uintptr) {

	n := FALSE
	if normalized {
		n = TRUE
	}
	procVertexAttribPointer.Call(
		uintptr(index),
		uintptr(size),
		uintptr(ty),
		uintptr(n),
		uintptr(stride),
		uintptr(pointer))
}

func Viewport(x, y, w, h int32) {
	procViewport.Call(uintptr(x), uintptr(y), uintptr(w), uintptr(h))
}

func bool2uint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func f2ptr(f float32) uintptr {
	return uintptr(math.Float32bits(f))
}

func goString(buf []byte) string {
	for i, b := range buf {
		if b == 0 {
			return string(buf[:i])
		}
	}
	panic("buf is not NUL-terminated")
}

func cGLboolean(b bool) uintptr {
	if b {
		return uintptr(TRUE)
	}
	return uintptr(FALSE)
}
