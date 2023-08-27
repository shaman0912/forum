package handlers

import (
	"fmt"
	"net/http"

	"01.alem.school/git/atastemi/forum/forum/internal"
)

func (hh *HttpHandler) HandleLikedPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		sessionCookie, err := r.Cookie("session_id")
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		sessionID := sessionCookie.Value

		session, err := hh.business.Session(sessionID)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		fmt.Println(session.UserId)
		likedPosts, err := hh.business.GetLikedPosts(session.UserId)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		internal.RenderLikePages(w, r, session.Username, likedPosts)
	}
}
