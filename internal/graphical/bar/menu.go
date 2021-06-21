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
	menuService := cocoa.NSMenuItem_New()
	menuService.SetTitle("Service")

	menuService.SetSubmenu(serviceMenu())

	menuQuit := cocoa.NSMenuItem_New()
	menuQuit.SetTitle("Quit")
	menuQuit.SetAction(objc.Sel("terminate:"))

	menu := cocoa.NSMenu_New()
	menu.SetAutoenablesItems(false)
	menu.AddItem(menuEnabled)
	menu.AddItem(cocoa.NSMenuItem_Separator())
	menu.AddItem(menuService)
	menu.AddItem(cocoa.NSMenuItem_Separator())
	menu.AddItem(menuQuit)
	return &menu
}

func serviceMenu() cocoa.NSMenu {
	menuServiceMenu := cocoa.NSMenu_New()
	menus1 := cocoa.NSMenuItem_New()
	menus1.SetTitle("系统门户")
	menus1.SetState(1)
	menus2 := cocoa.NSMenuItem_New()
	menus2.SetTitle("日志中心")
	menus2.SetState(1)
	menus3 := cocoa.NSMenuItem_New()
	menus3.SetTitle("文件中心")
	menus3.SetState(1)
	menuServiceMenu.AddItem(menus1)
	menuServiceMenu.AddItem(menus2)
	menuServiceMenu.AddItem(menus3)
	return menuServiceMenu
}
