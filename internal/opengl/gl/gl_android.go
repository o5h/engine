//go:build android

package gl

/*
#cgo LDFLAGS: -lGLESv2

#include <stdlib.h>
#include <GLES2/gl2.h>
*/
import "C"

import (
	"unsafe"
)

func ClearColor(r, g, b, a float32) {
	C.glClearColor(C.GLfloat(r), C.GLfloat(g), C.GLfloat(b), C.GLfloat(a))
}

// void glClear(	GLbitfield mask);
func Clear(mask uint32) {
	C.glClear(C.uint(mask))
}

func GenTexture() Texture {
	textures := []Texture{0}
	GenTextures(textures)
	return textures[0]
}

// void glGenTextures(	GLsizei n,	GLuint * textures);
func GenTextures(textures []Texture) {
	if len(textures) == 0 {
		return
	}
	C.glGenTextures(C.GLsizei(len(textures)), (*C.GLuint)(unsafe.Pointer(&textures[0])))
}

// void glBindTexture(	GLenum target, GLuint texture);
func BindTexture(target Enum, texture Texture) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

// void glBindBuffer(	GLenum target, GLuint buffer);
func BindBuffer(target Enum, buffer Buffer) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}

func GenBuffer() Buffer {
	buffers := []Buffer{0}
	GenBuffers(buffers)
	return buffers[0]
}

// void glGenBuffers(	GLsizei n,	GLuint * buffers)
func GenBuffers(buffers []Buffer) {
	C.glGenBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(unsafe.Pointer(&buffers[0])))
}

// void glTexImage2D(	GLenum target,
// 	GLint level,
// 	GLint internalFormat,
// 	GLsizei width,
// 	GLsizei height,
// 	GLint border,
// 	GLenum format,
// 	GLenum type,
// 	const void * data);

func TexImage2D(target Enum, level int, internalFormat TextureFormat,
	width int, height int, border int,
	format TextureFormat,
	ty TextureDataFormat,
	data uintptr) {
	C.glTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(internalFormat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLint(border),
		C.GLenum(format),
		C.GLenum(ty),
		unsafe.Pointer(data))
}

// void glTexParameteri(	GLenum target,
// 	GLenum pname,
// 	GLint param);

func TexParameteri(target Enum, name Enum, value int) {
	C.glTexParameteri(
		C.GLenum(target),
		C.GLenum(name),
		C.GLint(value))
}

// void glGenerateMipmap(	GLenum target);
func GenerateMipmap(target Enum) {
	C.glGenerateMipmap(C.GLenum(target))
}

// void glActiveTexture(	GLenum texture);
func ActiveTexture(texture Enum) {
	C.glActiveTexture(C.GLenum(texture))
}

// void glDrawElements(	GLenum mode,	GLsizei count,	GLenum type,	const void * indices);
func DrawElements(mode DrawMode, count int, ty IndexFormat, indices uintptr) {
	C.glDrawElements(
		C.GLenum(mode),
		C.GLsizei(count),
		C.GLenum(ty),
		unsafe.Pointer(indices))
}

// void glBufferData(GLenum target,	GLsizeiptr size,	const void * data,	GLenum usage);
func BufferData(target Enum, size int, data uintptr, usage BufferUsage) {
	C.glBufferData(
		C.GLenum(target),
		C.GLsizeiptr(size),
		unsafe.Pointer(data),
		C.GLenum(usage))
}

// GLuint glCreateShader(	GLenum shaderType);
func CreateShader(ty ShaderType) Shader {
	return Shader(C.glCreateShader(C.GLenum(ty)))
}

func IsShader(shader Shader) bool {
	return C.glIsShader(C.GLuint(shader)) == TRUE
}

