package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/LinuxSploit/SendMe/internal/download"
	"github.com/LinuxSploit/SendMe/internal/resource"
	"github.com/LinuxSploit/SendMe/internal/user"
	"github.com/LinuxSploit/SendMe/ui/activity"
	"github.com/LinuxSploit/SendMe/ui/home/homeTabs/actives"
	"github.com/LinuxSploit/SendMe/ui/shared"
)

type DataTaxi struct {
	Files    []resource.Resource `json:"Files"`
	Checksum string              `json:"Checksum"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := r.FormValue("username")
	if strings.Trim(username, " ") == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// setting session
	tmp := user.NewUser(username, r)
	actives.ActiveUsers = append(actives.ActiveUsers, tmp)

	// Create a new cookie
	cookie := http.Cookie{
		Name:     "token",
		Value:    tmp.Token,
		Expires:  time.Now().Add(24 * time.Hour), // Cookie expires in 24 hours
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}

	// Set the cookie in the response
	http.SetCookie(w, &cookie)

	actives.ActiveList.Refresh()
}

// ActiveSharedFilesJSON marks user as active and update SharedFiles
func ActiveSharedFilesJSON(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		token, err := r.Cookie("token")
		if err != nil {
			return
		}

		r.ParseForm()
		checksum := r.FormValue("checksum")

		// needs to be optimize
		userFound := false
		for i, user := range actives.ActiveUsers {
			if user.Token == token.Value {
				// setting found user status "Online"
				actives.ActiveUsers[i].LastSeen = time.Now()
				actives.ActiveUsers[i].Online = true
				actives.ActiveList.Refresh()
				userFound = true
			}
		}
		if !userFound {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		if checksum == "" {
			writeSharedFiles(w)
			return
		}

		if checksum != shared.CurrentResourcesChecksum {
			writeSharedFiles(w)
			return
		}
	}

}

func writeSharedFiles(w http.ResponseWriter) {

	jsonBytes, err := json.Marshal(
		DataTaxi{
			shared.SharedResources,
			shared.CurrentResourcesChecksum,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
}

func DownloadFile(w http.ResponseWriter, r *http.Request, reqUser user.User) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	filename := r.FormValue("filename")
	if filename == "" {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	// Iterate through shared files
	for _, afile := range shared.SharedResources {
		// Check if afile is the requested file
		if afile.FileName == filename {
			// Download afile if AllowedUser is empty
			if len(afile.FileAllowedUsers) == 0 {
				downloadActivity := &download.DownloadActivity{
					User:           reqUser,
					FileName:       afile.FileName,
					FilePath:       afile.FilePath,
					FileSize:       float64(afile.FileSize),
					DownloadedSize: 0,
					ProgressBar:    widget.NewProgressBar(),
				}
				activity.Activities = append(activity.Activities, downloadActivity)
				activity.ActivitiesList.Refresh()
				afile.Download(w, downloadActivity)
				return
			}

			// Check if the requesting user is in the list of allowed users
			for _, auser := range afile.FileAllowedUsers {
				if reqUser.Token == auser.Token {
					return
				}
			}

			// If the loop completes without finding the user, they are not allowed
			http.Error(w, "You're not allowed to download this file, please request access!", http.StatusForbidden)
			return
		}
	}

	// If the loop completes without finding the file, it doesn't exist
	http.Error(w, "File not found", http.StatusNotFound)
}
