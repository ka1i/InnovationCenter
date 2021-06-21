package window

import (
	"log"

	"github.com/ka1i/InnovationCenter/internal/graphical/bar"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/webkit"
)

func Browser() {
	config := webkit.WKWebViewConfiguration_New()
	config.Preferences().SetValueForKey(core.True, core.String("developerExtrasEnabled"))

	screen := cocoa.NSScreen_Main().Frame().Size
	windowSize := core.Rect(0, 0, screen.Width, screen.Height-25)

	log.Printf("Create [Browser] %v\n", windowSize)
	url := core.URL("https://www.mardan.wiki")
	req := core.NSURLRequest_Init(url)

	wv := webkit.WKWebView_Init(windowSize, config)
	wv.Retain()
	wv.SetOpaque(false)
	wv.SetBackgroundColor(cocoa.NSColor_Clear())
	wv.SetValueForKey(core.False, core.String("drawsBackground"))
	wv.LoadRequest(req)

	win := cocoa.NSWindow_Init(windowSize,
		cocoa.NSClosableWindowMask|cocoa.NSBorderlessWindowMask,
		cocoa.NSBackingStoreBuffered, false)
	win.SetContentView(wv)
	win.SetBackgroundColor(cocoa.NSColor_Clear())
	win.SetOpaque(false)
	win.SetTitleVisibility(cocoa.NSWindowTitleHidden)
	win.SetTitlebarAppearsTransparent(true)
	win.SetIgnoresMouseEvents(true)
	win.SetLevel(cocoa.NSMainMenuWindowLevel - 25)
	win.MakeKeyAndOrderFront(win)
	win.SetCollectionBehavior(cocoa.NSWindowCollectionBehaviorCanJoinAllSpaces)
	win.Send("setHasShadow:", false)

	bar.InitBar(&win)
}
