//go:build windows

package egl

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type NativeWindow uintptr
type NativeDisplay uintptr

var (
	dll *windows.DLL

	//angle
	procGetPlatformDisplayEXT *windows.Proc

	procBindAPI                       *windows.Proc
	procBindTexImage                  *windows.Proc
	procChooseConfig                  *windows.Proc
	procClientWaitSync                *windows.Proc
	procCopyBuffers                   *windows.Proc
	procCreateContext                 *windows.Proc
	procCreateImage                   *windows.Proc
	procCreatePbufferFromClientBuffer *windows.Proc
	procCreatePbufferSurface          *windows.Proc
	procCreatePixmapSurface           *windows.Proc
	procCreatePlatformPixmapSurface   *windows.Proc
	procCreatePlatformWindowSurface   *windows.Proc
	procCreateSync                    *windows.Proc
	procCreateWindowSurface           *windows.Proc
	procDestroyContext                *windows.Proc
	procDestroyImage                  *windows.Proc
	procDestroySurface                *windows.Proc
	procDestroySync                   *windows.Proc
	procGetConfigAttrib               *windows.Proc
	procGetConfigs                    *windows.Proc
	procGetCurrentContext             *windows.Proc
	procGetCurrentDisplay             *windows.Proc
	procGetCurrentSurface             *windows.Proc
	procGetDisplay                    *windows.Proc
	procGetError                      *windows.Proc
	procGetPlatformDisplay            *windows.Proc
	procGetProcAddress                *windows.Proc
	procGetSyncAttrib                 *windows.Proc
	procInitialize                    *windows.Proc
	procMakeCurrent                   *windows.Proc
	procQueryAPI                      *windows.Proc
	procQueryContext                  *windows.Proc
	procQueryString                   *windows.Proc
	procQuerySurface                  *windows.Proc
	procReleaseTexImage               *windows.Proc
	procReleaseThread                 *windows.Proc
	procSurfaceAttrib                 *windows.Proc
	procSwapBuffers                   *windows.Proc
	procSwapInterval                  *windows.Proc
	procTerminate                     *windows.Proc
	procWaitClient                    *windows.Proc
	procWaitGL                        *windows.Proc
	procWaitNative                    *windows.Proc
	procWaitSync                      *windows.Proc
)

