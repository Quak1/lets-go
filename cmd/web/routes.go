package main

import (
	"net/http"

	"github.com/Quak1/snippetbox/ui"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /ping", ping)

	dynamicChain := alice.New(app.sessionManager.LoadAndSave, app.noSurf, app.authenticate)

	mux.Handle("GET /{$}", dynamicChain.ThenFunc(app.home))
	mux.Handle("GET /about", dynamicChain.ThenFunc(app.aboutGet))
	mux.Handle("GET /snippet/view/{id}", dynamicChain.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", dynamicChain.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamicChain.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamicChain.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamicChain.ThenFunc(app.userLoginPost))

	protectedChain := dynamicChain.Append(app.requireAuthentication)

	mux.Handle("GET /account/view", protectedChain.ThenFunc(app.accountView))
	mux.Handle("GET /account/password/update", protectedChain.ThenFunc(app.accountPasswordUpdate))
	mux.Handle("POST /account/password/update", protectedChain.ThenFunc(app.accountPasswordUpdatePost))
	mux.Handle("GET /snippet/create", protectedChain.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protectedChain.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /user/logout", protectedChain.ThenFunc(app.userLogoutPost))

	standardChain := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standardChain.Then(mux)
}
