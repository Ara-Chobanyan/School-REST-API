package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	//Better alternative to the default mux from net/http from what I read and heard
	router := httprouter.New()

	// Route to pick a single student by there id
	router.HandlerFunc(http.MethodGet, "/v1/class/id/:id", app.getAStudent)

	// Gets a student by name
	router.HandlerFunc(http.MethodGet, "/v1/class/name/:name", app.getAStudentByName)

	// Gets the entire classroom
	router.HandlerFunc(http.MethodGet, "/v1/class/", app.getClass)

	router.HandlerFunc(http.MethodPost, "/v1/class/add", app.editClass)

	router.HandlerFunc(http.MethodGet, "/v1/class/remove/:id", app.deleteStudent)

	return router
}