func init() {
	dll = windows.MustLoadDLL("libEGL.dll")

	//angle
	procGetPlatformDisplayEXT = dll.MustFindProc("eglGetPlatformDisplayEXT")

	procBindAPI = dll.MustFindProc("eglBindAPI")
	procBindTexImage = dll.MustFindProc("eglBindTexImage")
	procChooseConfig = dll.MustFindProc("eglChooseConfig")
	procClientWaitSync = dll.MustFindProc("eglClientWaitSync")
	procCopyBuffers = dll.MustFindProc("eglCopyBuffers")
	procCreateContext = dll.MustFindProc("eglCreateContext")
	procCreateImage = dll.MustFindProc("eglCreateImage")
	procCreatePbufferFromClientBuffer = dll.MustFindProc("eglCreatePbufferFromClientBuffer")
	procCreatePbufferSurface = dll.MustFindProc("eglCreatePbufferSurface")
	procCreatePixmapSurface = dll.MustFindProc("eglCreatePixmapSurface")
	procCreatePlatformPixmapSurface = dll.MustFindProc("eglCreatePlatformPixmapSurface")
	procCreatePlatformWindowSurface = dll.MustFindProc("eglCreatePlatformWindowSurface")
	procCreateSync = dll.MustFindProc("eglCreateSync")
	procCreateWindowSurface = dll.MustFindProc("eglCreateWindowSurface")
	procDestroyContext = dll.MustFindProc("eglDestroyContext")
	procDestroyImage = dll.MustFindProc("eglDestroyImage")
	procDestroySurface = dll.MustFindProc("eglDestroySurface")
	procDestroySync = dll.MustFindProc("eglDestroySync")
	procGetConfigAttrib = dll.MustFindProc("eglGetConfigAttrib")
	procGetConfigs = dll.MustFindProc("eglGetConfigs")
	procGetCurrentContext = dll.MustFindProc("eglGetCurrentContext")
	procGetCurrentDisplay = dll.MustFindProc("eglGetCurrentDisplay")
	procGetCurrentSurface = dll.MustFindProc("eglGetCurrentSurface")
	procGetDisplay = dll.MustFindProc("eglGetDisplay")
	procGetError = dll.MustFindProc("eglGetError")
	procGetPlatformDisplay = dll.MustFindProc("eglGetPlatformDisplay")
	procGetProcAddress = dll.MustFindProc("eglGetProcAddress")
	procGetSyncAttrib = dll.MustFindProc("eglGetSyncAttrib")
	procInitialize = dll.MustFindProc("eglInitialize")
	procMakeCurrent = dll.MustFindProc("eglMakeCurrent")
	procQueryAPI = dll.MustFindProc("eglQueryAPI")
	procQueryContext = dll.MustFindProc("eglQueryContext")
	procQueryString = dll.MustFindProc("eglQueryString")
	procQuerySurface = dll.MustFindProc("eglQuerySurface")
	procReleaseTexImage = dll.MustFindProc("eglReleaseTexImage")
	procReleaseThread = dll.MustFindProc("eglReleaseThread")
	procSurfaceAttrib = dll.MustFindProc("eglSurfaceAttrib")
	procSwapBuffers = dll.MustFindProc("eglSwapBuffers")
	procSwapInterval = dll.MustFindProc("eglSwapInterval")
	procTerminate = dll.MustFindProc("eglTerminate")
	procWaitClient = dll.MustFindProc("eglWaitClient")
	procWaitGL = dll.MustFindProc("eglWaitGL")
	procWaitNative = dll.MustFindProc("eglWaitNative")
	procWaitSync = dll.MustFindProc("eglWaitSync")

}

func ChooseConfig(display Display, attribList []int32) (config Config, numConfigs int32, err error) {
	r0, _, _ := procChooseConfig.Call(
		uintptr(display),
		uintptr(unsafe.Pointer(&attribList[0])),
		uintptr(unsafe.Pointer(&config)),
		1,
		uintptr(unsafe.Pointer(&numConfigs)),
	)
	if r0 == FALSE {
		return 0, 0, GetError()
	}
	return config, numConfigs, nil
}

func CopyBuffers(display Display, surface Surface, target unsafe.Pointer /*EGLNativePixmapType*/) int {
	r0, _, _ := procCopyBuffers.Call(
		uintptr(display),
		uintptr(surface),
		uintptr(target))
	return int(r0)
}

func CreateContext(display Display, config Config, shareContext Context, contextAttribs []int32) (Context, error) {
	r0, _, _ := procCreateContext.Call(
		uintptr(display),
		uintptr(config),
		uintptr(shareContext),
		uintptr(unsafe.Pointer(&contextAttribs[0])),
	)
	context := Context(r0)
	if context == NO_CONTEXT {
		return context, GetError()
	}
	return context, nil
}

func CreatePbufferSurface(
	display Display,
	config Config,
	attrib_list unsafe.Pointer /*const EGLint **/) Surface {
	r0, _, _ := procCreatePbufferSurface.Call(
		uintptr(display),
		uintptr(config),
		uintptr(attrib_list))
	return Surface(r0)
}

func CreatePixmapSurface(
	display Display,
	config Config,
	pixmap unsafe.Pointer, /*EGLNativePixmapType*/
	attribList unsafe.Pointer /*const EGLint **/) Surface {
	r0, _, _ := procCreatePixmapSurface.Call(
		uintptr(display),
		uintptr(config),
		uintptr(pixmap),
		uintptr(attribList))
	return Surface(r0)
}