// void glShaderSource(	GLuint shader,
//
//	GLsizei count,
//	const GLchar **string,
//	const GLint *length);
func ShaderSource(shader Shader, src string) {
	cstr := unsafe.Pointer(C.CString(src))
	ptr := C.malloc(C.size_t(unsafe.Sizeof((*int)(nil)))) //???
	*(*uintptr)(ptr) = uintptr(cstr)
	length := C.GLint(len(src))
	defer C.free(ptr)
	defer C.free(cstr)
	C.glShaderSource(C.GLuint(shader), 1, (**C.GLchar)(unsafe.Pointer(ptr)), &length)
}

// void glCompileShader(	GLuint shader);
func CompileShader(shader Shader) {
	C.glCompileShader(C.GLuint(shader))
}

// void glGetShaderiv(	GLuint shader, GLenum pname, GLint *params);
func GetShaderiv(shader Shader, pname Enum) int {
	var params C.GLint
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(pname), &params)
	return int(params)
}

// GLenum glGetError(	void);
func GetError() error {
	err := C.glGetError()

	switch err {
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

// void glGetShaderInfoLog(	GLuint shader,	GLsizei maxLength,	GLsizei *length,	GLchar *infoLog);
func GetShaderInfoLog(shader Shader) string {
	maxLength := GetShaderiv(shader, INFO_LOG_LENGTH)
	if maxLength == 0 {
		return ""
	}
	buf := make([]byte, maxLength)
	var len C.GLsizei
	C.glGetShaderInfoLog(
		C.GLuint(shader),
		C.GLsizei(maxLength),
		&len,
		(*C.GLchar)(unsafe.Pointer(&buf[0])))
	return goString(buf)
}

func goString(buf []byte) string {
	for i, b := range buf {
		if b == 0 {
			return string(buf[:i])
		}
	}
	return string(buf)
}

// GLuint glCreateProgram(	void);
func CreateProgram() Program {
	return Program(C.glCreateProgram())
}

// void glDeleteProgram(	GLuint program);
func DeleteProgram(program Program) {
	C.glDeleteProgram(C.GLuint(program))
}

// void glDeleteShader(	GLuint shader);
func DeleteShader(shader Shader) {
	C.glDeleteShader(C.GLuint(shader))
}

// void glBindAttribLocation(	GLuint program,	GLuint index,	const GLchar *name);
func BindAttribLocation(program Program, index uint32, name string) {
	cstr := unsafe.Pointer(C.CString(name))
	C.glBindAttribLocation(C.GLuint(program), C.GLuint(index), (*C.GLchar)(cstr))
	defer C.free(cstr)
}

// void glAttachShader(	GLuint program, GLuint shader);
func AttachShader(program Program, shader Shader) {
	C.glAttachShader(C.GLuint(program), C.GLuint(shader))
}

// void glLinkProgram(	GLuint program);
func LinkProgram(program Program) {
	C.glLinkProgram(C.GLuint(program))
}

// void glGetProgramiv(	GLuint program,GLenum pname,GLint *params);
func GetProgramiv(program Program, pname Enum) int {
	var params C.GLint
	C.glGetProgramiv(C.GLuint(program), C.GLenum(pname), &params)
	return int(params)

}

// void glGetProgramInfoLog(	GLuint program,
// 	GLsizei maxLength,
// 	GLsizei *length,
// 	GLchar *infoLog);

func GetProgramInfoLog(program Program) string {

	maxLength := GetProgramiv(program, INFO_LOG_LENGTH)
	if maxLength == 0 {
		return ""
	}
	buf := make([]byte, maxLength)

	var len C.GLsizei
	C.glGetProgramInfoLog(
		C.GLuint(program),
		C.GLsizei(maxLength),
		&len,
		(*C.GLchar)(unsafe.Pointer(&buf[0])))
	return goString(buf)

}

// void glEnableVertexAttribArray(	GLuint index);
func EnableVertexAttribArray(index uint32) {
	C.glEnableVertexAttribArray(C.GLuint(index))
}

// void glDisableVertexAttribArray(	GLuint index);
func DisableVertexAttribArray(index uint32) {
	C.glDisableVertexAttribArray(C.GLuint(index))
}

// void glGetActiveAttrib(	GLuint program,
// 	GLuint index,
// 	GLsizei bufSize,
// 	GLsizei *length,
// 	GLint *size,
// 	GLenum *type,
// 	GLchar *name);

func GetActiveAttrib(program Program, index uint32) (string, int32, int) {
	bufSize := C.GLsizei(1024)
	buf := make([]byte, 1024)
	var length C.GLsizei
	var size C.GLint
	var ty C.GLenum
	C.glGetActiveAttrib(
		C.GLuint(program),
		C.GLuint(index),
		bufSize,
		&length,
		&size,
		&ty,
		(*C.GLchar)(unsafe.Pointer(&buf[0])))
	return goString(buf), int32(size), int(ty)
}

//GLint glGetAttribLocation(	GLuint program, const GLchar *name);

func GetAttribLocation(program Program, name string) int {
	cstr := unsafe.Pointer(C.CString(name))
	defer C.free(cstr)
	return int(C.glGetAttribLocation(C.GLuint(program), (*C.GLchar)(cstr)))
}

// void glGetActiveUniform(	GLuint program,
// 	GLuint index,
// 	GLsizei bufSize,
// 	GLsizei *length,
// 	GLint *size,
// 	GLenum *type,
// 	GLchar *name);

func GetActiveUniform(program Program, index uint32) (string, int32, int) {
	bufSize := C.GLsizei(1024)
	buf := make([]byte, 1024)
	var length C.GLsizei
	var size C.GLint
	var ty C.GLenum
	C.glGetActiveUniform(
		C.GLuint(program),
		C.GLuint(index),
		bufSize,
		&length,
		&size,
		&ty,
		(*C.GLchar)(unsafe.Pointer(&buf[0])))
	return goString(buf), int32(size), int(ty)
}

// GLint glGetUniformLocation(	GLuint program, const GLchar *name);
func GetUniformLocation(program Program, name string) int {
	cstr := unsafe.Pointer(C.CString(name))
	defer C.free(cstr)
	return int(C.glGetUniformLocation(C.GLuint(program), (*C.GLchar)(cstr)))
}

func cGLboolean(b bool) C.GLboolean {
	if b {
		return C.GLboolean(TRUE)
	}
	return C.GLboolean(FALSE)
}

// void glVertexAttribPointer(	GLuint index,
// 	GLint size,
// 	GLenum type,
// 	GLboolean normalized,
// 	GLsizei stride,
// 	const void * pointer);

func VertexAttribPointer(
	index uint32,
	size int,
	ty int,
	normalized bool,
	stride int,
	pointer uintptr) {

	C.glVertexAttribPointer(
		C.GLuint(index),
		C.GLint(size),
		C.GLenum(ty),
		cGLboolean(normalized),
		C.GLsizei(stride),
		unsafe.Pointer(pointer))
}

// void glUniformMatrix4fv(	GLint location,
//
//	GLsizei count,
//	GLboolean transpose,
//	const GLfloat *value);
func UniformMatrix4fv(location int, count int, transpose bool, value uintptr) {
	C.glUniformMatrix4fv(
		C.GLint(location),
		C.GLsizei(count),
		cGLboolean(transpose),
		(*C.GLfloat)(unsafe.Pointer(value)))
}

// void glUniform3f(	GLint location,
// 	GLfloat v0,
// 	GLfloat v1,
// 	GLfloat v2);

func Uniform3f(location int, v0, v1, v2 float32) {
	C.glUniform3f(
		C.GLint(location),
		C.GLfloat(v0),
		C.GLfloat(v1),
		C.GLfloat(v2))
}

// void glUseProgram(	GLuint program);
func UseProgram(program Program) {
	C.glUseProgram(C.GLuint(program))
}

// void glViewport(	GLint x,
// 	GLint y,
// 	GLsizei width,
// 	GLsizei height);

func Viewport(x, y, width, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
