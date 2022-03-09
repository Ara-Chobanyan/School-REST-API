package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
	user     = "name"
	password = "password"
	db       = "school"
)

type DB struct {
	DB *sql.DB
}

type Student struct {
	ID         int     `json:"id"`
	First_Name string  `json:"first_name"`
	Last_Name  string  `json:"last_name"`
	Comments   string  `json:"comments"`
	Behvaior   string  `json:"behavior"`
	Grade      string  `json:"grade"`
	Average    float64 `json:"average"`
}

type Account struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
func (s *DB) GetById(id int) (*Student, error) {
	// Declare a student data structure
	var student Student

	// Use context to close the connection if it takes longer then 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// create a PSQL query to and have the id dynamic to the params
	query := `select * from mrsmith_class where id = $1`

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

func (s *DB) GetAll() ([]*Student, error) {
	// context to end connection if it takes longer then 3 seoncds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// query everything from test_
	query := `select * from mrsmith_class`
	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Need a json array to get all of the students so loop through the columns and rows and appends it into a []*student
	var students []*Student
	for rows.Next() {
		var student Student
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

func (s *DB) GetByName(name string) (*Student, error) {
	// Declare a student data structure
	var student Student

	// Use context to close the connection if it takes longer then 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// create a PSQL query to and have the id dynamic to the params
	query := `select * from mrsmith_class where first_name=$1`

	// Returning a row if any exist from the query
	row := s.DB.QueryRowContext(ctx, query, name)

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

// InsertAStudent - Inserts a new student into the class table
func (s *DB) InsertAStudent(student Student) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into mrsmith_class (first_name, last_name,comments,behavior,grade,average) 
            values($1,$2,$3,$4,$5,$6)`

	_, err := s.DB.ExecContext(ctx, query,
		student.First_Name,
		student.Last_Name,
		student.Comments,
		student.Behvaior,
		student.Grade,
		student.Average,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *DB) UpdateStudent(student Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update mrsmith_class set first_name = $1, last_name = $2, comments = $3,
            behavior = $4, grade = $5, average = $6 where id = $7`

	_, err := s.DB.ExecContext(ctx, query,
		student.First_Name,
		student.Last_Name,
		student.Comments,
		student.Behvaior,
		student.Grade,
		student.Average,
		student.ID,
	)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (s *DB) DeleteStudent(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from mrsmith_class where id=$1`

	_, err := s.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *DB) InsertAccount(user Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into accounts (email, password) values($1,$2)`

	_, err := s.DB.ExecContext(ctx, query,
		user.Email,
		user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *DB) GetAccount(email string) (*Account, error) {
	var user Account

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from accounts where email=$1`

	row := s.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