func CreateWindowSurface(
	display Display,
	config Config,
	nativeWindow NativeWindow,
	attribList []int32) (surface Surface, err error) {
	var attribListPtr uintptr
	if attribList != nil && len(attribList) > 0 {
		attribListPtr = uintptr(unsafe.Pointer(&attribList[0]))
	}
	r0, _, _ := procCreateWindowSurface.Call(
		uintptr(display),
		uintptr(config),
		uintptr(nativeWindow),
		attribListPtr)

	surface = Surface(r0)
	if surface == NO_SURFACE {
		return surface, GetError()
	}
	return surface, nil
}

func DestroyContext(display Display, ctx Context) int {
	r0, _, _ := procDestroyContext.Call(
		uintptr(display),
		uintptr(ctx))
	return int(r0)
}

func DestroySurface(display Display, surface Surface) int {
	r0, _, _ := procDestroySurface.Call(
		uintptr(display),
		uintptr(surface))
	return int(r0)
}

func GetConfigAttrib(display Display, config Config, attribute int, value unsafe.Pointer /*EGLint **/) int {
	r0, _, _ := procGetConfigAttrib.Call(
		uintptr(display),
		uintptr(config),
		uintptr(attribute),
		uintptr(value))
	return int(r0)
}

func GetConfigs(display Display, configs unsafe.Pointer /*Config**/, configSize int, numConfig unsafe.Pointer /*EGLint **/) int {
	r0, _, _ := procGetConfigs.Call(
		uintptr(display),
		uintptr(configs),
		uintptr(configSize),
		uintptr(numConfig))
	return int(r0)
}

func GetCurrentDisplay() Display {
	r0, _, _ := procGetCurrentDisplay.Call()
	return Display(r0)
}

func GetCurrentSurface(readdraw int) Surface {
	r0, _, _ := procGetCurrentSurface.Call(uintptr(readdraw))
	return Surface(r0)
}

func GetDisplay(display_id unsafe.Pointer /*EGLNativeDisplayType*/) Display {
	r0, _, _ := procGetDisplay.Call(uintptr(display_id))
	return Display(r0)
}

func GetError() error {
	ret, _, _ := procGetError.Call()
	switch int(ret) {
	case SUCCESS:
		return nil
	case NOT_INITIALIZED:
		return ErrNotInitialized
	case BAD_ACCESS:
		return ErrBadAccess
	case BAD_ALLOC:
		return ErrBadAlloc
	case BAD_ATTRIBUTE:
		return ErrBadAttribute
	case BAD_CONTEXT:
		return ErrBadContext
	case BAD_CONFIG:
		return ErrBadConfig
	case BAD_CURRENT_SURFACE:
		return ErrBadCurrentSurface
	case BAD_DISPLAY:
		return ErrBadDisplay
	case BAD_SURFACE:
		return ErrBadSurface
	case BAD_MATCH:
		return ErrBadMatch
	case BAD_PARAMETER:
		return ErrBadParameter
	case BAD_NATIVE_PIXMAP:
		return ErrBadNativePixmap
	case BAD_NATIVE_WINDOW:
		return ErrBadNativeWindow
	case CONTEXT_LOST:
		return ErrContextLost
	default:
		return ErrUnknown
	}
}

func GetProcAddress(procname unsafe.Pointer /*const char **/) unsafe.Pointer /*__eglMustCastToProperFunctionPointerType*/ {
	r0, _, _ := procGetProcAddress.Call(uintptr(procname))
	return unsafe.Pointer /*__eglMustCastToProperFunctionPointerType*/ (r0)
}

func Initialize(display Display) (int, int, error) {
	var major, minor int
	r0, _, _ := procInitialize.Call(
		uintptr(display),
		uintptr(unsafe.Pointer(&major)),
		uintptr(unsafe.Pointer(&minor)))
	if r0 == FALSE {
		return 0, 0, GetError()
	}
	return major, minor, nil
}

