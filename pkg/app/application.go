package app

import "runtime"

type Application interface {
	OnCreate(Context)
	OnUpdate(float32)
	OnDestroy()
}

func Start(app Application, options ...func(*Config)) {

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	cfg := DefaultConfig
	for _, opt := range options {
		opt(&cfg)
	}
	ctx := newContext(app, &cfg)
	defer ctx.onDestroy()
	ctx.mainLoop()
}
