package activity

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/LinuxSploit/SendMe/internal/download"
)

var (
	Activities     []*download.DownloadActivity
	ActivitiesList *widget.List
	ActivityTab    *fyne.Container
)

func Init() {

	ActivitiesList = widget.NewList(
		func() int { return len(Activities) },
		func() fyne.CanvasObject {
			return widget.NewCard(
				"Username",
				"filename.ext",
				widget.NewProgressBar(),
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {

			Activities[lii].ProgressBar.SetValue(float64(Activities[lii].DownloadedSize / Activities[lii].FileSize))

			co.(*widget.Card).SetTitle(Activities[lii].User.Username)
			co.(*widget.Card).SetSubTitle(Activities[lii].FileName)
			co.(*widget.Card).SetContent(
				Activities[lii].ProgressBar,
			)
		},
	)

	tabLabel := canvas.NewText("Activies", color.RGBA{53, 132, 228, 255})
	tabLabel.TextStyle.Bold = true

	ActivityTab =
		container.NewBorder(
			container.NewCenter(
				tabLabel,
			),
			nil,
			nil,
			nil,
			ActivitiesList,
		) // Downloads resources list
}
