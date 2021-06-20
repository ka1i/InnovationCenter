package graphical

import (
	"runtime"

	"github.com/ka1i/InnovationCenter/internal/graphical/window"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

func init() {
	runtime.LockOSThread()
}

func UserInterface() cocoa.NSApplication {
	cocoa.TerminateAfterWindowsClose = false

	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		// create window
		window.WaterMark()
		//window.WindowBrowser()

		// listen events
	})
	app.ActivateIgnoringOtherApps(true)
	return app
}
