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

	// Creates a new student if no id matches any of the existing ones or if it exists then it edits there data
	router.HandlerFunc(http.MethodPost, "/v1/admin/add", app.editClass)

	// Removes a student by there id
	router.HandlerFunc(http.MethodGet, "/v1/admin/remove/:id", app.deleteStudent)

	// creates an account and hashes the password
	router.HandlerFunc(http.MethodPost, "/v1/admin/account/", app.createAccount)

	//test route
	router.HandlerFunc(http.MethodPost, "/v1/test/account/signin", app.signIn)

	// struggling too much on it will get back to it
	// router.HandlerFunc(http.MethodGet, "/v1/admin/class/:name", app.createClass)

	return router
}
