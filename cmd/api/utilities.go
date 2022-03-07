package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Use to write the data from the database into json
func (app *application) writeJson(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

// Help with json errors if any occur
func (app *application) errorJson(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	err = app.writeJson(w, statusCode, theError, "error")
	if err != nil {
		log.Println(err)
	}
}

// Used to log errors from json return
func (app *application) handleErrorJson(w http.ResponseWriter, err error) {
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

// Used to log errors from error var returns
func (app *application) logError(err error) {
	if err != nil {
		app.logger.Print(fmt.Errorf("Something went wrong %s", err))
	}
}
