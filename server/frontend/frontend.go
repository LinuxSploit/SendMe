package frontend

import (
	"html/template"
	"net/http"

	"github.com/LinuxSploit/SendMe/ui/shared"
)

func SharedFilePage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		"./template/index.html",
	))
	tmpl.Execute(w, shared.SharedResources)
}

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/welcome.html"))
	tmpl.Execute(w, nil)
}

func RequestPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/request.html"))
	tmpl.Execute(w, nil)
}
