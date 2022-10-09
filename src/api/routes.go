package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/home", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/listing/:id", app.getOneListing)
	router.HandlerFunc(http.MethodGet, "/v1/listing", app.getAllListings)

	return router
}
