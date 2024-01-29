package shared

import (
	"encoding/hex"
	"fmt"
	"log"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/LinuxSploit/SendMe/bundled"
	"github.com/LinuxSploit/SendMe/internal/resource"
	"github.com/gorilla/securecookie"
)

var (
	SharedTab                *fyne.Container
	SharedList               *widget.List
	SharedResources          []resource.Resource
	CurrentResourcesChecksum string
)

func randomizeChecksum() {
	// Generate 16 random bytes (128 bits)
	randomBytes := securecookie.GenerateRandomKey(64)

	// Convert the random bytes to a hexadecimal string
	hashString := hex.EncodeToString(randomBytes)
	CurrentResourcesChecksum = hashString
}

func Init(w fyne.Window) {
	randomizeChecksum()

	SharedList = widget.NewList(
		func() int { return len(SharedResources) },
		func() fyne.CanvasObject {
			tmp := widget.NewButtonWithIcon(
				"",
				theme.FileIcon(), // this icon will be replaced by [ bundled.ResourceProtectedPng, bundled.ResourceFilePng ]
				func() {},
			)
			tmp.Alignment = widget.ButtonAlignLeading
			tmp.Importance = widget.LowImportance
			return tmp
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			if SharedResources[lii].FileStatus {
				switch path.Ext(SharedResources[lii].FileName) {
				case ".mp4":
					co.(*widget.Button).SetIcon(bundled.ResourceFileMp4Svg)
				case ".pdf":
					co.(*widget.Button).SetIcon(bundled.ResourceFilePdfSvg)
				case ".zip":
					co.(*widget.Button).SetIcon(bundled.ResourceFileZipSvg)
				case ".png":
					co.(*widget.Button).SetIcon(bundled.ResourceFilePngSvg)

				default:
					co.(*widget.Button).SetIcon(bundled.ResourceFilePng)
				}
			} else {
				co.(*widget.Button).SetIcon(bundled.ResourceProtectedPng)
			}
			co.(*widget.Button).SetText(SharedResources[lii].FileName)

			// Func will trigger when tapped any shared resources
			co.(*widget.Button).OnTapped = func() {
				// creating dialog for actions
				var actionsDialog *dialog.CustomDialog
				var statusBtn *widget.Button

				if SharedResources[lii].FileStatus {
					statusBtn = widget.NewButtonWithIcon("Private", theme.ErrorIcon(), func() {
						SharedResources[lii].FileStatus = false
						randomizeChecksum()
						SharedList.Refresh()
						actionsDialog.Hide()
					})
				} else {
					statusBtn = widget.NewButtonWithIcon("Public", theme.ComputerIcon(), func() {
						SharedResources[lii].FileStatus = true
						randomizeChecksum()
						SharedList.Refresh()
						actionsDialog.Hide()
					})
				}

				//
				statusBtn.Importance = widget.HighImportance

				RemoveBtn := widget.NewButtonWithIcon("Remove", theme.DeleteIcon(), func() {
					if lii >= 0 && lii < len(SharedResources) {
						SharedResources = append(SharedResources[:lii], SharedResources[lii+1:]...)
					} else {
						log.Println("Index out of range")
					}
					randomizeChecksum()
					SharedList.Refresh()
					actionsDialog.Hide()
				})
				RemoveBtn.Importance = widget.DangerImportance
				actionsDialog = dialog.NewCustom(
					"Select following actions:",
					"Done",
					container.NewVBox(
						statusBtn,
						RemoveBtn,
					),
					w,
				)
				actionsDialog.Show()
			}
		},
	)

	// button to share resources
	browseBtn := widget.NewButtonWithIcon("Browse", theme.FileIcon(), func() {
		// need to be improve
		dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				return
			}
			if uc == nil {
				// User canceled the dialog
				fmt.Println("File open canceled by user.")
				return
			}
			if res, err := resource.NewResource(uc.URI().Path()); err != nil {
				log.Println("Can't share selected resource: " + uc.URI().Path())
			} else {
				SharedResources = append(SharedResources, *res)
				randomizeChecksum()
			}

			_ = uc.Close()
			SharedList.Refresh()
		}, w).Show()
	})
	browseBtn.Importance = widget.HighImportance

	SharedTab = container.NewBorder(
		nil,
		container.NewVBox(
			browseBtn,             // browse button to share files
			widget.NewSeparator(), // line separator above TriTabNav bar
		),
		nil,
		nil,
		SharedList, // shared resources list
	)
}
