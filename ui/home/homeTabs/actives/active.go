package actives

import (
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/LinuxSploit/SendMe/internal/user"
)

var (
	ActiveTab   *fyne.Container
	ActiveList  *widget.List
	ActiveUsers []user.User
)

func Init() {

	// need to be optimize
	go func() {
		for {
			UpdateUserStatus()
			time.Sleep(time.Second)
		}
	}()

	ActiveList = widget.NewList(
		func() int { return len(ActiveUsers) },
		// active user template
		func() fyne.CanvasObject {
			return widget.NewCard(
				"Share Khan",
				"127.0.0.1",
				container.NewAdaptiveGrid(1,
					container.NewHBox(
						container.NewCenter(
							canvas.NewText(" Online ", color.RGBA{255, 0, 0, 255}),
						),
						widget.NewSeparator(),
						container.NewCenter(
							canvas.NewText("Safari", color.Black),
						),
					),
				),
			)
		},
		// active users info
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			content := container.NewAdaptiveGrid(1,
				container.NewHBox(
					container.NewCenter(
						canvas.NewText(ActiveUsers[lii].GetUserStatusWithColor()),
					),
					widget.NewSeparator(),
					container.NewCenter(
						canvas.NewText(ActiveUsers[lii].Browser, color.Black),
					),
				),
			)

			co.(*widget.Card).SetTitle(ActiveUsers[lii].Username)
			co.(*widget.Card).SetSubTitle(ActiveUsers[lii].IP.String() + "  " + ActiveUsers[lii].Device)
			co.(*widget.Card).SetContent(content)
		},
	)

	// Creating the ActiveTab container
	ActiveTab = container.NewStack(
		ActiveList,
	)

}

func UpdateUserStatus() {
	// need to be optimize
	for i, user := range ActiveUsers {
		dur := time.Since(user.LastSeen)
		if dur.Seconds() > 6 && ActiveUsers[i].Online {
			ActiveUsers[i].Online = false
			ActiveList.Refresh()

			// debug:
			log.Println(dur.Seconds())
			log.Println("Last Seen: ", user.LastSeen)
		}
	}
}
