package internal

import (
	"fmt"
	"html/template"
	"net/http"

	"01.alem.school/git/atastemi/forum/forum/domain"
)

func RenderMainPage(w http.ResponseWriter, r *http.Request, userSession *domain.Session, posts []domain.Posts) {
	tmpl, err := template.ParseFiles("./forum/templates/index.html", "./forum/templates/base.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if userSession != nil {
	}
	data := struct {
		Name  string
		Posts []domain.Posts
	}{
		Posts: posts,
	}
	if userSession == nil {
		data.Name = "Guest"
	} else {
		data.Name = userSession.Username
	}
	template.ParseGlob("template/*")

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RenderLoginPage(w http.ResponseWriter, r *http.Request, errorMessage string) {
	tmpl, err := template.ParseFiles("./forum/templates/login.html")
	if err != nil {
		fmt.Println("Cant get the HTML files")
		return
	}
	data := struct {
		ErrorMessage string
	}{
		ErrorMessage: errorMessage,
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RenderRegisterPage(w http.ResponseWriter, r *http.Request, errorMessage string) {
	tmpl, err := template.ParseFiles("./forum/templates/register.html")
	if err != nil {
		fmt.Println("Cant get the HTML files")
		return
	}
	data := struct {
		ErrorMessage string
	}{
		ErrorMessage: errorMessage,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RenderLikePages(w http.ResponseWriter, r *http.Request, username string, posts []domain.Posts) {
	tmpl, err := template.ParseFiles("./forum/templates/likedPosts.html", "./forum/templates/base.html")
	if err != nil {
		fmt.Println(err)
		fmt.Println("dada")

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Username string
		Posts    []domain.Posts
	}{
		Username: username,
		Posts:    posts,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("dada")

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RenderPostPage(w http.ResponseWriter, r *http.Request, username string, error string) {
	tmpl, err := template.ParseFiles("./forum/templates/createPost.html", "./forum/templates/base.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Username string
		Error    string
	}{
		Username: username,
		Error:    error,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RenderMyPostPage(w http.ResponseWriter, r *http.Request, username string, posts []domain.Posts) {
	tmpl, err := template.ParseFiles("./forum/templates/my_posts.html", "./forum/templates/base.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Username string
		Posts    []domain.Posts
	}{
		Username: username,
		Posts:    posts,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RenderErrorPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./forum/templates/404.html")
	if err != nil {
		fmt.Println("Cant get the HTML files")
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RenderAboutPage(w http.ResponseWriter, r *http.Request, userSession *domain.Session, posts domain.Posts, comments []domain.Comments) {
	tmpl, err := template.ParseFiles("./forum/templates/About.html", "./forum/templates/base.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Name     string
		UserId   int
		Post     domain.Posts
		Comments []domain.Comments
	}{
		Post:     posts,
		Comments: comments,
	}
	if userSession == nil {
		data.Name = "Guest"
		data.UserId = 0
	} else {
		data.Name = userSession.Username
		data.UserId = userSession.UserId
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
