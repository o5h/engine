package main

import (
	"log"
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/o5h/opengles/egl"
	"github.com/o5h/opengles/gl"
	"github.com/o5h/winapi"
	"github.com/o5h/winapi/kernel32"
	"github.com/o5h/winapi/user32"
	"golang.org/x/sys/windows"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	ctx := Create("Example", 640, 480)
	ctx.MainLoop()
}

type context struct {
	nativeWindow  egl.NativeWindow
	nativeDisplay egl.NativeDisplay
	Context       egl.Context
	Surface       egl.Surface
	Display       egl.Display
	handle        cgo.Handle
	title         string
	width, height int32
	Done          chan bool
}

func Create(title string, w, h int32) *context {
	ctx := &context{
		title:  title,
		width:  w,
		height: h,
		Done:   make(chan bool, 1),
	}
	ctx.handle = cgo.NewHandle(ctx)

	ctx.createWindow(w, h)
	return ctx
}

func (ctx *context) createWindow(w, h int32) {
	wndproc := winapi.WNDPROC(windows.NewCallback(wndProc))
	mh, _ := kernel32.GetModuleHandle(nil)
	myicon, _ := user32.LoadIconW(0, user32.IDI_APPLICATION)
	mycursor, _ := user32.LoadCursorW(0, user32.IDC_ARROW)

	var wc user32.WNDCLASSEX
	wc.Size = uint32(unsafe.Sizeof(wc))
	wc.WndProc = wndproc
	wc.Instance = winapi.HINSTANCE(mh)
	wc.Icon = myicon
	wc.Cursor = mycursor
	wc.Background = user32.COLOR_BTNFACE + 1
	wc.MenuName = nil
	wcname, _ := windows.UTF16PtrFromString("OPENGLES_WindowClass")
	wc.ClassName = wcname
	wc.IconSm = myicon
	user32.RegisterClassExW(&wc)

	windowTitle, _ := windows.UTF16PtrFromString(ctx.title)
	user32.CreateWindowExW(
		0,
		wcname,
		windowTitle,
		// No border, no title
		user32.WS_POPUP|user32.WS_CLIPSIBLINGS|user32.WS_CLIPCHILDREN|user32.WS_OVERLAPPEDWINDOW,
		user32.CW_USEDEFAULT,
		user32.CW_USEDEFAULT,
		w,
		h,
		winapi.HWND(0),
		winapi.HMENU(0),
		winapi.HINSTANCE(mh),
		winapi.LPVOID(ctx.handle))
}

func (ctx *context) onPaint() {
	gl.ClearColor(1, 0, 1, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

}

func (ctx *context) initialize(hWnd winapi.HWND) error {

	ctx.nativeWindow = egl.NativeWindow(hWnd)
	dc := user32.GetDC(hWnd)
	ctx.nativeDisplay = egl.NativeDisplay(dc)

	user32.SetWindowLongPtrW(hWnd, user32.GWLP_USERDATA, winapi.LONG_PTR(ctx.handle))

	var err error

	ctx.Context, ctx.Display, ctx.Surface, err = egl.CreateEGLSurface(ctx.nativeDisplay, ctx.nativeWindow)
	if err != nil {
		return err
	}
	if err = egl.MakeCurrent(ctx.Display, ctx.Surface, ctx.Surface, ctx.Context); err != nil {
		return err
	}

	if err = egl.SwapInterval(ctx.Display, 1); err != nil {
		return err
	}

	user32.SetWindowPos(hWnd, user32.HWND_TOP, 0, 0, ctx.width, ctx.height, user32.SWP_SHOWWINDOW)
	user32.ShowWindow(hWnd, user32.SW_SHOW)
	user32.UpdateWindow(hWnd)
	user32.SetFocus(hWnd)
	return err
}

func (ctx *context) MainLoop() {
	defer ctx.onDestroy()
	for {
		select {
		case <-ctx.Done:
			return
		default:
			var message user32.Msg
			if ok, _ := user32.PeekMessageW(&message, 0, 0, 0, user32.PM_REMOVE); ok == winapi.TRUE {
				user32.TranslateMessage(&message)
				user32.DispatchMessageW(&message)
			} else {
				user32.SendMessageW(winapi.HWND(ctx.nativeWindow), user32.WM_PAINT, 0, 0)
			}
		}
	}

}

func (ctx *context) onDestroy() {
	user32.DestroyWindow(winapi.HWND(ctx.nativeWindow))
	ctx.handle.Delete()
}

func wndProc(hWnd winapi.HWND, msg winapi.UINT, wParam winapi.WPARAM, lParam winapi.LPARAM) (rc winapi.LRESULT) {
	var ctx *context
	if ptr, _ := user32.GetWindowLongPtrW(hWnd, user32.GWLP_USERDATA); ptr != 0 {
		ctx = cgo.Handle(ptr).Value().(*context)
	}
	switch msg {
	case user32.WM_CREATE:
		create := (*user32.CREATESTRUCTW)(unsafe.Pointer(lParam))
		ctx := cgo.Handle(create.CreateParams).Value().(*context)
		err := ctx.initialize(hWnd)
		if err != nil {
			log.Fatal(err)
		}
	case user32.WM_PAINT:
		ctx.onPaint()
		egl.SwapBuffers(ctx.Display, ctx.Surface)
	case user32.WM_SIZE:
		w := int(winapi.LOWORD(winapi.DWORD(lParam)))
		h := int(winapi.HIWORD(winapi.DWORD(lParam)))
		log.Println(w, h)
	case user32.WM_CLOSE:
		ctx.Done <- true
	case user32.WM_DESTROY:
		user32.PostQuitMessage(0)
	case user32.WM_KEYDOWN, user32.WM_KEYUP:
		if wParam == user32.VK_ESCAPE {
			ctx.Done <- true
		}
	default:
		rc = user32.DefWindowProcW(hWnd, msg, wParam, lParam)
	}
	return
}
