package handlers

// var (
// 	githubOAuthConfig = oauth2.Config{
// 		ClientID:     "f4aefacb16ef0a806638",
// 		ClientSecret: "58860bde90fe36e906cc853d983c029a2c18ce64",
// 		Endpoint:     github.Endpoint,
// 		RedirectURL:  "http://localhost:8080/login_git/callback",
// 		Scopes:       []string{"user:email"},
// 	}

// 	oauthStateString = "random"
// )

// func (hh *HttpHandler) GitHubLoginHandler(w http.ResponseWriter, r *http.Request) {
// 	// Create the dynamic redirect URL for GitHub login
// 	redirectURL := fmt.Sprintf(
// 		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
// 		githubOAuthConfig.ClientID,
// 		"http://localhost:8080/login_git/callback",
// 	)

// 	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
// }

// func (hh *HttpHandler) GitHubCallbackHandler(w http.ResponseWriter, r *http.Request) {
// 	code := r.URL.Query().Get("code")

// 	githubAccessToken := hh.GetGithubAccessToken(code)

// 	githubData := hh.GetGithubData(githubAccessToken)

// 	// Here, you can use the GitHub data to perform login or registration logic.
// 	// You can create your own functions to handle registration or login based on the GitHub data.

// 	// For example:
// 	// err := hh.PerformGitHubLoginOrRegister(githubData)
// 	// if err != nil {
// 	//     http.Error(w, "Failed to login or register", http.StatusInternalServerError)
// 	//     return
// 	// }

// 	// Redirect to the main page after successful login or registration.
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }
