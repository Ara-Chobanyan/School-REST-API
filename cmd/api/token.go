package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Ara-Chobanyan/placeholder/models"
	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
)

const SECRET = "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160"

type AccountPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credentials struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// adds a user to the database and hashes there password
func (app *application) createAccount(w http.ResponseWriter, r *http.Request) {
	var payload AccountPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	//hash the payload password
	payload.Password = app.hashPassword(payload.Password)

	var account models.Account

	account.Email = payload.Email
	account.Password = payload.Password

	err = app.models.InsertAccount(account)
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

// authenticates a user and sends them a jwt token
func (app *application) signIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJson(w, errors.New("unauthorized"))
		return
	}

	hashedPassword, err := app.models.GetAccount(creds.Email)
	if err != nil {
		app.errorJson(w, errors.New("unauthorized"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword.Password), []byte(creds.Password))
	if err != nil {
		log.Println(err)
		app.errorJson(w, errors.New("unauthorized password"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprintf(creds.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(SECRET))
	if err != nil {
		app.errorJson(w, errors.New("error signing"))
		return
	}

	err = app.writeJson(w, http.StatusOK, string(jwtBytes), "response")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}
