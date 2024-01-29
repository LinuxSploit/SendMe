package home

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/LinuxSploit/SendMe/custom/switchBtn"
	"github.com/LinuxSploit/SendMe/server"
	"github.com/LinuxSploit/SendMe/ui/home/homeTabs"
)

var (
	// Home Tab container
	HomeTab *fyne.Container
	// Point to HTTP server
	ServerRef *server.Server
	// Copy button for HTTP server address
	AddressBtn *widget.Button
)

func Init(w fyne.Window) {
	// init home sub tabs [ "Actives" , "Requests" ]
	homeTabs.Init()
	// app title "SendMe" as logo
	title := canvas.NewText("SendMe", theme.PrimaryColor())
	title.TextSize = 22
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	// UI Component to display server's URL
	AddressBtn = widget.NewButtonWithIcon("http://0.0.0.0:4556", theme.ContentCopyIcon(), func() {
		log.Println(AddressBtn.Text)
	})
	AddressBtn.Hide()
	// a custom switch button widget to turn server ON/OFF
	var serverSwitch *switchBtn.SwitchButton
	serverSwitch = switchBtn.NewSwitchButton(
		"Stop",
		"Start",
		theme.MediaStopIcon(),
		theme.MediaPlayIcon(),
		widget.DangerImportance,
		widget.HighImportance,
		func() {
			AddressBtn.Show()
			serverSwitch.ToggleSwitchState()
			log.Println(serverSwitch.Status)

			ServerRef = server.NewServer(":4556")
			// NOTE: passing serverSwitch, AddressBtn, w so that if any error occur turn serverSwitch off,
			// hide AddressBtn and display error message
			err := ServerRef.Start(serverSwitch, AddressBtn, w)
			if err != nil {
				log.Println(err)
				AddressBtn.Hide()
				serverSwitch.ToggleSwitchState()
				dialog.ShowError(err, w)

			}
		},
		func() {
			serverSwitch.ToggleSwitchState()
			log.Println("State: OFF")

			err := ServerRef.Stop()
			if err != nil {
				log.Println(err)
				return
			}

			AddressBtn.Hide()

		},
	)

	// home tab display
	HomeTab = container.NewBorder(
		container.NewVBox(
			title,      // app title logo
			AddressBtn, // Copy button for HTTP server address
		),
		container.NewVBox(
			serverSwitch,          // Server Start/Stop Button at the bottom
			widget.NewSeparator(), // line separator above TriTabNav bar
		),
		nil,
		nil,
		container.NewStack(homeTabs.HomeTabs.AppScreen), // home subtabs [actives,request] content
	)
}
