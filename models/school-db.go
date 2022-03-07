package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Not too sure but I believe in a bigger project this should be a struct if the
// project would require multiple databases or differnt ports and config
// but I dont see a reason to add that type code cause it does not serve me in
//this project then again I am very inexpeinced so I can't be sure. Added this
//to be reminded to ask people and to just keep me thinking.
const (
	host     = "localhost"
	port     = "5432"
	user     = "ara"
	password = "arayik01"
	db       = "school"
)

type DB struct {
	DB *sql.DB
}

type student struct {
	ID         int     `json:"id"`
	First_Name string  `json:"first_name"`
	Last_Name  string  `json:"last_name"`
	Comments   string  `json:"comments"`
	Behvaior   string  `json:"behavior"`
	Grade      string  `json:"grade"`
	Average    float64 `json:"average"`
}

func NewDB(db *sql.DB) DB {
	return DB{
		DB: db,
	}
}

// OpenDB - Opens connection to the database
func OpenDB() (*sql.DB, error) {
	// Info used to connect to the PSQL database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable", host, port, user, password, db)

	// PSQL credentials fed into the driver to open up the database
	db, err := sql.Open("postgres", psqlInfo)
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


// Get - Quires for a matching id
func (s *DB) Get(id int) (*student, error) {
	// Declare a student data structure
	var student student

	// Use context to close the connection if it takes longer then 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// create a PSQL query to and have the id dynamic to the params
	query := `select id, first_name, last_name, comments, behavior, grade, average from test_ where id = $1`

	// Returning a row if any exist from the query
	row := s.DB.QueryRowContext(ctx, query, id)

	//Scan the row and copy the values that then returns a student data type which can be used to for the json function - writeJson
	err := row.Scan(
		&student.ID,
		&student.First_Name,
		&student.Last_Name,
		&student.Comments,
		&student.Behvaior,
		&student.Grade,
		&student.Average,
	)
	if err != nil {
		return nil, err
	}

	// return the student filled with the db data
	return &student, nil
}

func (s *DB) GetAll() ([]*student, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from test_`

	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*student
	for rows.Next() {
		var student student
		err := rows.Scan(
			&student.ID,
			&student.First_Name,
			&student.Last_Name,
			&student.Comments,
			&student.Behvaior,
			&student.Grade,
			&student.Average,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}
