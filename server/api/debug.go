package api

import (
	"encoding/json"
	"net/http"

	"github.com/LinuxSploit/SendMe/ui/activity"
)

func Debug(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	for _, act := range activity.Activities {
		actJson, err := json.Marshal(act)
		if err != nil {
			break
		}
		w.Write(actJson)
	}
}
