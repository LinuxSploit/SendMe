package user

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/mileusna/useragent"
)

type User struct {
	Token    string
	Username string
	IP       net.IP
	Browser  string
	Device   string
	Online   bool

	LastSeen time.Time
}

func NewUser(username string, r *http.Request) User {

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// Handle the error, for example, log it
		fmt.Println("Error splitting host and port:", err)
		// Set a default IP or take alternative action
		ip = "127.0.0.1"
	}

	user := User{
		Token:    generateRandomHash(),
		Username: username,
		IP:       net.ParseIP(ip),
		Online:   true, // Set a default status or retrieve it from the request
		LastSeen: time.Now(),
	}

	user.UpdateFromUserAgent(r.UserAgent())

	return user
}

func (user *User) UpdateFromUserAgent(uaStr string) {
	// Implement logic to extract Browser and Device info from userAgent string
	ua := useragent.Parse(uaStr)
	// Update the corresponding fields in the User struct
	user.Browser = fmt.Sprintf("%s %s", ua.Name, ua.Version)
	user.Device = fmt.Sprintf("%s %s", ua.OS, ua.OSVersion)
}

func (user User) GetUserStatusWithColor() (string, color.Color) {
	if user.Online {
		return "Online", color.RGBA{0, 200, 50, 255}
	} else {
		return "Offline", color.RGBA{255, 0, 50, 255}
	}
}

func generateRandomHash() string {
	// Generate 16 random bytes (128 bits)
	randomBytes := securecookie.GenerateRandomKey(64)

	// Convert the random bytes to a hexadecimal string
	hashString := hex.EncodeToString(randomBytes)

	return hashString
}
