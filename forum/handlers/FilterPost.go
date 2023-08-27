package handlers

import (
	"net/http"

	"01.alem.school/git/atastemi/forum/forum/domain"
	"01.alem.school/git/atastemi/forum/forum/internal"
)

func (hh *HttpHandler) HandleFilteredPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		categories := r.URL.Query()["category"]

		var posts []domain.Posts
		var err error

		if len(categories) == 0 || (len(categories) == 1 && categories[0] == "none") {
			posts, err = hh.business.GetAllPosts()
		} else {
			posts, err = hh.business.GetPostsByCategories(categories)
		}

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		username, err := hh.GetUsername(w, r)
		if err != nil {
			internal.RenderMainPage(w, r, username, posts)
			return
		}
		if username.Username == "" {
			internal.RenderMainPage(w, r, username, posts)
			return
		}

		internal.RenderMainPage(w, r, username, posts)
	}
}
