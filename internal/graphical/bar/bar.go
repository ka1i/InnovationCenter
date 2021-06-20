package bar

import "github.com/progrium/macdriver/cocoa"

const (
	MenuTitle = "☘️"
)

func InitBar(win *cocoa.NSWindow) {
	statusBar := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
	statusBar.Retain()
	statusBar.Button().SetTitle(MenuTitle)

	statusBar.SetMenu(*Menu(win))
}