func MakeCurrent(display Display, draw Surface, read Surface, context Context) error {
	r0, _, _ := procMakeCurrent.Call(
		uintptr(display),
		uintptr(draw),
		uintptr(read),
		uintptr(context))

	if r0 == FALSE {
		return GetError()
	}
	return nil
}
func QueryContext(
	display Display,
	ctx Context,
	attribute int,
	value unsafe.Pointer /*EGLint **/) int {
	r0, _, _ := procQueryContext.Call(
		uintptr(display),
		uintptr(ctx),
		uintptr(attribute),
		uintptr(value))
	return int(r0)
}

func QueryString(
	display Display,
	name int) unsafe.Pointer /*const char **/ {
	r0, _, _ := procQueryString.Call(
		uintptr(display),
		uintptr(name))
	return unsafe.Pointer /*const char **/ (r0)
}

func QuerySurface(
	display Display,
	surface Surface,
	attribute int,
	value unsafe.Pointer /*EGLint **/) int {
	r0, _, _ := procQuerySurface.Call(
		uintptr(display),
		uintptr(surface),
		uintptr(attribute),
		uintptr(value))
	return int(r0)
}

func SwapBuffers(display Display, surface Surface) bool {
	r0, _, _ := procSwapBuffers.Call(uintptr(display), uintptr(surface))
	return r0 != FALSE
}

func Terminate(display Display) int {
	r0, _, _ := procTerminate.Call(uintptr(display))
	return int(r0)
}

func WaitGL() int {
	r0, _, _ := procWaitGL.Call()
	return int(r0)
}

func WaitNative(engine int) int {
	r0, _, _ := procWaitNative.Call(uintptr(engine))
	return int(r0)
}

func BindTexImage(
	display Display,
	surface Surface,
	buffer int) int {
	r0, _, _ := procBindTexImage.Call(
		uintptr(display),
		uintptr(surface),
		uintptr(buffer))
	return int(r0)
}

func ReleaseTexImage(display Display, surface Surface, buffer int) int {
	r0, _, _ := procReleaseTexImage.Call(
		uintptr(display),
		uintptr(surface),
		uintptr(buffer))
	return int(r0)
}

func SurfaceAttrib(display Display, surface Surface, attribute int, value int) int {
	r0, _, _ := procSurfaceAttrib.Call(
		uintptr(display),
		uintptr(surface),
		uintptr(attribute),
		uintptr(value))
	return int(r0)
}

func SwapInterval(display Display, interval int) error {
	r0, _, _ := procSwapInterval.Call(uintptr(display), uintptr(interval))
	if r0 == 0 {
		return GetError()
	}
	return nil
}

func BindAPI(api uint32) error {
	r0, _, _ := procBindAPI.Call(uintptr(api))
	if r0 == FALSE {
		return GetError()
	}
	return nil
}

func QueryAPI() uint32 {
	r0, _, _ := procQueryAPI.Call()
	return uint32(r0)
}

func CreatePbufferFromClientBuffer(
	display Display,
	buftype uint32,
	buffer unsafe.Pointer, /*EGLClientBuffer*/
	config Config,
	attrib_list unsafe.Pointer /*const EGLint **/) Surface {
	r0, _, _ := procCreatePbufferFromClientBuffer.Call(
		uintptr(display),
		uintptr(buftype),
		uintptr(buffer),
		uintptr(config),
		uintptr(attrib_list))
	return Surface(r0)
}

func ReleaseThread() int {
	r0, _, _ := procReleaseThread.Call()
	return int(r0)
}

func WaitClient() int {
	r0, _, _ := procWaitClient.Call()
	return int(r0)
}

func GetCurrentContext() Context {
	r0, _, _ := procGetCurrentContext.Call()
	return Context(r0)
}

