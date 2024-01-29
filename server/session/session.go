package session

import (
	"errors"
	"net/http"

	"github.com/LinuxSploit/SendMe/internal/user"
	"github.com/LinuxSploit/SendMe/ui/home/homeTabs/actives"
)

func CheckSession(r *http.Request, w http.ResponseWriter) (user.User, error) {
	sessionToken, err := r.Cookie("token")
	if err != nil {
		return user.User{}, err
	}

	for _, auser := range actives.ActiveUsers {
		if auser.Token == sessionToken.Value {
			return auser, nil
		}
	}

	return user.User{}, errors.New("invalid session")
}
