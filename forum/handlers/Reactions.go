package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func (hh *HttpHandler) HandleLikeDislikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, err := hh.GetUsername(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	postIDStr := r.FormValue("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	action := r.FormValue("action")

	switch action {
	case "like":
		err = hh.business.LikePost(postID, session.UserId)
	case "dislike":
		err = hh.business.DislikePost(postID, session.UserId)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err != nil {
		hh.Handle404(w, r)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	redirectURL := r.Referer()
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func (hh *HttpHandler) HandleLikeDislikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, err := hh.GetUsername(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	commentIDStr := r.FormValue("comment_id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	action := r.FormValue("action")

	switch action {
	case "like":
		err = hh.business.LikeComment(commentID, session.UserId)
		if err != nil {
			fmt.Println(err)
		}
	case "dislike":
		err = hh.business.DislikeComment(commentID, session.UserId)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err != nil {
		hh.Handle404(w, r)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	redirectURL := r.Referer()
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
