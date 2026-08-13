package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"net/http"
)

const CookieName = "dokuroku_session"

type Manager struct {
	password string
	secret   string
}

func NewManager(password, secret string) Manager {
	return Manager{password: password, secret: secret}
}

func (m Manager) VerifyPassword(password string) bool {
	if m.password == "" {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(password), []byte(m.password)) == 1
}

func (m Manager) SessionValue() string {
	mac := hmac.New(sha256.New, []byte(m.secret))
	mac.Write([]byte("dokuroku"))
	return hex.EncodeToString(mac.Sum(nil))
}

func (m Manager) SetCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{Name: CookieName, Value: m.SessionValue(), Path: "/", HttpOnly: true, SameSite: http.SameSiteLaxMode})
}

func (m Manager) Authorized(r *http.Request) bool {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(cookie.Value), []byte(m.SessionValue())) == 1
}
