package download

import (
	"fyne.io/fyne/v2/widget"
	"github.com/LinuxSploit/SendMe/internal/user"
)

type DownloadActivity struct {
	User           user.User
	FileName       string
	FilePath       string
	FileSize       float64
	DownloadedSize float64 // download progress percentage
	ProgressBar    *widget.ProgressBar
}
