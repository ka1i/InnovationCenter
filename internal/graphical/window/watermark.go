package window

import (
	"fmt"
	"log"

	"github.com/ka1i/InnovationCenter/pkg/version"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
)

func WaterMark() {
	watermark := fmt.Sprintf("InnovationCenter Pro Inside Preview\nMardan(ka1i) Build Version:%s", version.VERSION)
	screen := cocoa.NSScreen_Main().Frame().Size
	windowSize := core.Rect(5, 10, screen.Width*0.25, screen.Height*0.05)

	log.Printf("Create [WaterMark] %v\n", windowSize)
	t := cocoa.NSTextView_Init(core.Rect(0, 0, screen.Width*0.25, screen.Height*2*0.018))
	t.SetString(watermark)
	t.SetFont(cocoa.Font("Monaco", 18))
	t.SetTextColor(cocoa.Color(1, 1, 1, 0.85))
	t.SetEditable(false)
	t.SetImportsGraphics(false)
	t.SetDrawsBackground(false)

	c := cocoa.NSView_Init(core.Rect(0, 0, screen.Width*0.25, screen.Height*0.05))
	c.AddSubviewPositionedRelativeTo(t, cocoa.NSWindowAbove, nil)

	w := cocoa.NSWindow_Init(windowSize,
		cocoa.NSBorderlessWindowMask,
		cocoa.NSBackingStoreBuffered,
		false,
	)
	w.SetContentView(c)
	w.SetTitlebarAppearsTransparent(true)
	w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
	w.SetOpaque(false)
	w.SetBackgroundColor(cocoa.NSColor_Clear())
	w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
	w.SetFrameDisplay(windowSize, true)
	w.MakeKeyAndOrderFront(nil)
}
