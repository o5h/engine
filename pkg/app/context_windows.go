package app

import (
	"log"
	"runtime/cgo"
	"unsafe"

	"github.com/o5h/engine/pkg/app/input/keyboard"
	"github.com/o5h/engine/pkg/app/input/mouse"
	"github.com/o5h/opengles/egl"
	"github.com/o5h/winapi"
	"github.com/o5h/winapi/kernel32"
	"github.com/o5h/winapi/user32"
	"golang.org/x/sys/windows"
)

type windowsContext struct {
	app           Application
	Width, Height int32

	hWnd winapi.HWND
	hDC  winapi.HDC

	NativeWindow  egl.NativeWindow
	NativeDisplay egl.NativeDisplay
	Context       egl.Context
	Surface       egl.Surface
	Display       egl.Display

	handle cgo.Handle
	done   chan struct{}
}

func (ctx *windowsContext) Done() {
	ctx.done <- struct{}{}
}

func newContext(app Application, cfg *Config) *windowsContext {
	ctx := &windowsContext{
		app:  app,
		done: make(chan struct{}, 1),
	}
	ctx.handle = cgo.NewHandle(ctx)
	ctx.createWindow(cfg)
	return ctx
}

func (ctx *windowsContext) createWindow(cfg *Config) {
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

	windowTitle, _ := windows.UTF16PtrFromString(cfg.Title)
	user32.CreateWindowExW(
		0,
		wcname,
		windowTitle,
		// No border, no title
		user32.WS_POPUP|user32.WS_CLIPSIBLINGS|user32.WS_CLIPCHILDREN|user32.WS_OVERLAPPEDWINDOW,
		user32.CW_USEDEFAULT,
		user32.CW_USEDEFAULT,
		cfg.Width,
		cfg.Height,
		winapi.HWND(0),
		winapi.HMENU(0),
		winapi.HINSTANCE(mh),
		winapi.LPVOID(ctx.handle))
}

func (ctx *windowsContext) onCreate(hWnd winapi.HWND) error {
	ctx.hWnd = hWnd
	ctx.hDC = user32.GetDC(hWnd)
	ctx.NativeWindow = egl.NativeWindow(hWnd)
	ctx.NativeDisplay = egl.NativeDisplay(ctx.hDC)

	user32.SetWindowLongPtrW(hWnd, user32.GWLP_USERDATA, winapi.LONG_PTR(ctx.handle))

	var err error
	if ctx.Context, ctx.Display, ctx.Surface, err = egl.CreateEGLSurface(ctx.NativeDisplay, ctx.NativeWindow); err != nil {
		return err
	}
	if err = egl.MakeCurrent(ctx.Display, ctx.Surface, ctx.Surface, ctx.Context); err != nil {
		return err
	}

	if err = egl.SwapInterval(ctx.Display, 1); err != nil {
		return err
	}

	//user32.SetWindowPos(hWnd, user32.HWND_TOP, 0, 0, ctx.Width, ctx.Height, user32.SWP_SHOWWINDOW)
	ctx.app.OnCreate(ctx)
	user32.ShowWindow(hWnd, user32.SW_SHOW)
	user32.SetFocus(hWnd)
	user32.UpdateWindow(hWnd)
	return err
}

func (ctx *windowsContext) onDestroy() {
	user32.ReleaseDC(ctx.hWnd, ctx.hDC)
	ctx.app.OnDestroy()
	ctx.handle.Delete()
}

func (ctx *windowsContext) onResize(w, h int32) {
}

func (ctx *windowsContext) onPaint() {
	ctx.app.OnUpdate(0)
}

func (ctx *windowsContext) mainLoop() {
	for {
		select {
		case <-ctx.done:
			return
		default:
			var message user32.Msg
			if ok, _ := user32.PeekMessageW(&message, 0, 0, 0, user32.PM_REMOVE); ok == winapi.TRUE {
				user32.TranslateMessage(&message)
				user32.DispatchMessageW(&message)
			} else {
				user32.SendMessageW(winapi.HWND(ctx.NativeWindow), user32.WM_PAINT, 0, 0)
			}
		}
	}

}
func (ctx *windowsContext) onKey(msg winapi.UINT, wParam winapi.WPARAM, lParam winapi.LPARAM) {
	code := WindowsVKToCode(wParam)
	r := rune(user32.MapVirtualKeyW(winapi.UINT(wParam), user32.MAPVK_VK_TO_CHAR))
	var dir = keyboard.Press
	switch msg {
	case user32.WM_KEYDOWN:
		dir = keyboard.Press
	case user32.WM_KEYUP:
		dir = keyboard.Release
	}
	go keyboard.Events.Next(keyboard.Event{Direction: dir, Code: code, Rune: r})
}

func (ctx *windowsContext) onMouse(hWnd winapi.HWND, msg winapi.UINT, lParam winapi.LPARAM) {
	var action mouse.Action
	var btn mouse.Button
	x := int(winapi.GET_X_LPARAM(lParam))
	y := int(winapi.GET_Y_LPARAM(lParam))

	switch msg {
	case user32.WM_MOUSEMOVE:
		action = mouse.ActionNone
		btn = mouse.ButtonNone
	case user32.WM_LBUTTONDOWN:
		action = mouse.ActionPress
		btn = mouse.ButtonLeft
	case user32.WM_RBUTTONDOWN:
		action = mouse.ActionPress
		btn = mouse.ButtonLeft
	default:
		log.Println("mouse", msg)
	}
	mouse.Events.Next(mouse.Event{Action: action, X: x, Y: y, Button: btn})
}

func wndProc(hWnd winapi.HWND, msg winapi.UINT, wParam winapi.WPARAM, lParam winapi.LPARAM) (rc winapi.LRESULT) {
	var ctx *windowsContext
	if ptr, _ := user32.GetWindowLongPtrW(hWnd, user32.GWLP_USERDATA); ptr != 0 {
		ctx = cgo.Handle(ptr).Value().(*windowsContext)
	}
	switch msg {

	case user32.WM_CREATE:
		create := (*user32.CREATESTRUCTW)(unsafe.Pointer(lParam))
		ctx := cgo.Handle(create.CreateParams).Value().(*windowsContext)
		err := ctx.onCreate(hWnd)
		if err != nil {
			log.Fatal(err)
		}

	case user32.WM_PAINT:
		ctx.onPaint()
		egl.SwapBuffers(ctx.Display, ctx.Surface)

	case user32.WM_SIZE:
		w := int32(winapi.LOWORD(winapi.DWORD(lParam)))
		h := int32(winapi.HIWORD(winapi.DWORD(lParam)))
		ctx.onResize(w, h)

	case user32.WM_CLOSE:
		ctx.Done()

	case user32.WM_DESTROY:
		user32.PostQuitMessage(0)

	case user32.WM_KEYDOWN, user32.WM_KEYUP:
		ctx.onKey(msg, wParam, lParam)

	case user32.WM_LBUTTONDOWN, user32.WM_LBUTTONUP, user32.WM_MOUSEMOVE:
		ctx.onMouse(hWnd, msg, lParam)

	default:
		rc = user32.DefWindowProcW(hWnd, msg, wParam, lParam)
	}
	return
}
