package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ara-Chobanyan/placeholder/models"
)

const port = 4000

// To avoid gloabl imports and use this as a method for functions
// that either interact with the database or need use of logger
type application struct {
	logger *log.Logger
	models models.DB
}

func main() {
	// To log any potential errors
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Open the database to be able to access the data
	db, err := models.OpenDB()
	if err != nil {
		logger.Fatal(err)
	}
	// make sure to close the database to avoid any leakage
	defer db.Close()

	app := &application{
		logger: logger,
		models: models.NewDB(db),
	}

	// Server inputs
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      app.routes(),
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Running server")

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
