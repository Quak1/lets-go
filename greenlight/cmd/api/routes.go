package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.GET("/v1/healthcheck", app.healthcheckHandler)
	router.POST("/v1/movies", app.createMovieHandler)
	router.GET("/v1/movies/:id", app.showMovieHandler)

	return router
}
