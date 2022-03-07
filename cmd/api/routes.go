package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	//Better alternative to the default mux from net/http from what I read and heard
	router := httprouter.New()

	// Route to pick a single student by there id
	router.HandlerFunc(http.MethodGet, "/v1/student/:id", app.getAStudent)

	router.HandlerFunc(http.MethodGet, "/v1/class/", app.getClass)

	// Route to get all of the students in a class
	// router.HandlerFunc(http.MethodGet, "/v1/student/all", app.getAll)

	return router
}

/*
Have routes that connect to a class first then give the options to make edits or switch them around

/v1/:class/:id

*/
