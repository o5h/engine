//go:build android

package egl

/*
#cgo LDFLAGS: -lEGL
#include <EGL/egl.h>
*/
import "C"
import (
	"unsafe"
)

type NativeWindow uintptr
type NativeDisplay uintptr

func GetDisplay(ty NativeDisplay) Display {
	return Display(C.eglGetDisplay((C.EGLNativeDisplayType)(ty)))
}

func Initialize(display Display) (int, int, error) {
	var major, minor C.int
	res := C.eglInitialize(C.EGLDisplay(display), &major, &minor)
	if res == FALSE {
		return 0, 0, GetError()
	}
	return int(major), int(minor), nil
}

func ChooseConfig(display Display, attrsList ...int32) Config {
	var config C.EGLConfig
	var numConfigs C.int
	res := C.eglChooseConfig(
		C.EGLDisplay(display),
		cI32SToEGLintPtr(attrsList),
		&config,
		1,
		&numConfigs,
	)

	if res == FALSE {
		return NO_CONFIG
	}
	return Config(config)
}

func MakeCurrent(display Display, draw Surface, read Surface, context Context) bool {
	return C.eglMakeCurrent(
		C.EGLDisplay(display),
		C.EGLSurface(draw),
		C.EGLSurface(read),
		C.EGLContext(context)) == TRUE
}

func GetConfigAttrib(display Display, config Config, attribute int) (int, error) {
	var value C.EGLint
	res := C.eglGetConfigAttrib(C.EGLDisplay(display),
		C.EGLConfig(config),
		C.EGLint(attribute),
		&value)
	if res == FALSE {
		return 0, GetError()
	}

	return int(value), nil
}

func CreateWindowSurface(display Display, config Config, window NativeWindow, attrsList ...int32) Surface {
	return Surface(C.eglCreateWindowSurface(
		C.EGLDisplay(display),
		C.EGLConfig(config),
		C.EGLNativeWindowType(unsafe.Pointer(window)),
		cI32SToEGLintPtr(attrsList)))
}

func SwapBuffers(display Display, surface Surface) bool {
	return C.eglSwapBuffers(C.EGLDisplay(display), C.EGLSurface(surface)) == C.EGL_TRUE
}

func CreateContext(display Display, config Config, context Context, attrsList ...int32) Context {
	return Context(C.eglCreateContext(
		C.EGLDisplay(display),
		C.EGLConfig(config),
		C.EGLContext(context),
		cI32SToEGLintPtr(attrsList)))
}

func DestroySurface(display Display, surface Surface) bool {
	return C.eglDestroySurface(C.EGLDisplay(display), C.EGLSurface(surface)) == TRUE
}

func GetError() error {
	switch errNum := C.eglGetError(); errNum {
	case C.EGL_SUCCESS:
		return nil
	case C.EGL_NOT_INITIALIZED:
		return ErrNotInitialized
	case C.EGL_BAD_ACCESS:
		return ErrBadAccess
	case C.EGL_BAD_ALLOC:
		return ErrBadAlloc
	case C.EGL_BAD_ATTRIBUTE:
		return ErrBadAttribute
	case C.EGL_BAD_CONTEXT:
		return ErrBadContext
	case C.EGL_BAD_CONFIG:
		return ErrBadConfig
	case C.EGL_BAD_CURRENT_SURFACE:
		return ErrBadCurrentSurface
	case C.EGL_BAD_DISPLAY:
		return ErrBadDisplay
	case C.EGL_BAD_SURFACE:
		return ErrBadSurface
	case C.EGL_BAD_MATCH:
		return ErrBadMatch
	case C.EGL_BAD_PARAMETER:
		return ErrBadParameter
	case C.EGL_BAD_NATIVE_PIXMAP:
		return ErrBadNativePixmap
	case C.EGL_BAD_NATIVE_WINDOW:
		return ErrBadNativeWindow
	case C.EGL_CONTEXT_LOST:
		return ErrContextLost
	default:
		return ErrUnknown
	}
}

func cI32SToEGLintPtr(s []int32) *C.EGLint {
	if len(s) == 0 {
		return (*C.int)(unsafe.Pointer(uintptr(0)))
	} else {
		return (*C.EGLint)(unsafe.Pointer(&s[0]))
	}
}
