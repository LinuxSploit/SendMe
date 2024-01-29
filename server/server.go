package server

import (
	"context"
	"errors"
	"html/template"
	"log"
	"net/http"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/LinuxSploit/SendMe/custom/switchBtn"
	"github.com/LinuxSploit/SendMe/server/api"
	"github.com/LinuxSploit/SendMe/server/frontend"
	"github.com/LinuxSploit/SendMe/server/session"
)

// Server represents an HTTP server.
type Server struct {
	server *http.Server
	mux    *http.ServeMux
	lock   sync.Mutex
}

// NewServer creates a new HTTP server with the specified address.
func NewServer(addr string) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", serverHandler)

	return &Server{
		server: &http.Server{Addr: addr, Handler: mux},
		mux:    mux,
	}
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)

	case "/api/login":
		api.LoginHandler(w, r)
	case "/api/active.json":
		// autherization check is made inside ActiveSharedFilesJSON
		api.ActiveSharedFilesJSON(w, r)
	case "/download":
		reqUser, err := session.CheckSession(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, "unautherized", http.StatusUnauthorized)
			return
		}
		api.DownloadFile(w, r, reqUser)

	case "/home":
		_, err := session.CheckSession(r, w)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/welcome", http.StatusTemporaryRedirect)
		}
		frontend.SharedFilePage(w, r)

	case "/welcome":
		_, err := session.CheckSession(r, w)
		if err == nil {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
		}
		frontend.WelcomePage(w, r)

	case "/request":
		frontend.RequestPage(w, r)
	case "/debug":
		api.Debug(w, r)
	default:
		tmpl := template.Must(template.ParseFiles("./template/404.html"))
		tmpl.Execute(w, nil)
	}
}

// Start initializes and starts the HTTP server.
func (s *Server) Start(switchbtn *switchBtn.SwitchButton, addrBtn *widget.Button, w fyne.Window) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.server == nil {
		return errors.New("Server is already not initialized")
	}

	addrBtn.SetText("http://" + GetServerLocalIP() + ":4556")
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Log the error if it's not a manual server shutdown or "port bind: address already" like errors
			log.Println(err)
			// turn off switch and hide address
			switchbtn.ToggleSwitchState()
			addrBtn.Hide()

			dialog.ShowError(err, w)
		}
	}()

	return nil
}

// Stop gracefully shuts down the HTTP server.
func (s *Server) Stop() error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.server == nil {
		return errors.New("Server is already stopped or not initialized")
	}

	// Shutdown the server gracefully using a context
	err := s.server.Shutdown(context.Background())

	// Set variables to nil to release resources
	s.server = nil
	s.mux = nil

	return err
}