func CreateSync(
	display Display,
	ty uint32,
	attrib_list unsafe.Pointer /*const EGLAttrib **/) unsafe.Pointer /*EGLSync*/ {
	r0, _, _ := procCreateSync.Call(
		uintptr(display),
		uintptr(ty),
		uintptr(attrib_list))
	return unsafe.Pointer /*EGLSync*/ (r0)
}

func DestroySync(
	display Display,
	sync unsafe.Pointer /*EGLSync*/) int {
	r0, _, _ := procDestroySync.Call(
		uintptr(display),
		uintptr(sync))
	return int(r0)
}

func ClientWaitSync(
	display Display,
	sync unsafe.Pointer, /*EGLSync*/
	flags int,
	timeout int64) int {
	r0, _, _ := procClientWaitSync.Call(
		uintptr(display),
		uintptr(sync),
		uintptr(flags),
		uintptr(timeout))
	return int(r0)
}

func GetSyncAttrib(
	display Display,
	sync unsafe.Pointer, /*EGLSync*/
	attribute int,
	value unsafe.Pointer /*EGLAttrib **/) int {
	r0, _, _ := procGetSyncAttrib.Call(
		uintptr(display),
		uintptr(sync),
		uintptr(attribute),
		uintptr(value))
	return int(r0)
}

func CreateImage(
	display Display,
	ctx Context,
	target uint32,
	buffer unsafe.Pointer, /*EGLClientBuffer*/
	attrib_list unsafe.Pointer /*const EGLAttrib **/) unsafe.Pointer /*EGLImage*/ {
	r0, _, _ := procCreateImage.Call(
		uintptr(display),
		uintptr(ctx),
		uintptr(target),
		uintptr(buffer),
		uintptr(attrib_list))
	return unsafe.Pointer /*EGLImage*/ (r0)
}

func DestroyImage(
	display Display,
	image unsafe.Pointer /*EGLImage*/) int {
	r0, _, _ := procDestroyImage.Call(
		uintptr(display),
		uintptr(image))
	return int(r0)
}

func GetPlatformDisplay(
	platform uint32,
	nativeDisplay unsafe.Pointer, /*void **/
	attribList unsafe.Pointer /*const EGLAttrib **/) Display {
	r0, _, _ := procGetPlatformDisplay.Call(
		uintptr(platform),
		uintptr(nativeDisplay),
		uintptr(attribList))
	return Display(r0)
}

func CreatePlatformWindowSurface(
	display Display,
	config Config,
	nativeWindow unsafe.Pointer, /*void **/
	attribList unsafe.Pointer /*const EGLAttrib **/) Surface {
	r0, _, _ := procCreatePlatformWindowSurface.Call(
		uintptr(display),
		uintptr(config),
		uintptr(nativeWindow),
		uintptr(attribList))
	return Surface(r0)
}

func CreatePlatformPixmapSurface(
	display Display,
	config Config,
	nativePixmap unsafe.Pointer, /*void **/
	attribList unsafe.Pointer /*const EGLAttrib **/) Surface {
	r0, _, _ := procCreatePlatformPixmapSurface.Call(
		uintptr(display),
		uintptr(config),
		uintptr(nativePixmap),
		uintptr(attribList))
	return Surface(r0)
}

func WaitSync(
	display Display,
	sync unsafe.Pointer, /*EGLSync*/
	flags int) int {
	r0, _, _ := procWaitSync.Call(
		uintptr(display),
		uintptr(sync),
		uintptr(flags))
	return int(r0)
}

func GetPlatformDisplayEXT(platform uint32, nativeDisplay NativeDisplay, attribList []int32) (Display, error) {
	r0, _, _ := procGetPlatformDisplayEXT.Call(
		uintptr(platform),
		uintptr(nativeDisplay),
		uintptr(unsafe.Pointer(&attribList[0])))
	display := Display(r0)
	if display == NO_DISPLAY {
		return NO_DISPLAY, GetError()
	}
	return display, nil
}
