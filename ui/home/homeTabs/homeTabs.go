package homeTabs

import (
	"fyne.io/fyne/v2/theme"
	"github.com/LinuxSploit/SendMe/custom/DualTabNav"
	"github.com/LinuxSploit/SendMe/ui/home/homeTabs/actives"
	"github.com/LinuxSploit/SendMe/ui/home/homeTabs/requests"
)

var (
	HomeTabs *DualTabNav.DualTabNav
)

func Init() {
	actives.Init()  // a tab within home tab
	requests.Init() // a tab within home tab

	HomeTabs = DualTabNav.NewDualTabNav(
		"Actives",
		"Requests",
		theme.AccountIcon(),
		theme.HelpIcon(),
		actives.ActiveTab,
		requests.RequestTab,
	)

}
