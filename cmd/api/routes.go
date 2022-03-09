package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), httprouter.ParamsKey, ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	//Better alternative to the default mux from net/http from what I read and heard
	router := httprouter.New()

	// Used to chain middle ware
	secure := alice.New(app.checkToken)

	// Route to pick a single student by there id
	router.HandlerFunc(http.MethodGet, "/v1/class/id/:id", app.getAStudent)

	// Gets a student by name
	router.HandlerFunc(http.MethodGet, "/v1/class/name/:name", app.getAStudentByName)

	// Gets the entire classroom
	router.HandlerFunc(http.MethodGet, "/v1/class/", app.getClass)

	// creates an account and hashes the password
	router.HandlerFunc(http.MethodPost, "/v1/admin/signup/", app.createAccount)

	// checks the creds and compares it to a hash password
	router.HandlerFunc(http.MethodPost, "/v1/admin/signin/", app.signIn)

	// Creates a new student if no id matches any of the existing ones or if it exists then it edits there data
	router.POST("/v1/admin/add", app.wrap(secure.ThenFunc(app.editClass)))

	// Removes a student by there id
	router.GET("/v1/admin/delete/:id", app.wrap(secure.ThenFunc(app.deleteStudent)))

	return app.enableCORS(router)
}
