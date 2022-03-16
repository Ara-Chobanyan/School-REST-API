package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ara-Chobanyan/placeholder/models"
)

type config struct {
	port int
	db   struct {
		dsn string
	}
}

// To avoid gloabl imports and use this as a method for functions
// that either interact with the database or need use of logger
type application struct {
	logger *log.Logger
	models models.DB
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 3000, "Server port to listen on")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://user:password@localhost/school?sslmode=disable", "Postgress connection")
	flag.Parse()

	// To log any potential errors
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Open the database to be able to access the data
	db, err := OpenDB(cfg)
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
		Addr:         fmt.Sprintf(":%d", cfg.port),
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

// OpenDB - Opens connection to the database
func OpenDB(cfg config) (*sql.DB, error) {

	// PSQL credentials fed into the driver to open up the database
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Context used to cancel data transfer if it takes longer then 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Used to make sure of connection to the db is alive
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
