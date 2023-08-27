package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"01.alem.school/git/atastemi/forum/forum/domain"
	"01.alem.school/git/atastemi/forum/forum/internal"
)

func (hh *HttpHandler) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		sessionId, err := hh.business.Login(username, password)
		if err != nil {
			if errors.Is(err, domain.ErrInvalidUser) {
				internal.RenderLoginPage(w, r, "Invalid username or password")
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		expiry := time.Now().Add(24 * time.Hour)

		cookie := http.Cookie{
			Name:     "session_id",
			Value:    sessionId.String(),
			Expires:  expiry,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		internal.RenderLoginPage(w, r, "")
	}
}

func (hh *HttpHandler) HandleUserRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		username := strings.TrimSpace(r.FormValue("username"))
		email := strings.TrimSpace(r.FormValue("email"))
		password := r.FormValue("password")

		err := hh.business.Registration(username, password, email)
		if err != nil {
			if errors.Is(err, domain.ErrInvalidDataonRegistartion) {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				internal.RenderRegisterPage(w, r, "Invalid username or password")
				return
			}
			if err == domain.ErrUserAlreadyExist {
				internal.RenderRegisterPage(w, r, "This email or username already exist")
				return
			}
			log.Printf("Error with registration:%s", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		internal.RenderRegisterPage(w, r, "")
	}
}

func (hh *HttpHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	ClearSession(w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ClearSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	})
}
