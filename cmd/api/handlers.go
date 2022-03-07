package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getAStudent - Gets a student by id
func (app *application) getAStudent(w http.ResponseWriter, r *http.Request) {
	// gets the params from http request which is in the form of a id
	params := httprouter.ParamsFromContext(r.Context())

	// transform the id into a int from a string so it can be used in getting the correct col from the db
	id, err := strconv.Atoi(params.ByName("id"))
	app.logError(err)

	// Call the Get function which quires the id and returns a student data type
	student, err := app.models.Get(id)
	app.handleErrorJson(w, err)

	// Transform the student data type into a json
	err = app.writeJson(w, http.StatusOK, student, "student")
	app.handleErrorJson(w, err)
}

func (app *application) getClass(w http.ResponseWriter, r *http.Request) {

	// We will usually need to get params from here to find the classroom but for now just testing if the function and query will work

	class, err := app.models.GetAll()
	app.handleErrorJson(w, err)

	err = app.writeJson(w, http.StatusOK, class, "class")
	app.handleErrorJson(w, err)

}
