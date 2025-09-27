package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamicChain := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamicChain.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamicChain.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dynamicChain.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamicChain.ThenFunc(app.snippetCreatePost))

	mux.Handle("GET /user/signup", dynamicChain.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamicChain.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamicChain.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamicChain.ThenFunc(app.userLoginPost))
	mux.Handle("POST /user/logout", dynamicChain.ThenFunc(app.userLogoutPost))

	standardChain := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standardChain.Then(mux)
}
