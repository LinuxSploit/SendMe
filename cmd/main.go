//go:generate fyne bundle -o ../bundled/bundled.go ../asset/
//go:generate sh -c "sed -i 's/resource/Resource/g' ../bundled/bundled.go && sed -i 's/main/bundled/g' ../bundled/bundled.go"

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	xtheme "fyne.io/x/fyne/theme"
	"github.com/LinuxSploit/SendMe/bundled"
	"github.com/LinuxSploit/SendMe/custom/TriTabNav"
	"github.com/LinuxSploit/SendMe/ui"
	"github.com/LinuxSploit/SendMe/ui/activity"
	"github.com/LinuxSploit/SendMe/ui/home"
	"github.com/LinuxSploit/SendMe/ui/shared"
)

func main() {
	// create a app object
	a := app.New()
	settings := a.Settings()
	if settings != nil {
		settings.SetTheme(xtheme.AdwaitaTheme())

	}
	w := a.NewWindow("SendMe")
	w.Resize(fyne.NewSize(300, 550))
	// init app ui
	ui.Init(w)
	// create a custom AppTabs widget object
	AppTabs := TriTabNav.NewTriTabNav(

		bundled.ResourceHomeBlueSvg,
		bundled.ResourceHomeWhiteSvg,

		bundled.ResourceShareBlueSvg,
		bundled.ResourceShareWhiteSvg,

		bundled.ResourceActivityBlueSvg,
		bundled.ResourceActivityWhiteSvg,

		home.HomeTab,         // home tab
		shared.SharedTab,     // shared tab
		activity.ActivityTab, // activity tab
	)

	// setting custom AppTabs as main screen
	w.SetContent(
		AppTabs.AppScreen,
	)
	// launching app
	w.ShowAndRun()
}
