package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Ara-Chobanyan/placeholder/models"
	"github.com/julienschmidt/httprouter"
)

type StudentPayload struct {
	ID         string `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Comments   string `json:"comments"`
	Behavior   string `json:"behavior"`
	Grade      string `json:"grade"`
	Average    string `json:"average"`
}

type jsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// getAStudent - Gets a student by id
func (app *application) getAStudent(w http.ResponseWriter, r *http.Request) {
	// gets the params from http request which is in the form of a id
	params := httprouter.ParamsFromContext(r.Context())

	// transform the id into a int from a string so it can be used in getting the correct col from the db
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	// Call the Get function which quires the id and returns a student data type
	student, err := app.models.GetById(id)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	// Transform the student data type into a json
	err = app.writeJson(w, http.StatusOK, student, "student")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

// Gets the entire class
func (app *application) getClass(w http.ResponseWriter, r *http.Request) {

	// calls GetAll which just displays everything in the table into json
	class, err := app.models.GetAll()
	if err != nil {
		app.errorJson(w, err)
		return
	}

	// write back json
	err = app.writeJson(w, http.StatusOK, class, "classname")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

// getAStudentByName - Finds a student by there name
func (app *application) getAStudentByName(w http.ResponseWriter, r *http.Request) {
	// gets the params
	params := httprouter.ParamsFromContext(r.Context())

	name := params.ByName("name")

	student, err := app.models.GetByName(name)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, student, "student")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

// Either adds a new student or edits a students data
func (app *application) editClass(w http.ResponseWriter, r *http.Request) {
	// Init the payload var
	var payload StudentPayload

	//Get the client payload and decode into the payload data type
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	// Declare a student data type
	var student models.Student

	//Check if the payload student data already exists
	if payload.ID != "0" {
		//parse the string id into a int
		id, err := strconv.Atoi(payload.ID)
		if err != nil {
			log.Println(err)
		}

		//Get the student id if it already exists and dereference it to the student type so the edits can be correct
		s, err := app.models.GetById(id)
		if err != nil {
			log.Println(err)
		}
		student = *s
	}

	//Insert the paylaod to the student data type
	student.ID, _ = strconv.Atoi(payload.ID)
	student.First_Name = payload.First_Name
	student.Last_Name = payload.Last_Name
	student.Comments = payload.Comments
	student.Behavior = payload.Behavior
	student.Grade = payload.Grade
	student.Average, _ = strconv.ParseFloat(payload.Average, 64)

	// To check if its for a new student or to update a student
	if student.ID == 0 {
		err = app.models.InsertAStudent(student)
		if err != nil {
			app.errorJson(w, err)
			return
		}
	} else {
		err = app.models.UpdateStudent(student)
		if err != nil {
			app.errorJson(w, err)
			return
		}
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJson(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	log.Println(id)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.models.DeleteStudent(id)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJson(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}
