package ui

import (
	"fyne.io/fyne/v2"
	"github.com/LinuxSploit/SendMe/ui/activity"
	"github.com/LinuxSploit/SendMe/ui/home"
	"github.com/LinuxSploit/SendMe/ui/shared"
)

func Init(w fyne.Window) {
	home.Init(w)
	shared.Init(w)
	activity.Init()
}
