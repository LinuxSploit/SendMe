package resource

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/LinuxSploit/SendMe/internal/download"
	"github.com/LinuxSploit/SendMe/internal/user"
)

type Resource struct {
	FileName         string      `json:"FileName"`
	FilePath         string      `json:"FilePath"`
	FileSize         int64       `json:"FileSize"`
	FileStatus       bool        `json:"FileStatus"` // public/private
	FileAllowedUsers []user.User `json:"-"`
}

func NewResource(filepath string) (*Resource, error) {
	info, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	return &Resource{
		FileName:   info.Name(),
		FilePath:   filepath,
		FileSize:   info.Size(),
		FileStatus: true,
	}, nil
}

func (res *Resource) Download(w http.ResponseWriter, downActivity *download.DownloadActivity) {
	openedFile, err := os.Open(res.FilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error opening file: %v", err)
		return
	}
	defer openedFile.Close()

	buffer := make([]byte, 9024)
	for {
		n, err := openedFile.Read(buffer)
		if err == io.EOF || n <= 0 {
			break
		}
		if err != nil {
			log.Printf("Error reading file: %v", err)
			break
		}

		w.Write(buffer[:n])
		downActivity.DownloadedSize += float64(n)
		downActivity.ProgressBar.SetValue(downActivity.DownloadedSize / float64(res.FileSize))
		downActivity.ProgressBar.Refresh()
		log.Printf("Downloaded: %v / %v (%.2f%%)", downActivity.DownloadedSize, res.FileSize, (downActivity.DownloadedSize/float64(res.FileSize))*100)
	}
}
