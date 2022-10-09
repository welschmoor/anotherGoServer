package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/welschmoor/anotherGoServer/models"
)

func (app *application) getOneListing(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context()) 

	id, err := strconv.Atoi(params.ByName("id")) //convert id to int ?
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id: ", id)

	listing := models.Listing{
		ID:          id,
		Title:       "Fahrrad",
		Description: "100% nicht geklaut",
		Price:       15000,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, listing, "movie")
	if err != nil {
		app.logger.Println(errors.New("writeJSON throws error"))
	}
}

func (app *application) getAllListings(w http.ResponseWriter, r *http.Request) {

}
