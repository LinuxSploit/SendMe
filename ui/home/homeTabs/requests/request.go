package requests

import (
	"image/color"
	"log"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/LinuxSploit/SendMe/internal/request"
	"github.com/LinuxSploit/SendMe/internal/resource"
	"github.com/LinuxSploit/SendMe/internal/user"
)

var (
	Requests     []request.Request
	RequestsList *widget.List
	RequestTab   *fyne.Container
)

func Init() {
	demoRes, err := resource.NewResource("/home/linuxsploit/Documents/SendMe/go.mod")
	if err != nil {
		log.Fatal(err)
	}

	Requests = append(Requests,
		request.Request{
			User: &user.User{
				Username: "Khan Downloader",
				IP:       net.ParseIP("192.100.112.33"),
				Browser:  "Lite UC",
				Device:   "RouterOS",
				Online:   true,
			},
			Resource: demoRes,
		},
	)

	RequestsList = widget.NewList(
		func() int { return len(Requests) },
		func() fyne.CanvasObject {

			acceptBtn := widget.NewButtonWithIcon(
				"Accept", theme.ConfirmIcon(),
				func() {},
			)
			acceptBtn.Importance = widget.HighImportance

			rejectBtn := widget.NewButtonWithIcon(
				"Reject", theme.CancelIcon(),
				func() {},
			)
			rejectBtn.Importance = widget.DangerImportance

			return widget.NewCard(
				"Share Khan", // username
				"127.0.0.1",  // ip
				container.NewAdaptiveGrid(1,
					container.NewVBox(
						canvas.NewText("cod.rar", color.RGBA{255, 0, 0, 255}), // filename
						container.NewAdaptiveGrid(2,
							acceptBtn,
							rejectBtn,
						),
					),
				),
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {

			acceptBtn := widget.NewButtonWithIcon(
				"Accept", theme.ConfirmIcon(),
				func() {},
			)
			acceptBtn.Importance = widget.HighImportance

			rejectBtn := widget.NewButtonWithIcon(
				"Reject", theme.CancelIcon(),
				func() {},
			)
			rejectBtn.Importance = widget.DangerImportance

			co.(*widget.Card).SetTitle(Requests[lii].Username)
			co.(*widget.Card).SetSubTitle(Requests[lii].IP.String())
			co.(*widget.Card).SetContent(
				container.NewAdaptiveGrid(1,
					container.NewVBox(
						canvas.NewText(Requests[lii].FileName, color.RGBA{255, 0, 0, 255}), // filename
						container.NewAdaptiveGrid(2,
							acceptBtn,
							rejectBtn,
						),
					),
				),
			)
		},
	)

	RequestTab = container.NewStack(
		RequestsList,
	)
}
