package bar

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

func Menu(win *cocoa.NSWindow) *cocoa.NSMenu {
	menuEnabled := cocoa.NSMenuItem_New()
	menuEnabled.Retain()
	menuEnabled.SetTitle("Disable")
	menuEnabled.SetAction(objc.Sel("enabled:"))
	cocoa.DefaultDelegateClass.AddMethod("enabled:", func(_ objc.Object) {
		if win.IsVisible() {
			win.Send("orderOut:", win)
			menuEnabled.SetTitle("Enable")
		} else {
			win.Send("orderFront:", win)
			menuEnabled.SetTitle("Disable")
		}
	})

	menuQuit := cocoa.NSMenuItem_New()
	menuQuit.SetTitle("Quit")
	menuQuit.SetAction(objc.Sel("terminate:"))

	menu := cocoa.NSMenu_New()
	menu.SetAutoenablesItems(false)
	menu.AddItem(menuEnabled)
	menu.AddItem(cocoa.NSMenuItem_Separator())
	menu.AddItem(menuQuit)
	return &menu
}
