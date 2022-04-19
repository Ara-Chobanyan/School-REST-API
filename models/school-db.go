package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

type Student struct {
	ID         int     `json:"id"`
	First_Name string  `json:"first_name"`
	Last_Name  string  `json:"last_name"`
	Comments   string  `json:"comments"`
	Behavior   string  `json:"behavior"`
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

// Get - calls for a matching id
func (s *DB) GetById(id int) (*Student, error) {
	// Declare a student data structure
	var student Student

	// Use context to close the connection if it takes longer then 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Returning a row if any exist from the query
	row := s.DB.QueryRowContext(ctx, `SELECT * FROM mrsmith_class WHERE id = $1`, id)

	//Scan the row and copy the values that then returns a student data type which can be used to for the json function - writeJson
	err := row.Scan(
		&student.ID,
		&student.First_Name,
		&student.Last_Name,
		&student.Comments,
		&student.Behavior,
		&student.Grade,
		&student.Average,
	)
	if err != nil {
		log.Printf("Could not scan rows school-db.GetById: %v", err)
		return nil, err
	}

	// return the student filled with the db data
	return &student, nil
}

// Displays everything from the class table
func (s *DB) GetAll() ([]*Student, error) {
	// context to end connection if it takes longer then 3 seoncds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// query everything from test_
	query := `SELECT * from mrsmith_class`
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
			&student.Behavior,
			&student.Grade,
			&student.Average,
		)
		if err != nil {
			log.Printf("Could not scan rows at school-db.GetAll: %v", err)
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}

// Finds a student by there first name
func (s *DB) GetByName(name string) (*Student, error) {
	// Declare a student data structure
	var student Student

	// Use context to close the connection if it takes longer then 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// create a PSQL query to and have the id dynamic to the params
	query := `SELECT * from mrsmith_class where first_name=$1`

	// Returning a row if any exist from the query
	row := s.DB.QueryRowContext(ctx, query, name)

	//Scan the row and copy the values that then returns a student data type which can be used to for the json function - writeJson
	err := row.Scan(
		&student.ID,
		&student.First_Name,
		&student.Last_Name,
		&student.Comments,
		&student.Behavior,
		&student.Grade,
		&student.Average,
	)
	if err != nil {
		log.Printf("Could not scan rows at school-db.GetByName: %v", err)
		return nil, err
	}

	// return the student filled with the db data
	return &student, nil
}

// InsertAStudent - Inserts a new student into the class table
func (s *DB) InsertAStudent(student Student) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO mrsmith_class (first_name, last_name, comments, behavior, grade, average) VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := s.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Could not execute query at school-db.InsertAStudent: %V", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		student.First_Name,
		student.Last_Name,
		student.Comments,
		student.Behavior,
		student.Grade,
		student.Average,
	)

	if err != nil {
		log.Printf("Could not add a new student at school-db.InsertAStudent: %v", err)
		return err
	}

	return nil
}

// Updates a student from the database
func (s *DB) UpdateStudent(student Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE mrsmith_class SET first_name = $1, last_name = $2, comments = $3,
            behavior = $4, grade = $5, average = $6 WHERE id = $7`

	stmt, err := s.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = s.DB.ExecContext(ctx, query,
		student.First_Name,
		student.Last_Name,
		student.Comments,
		student.Behavior,
		student.Grade,
		student.Average,
		student.ID,
	)

	if err != nil {
		log.Printf("Could not update student at school-db.UpdateStudent: %v", err)
		return err
	}

	return nil
}

// Deletes a student from the database
func (s *DB) DeleteStudent(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE from mrsmith_class WHERE id=$1`

	stmt, err := s.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		log.Printf("Could not remove a student at school-db.DeleteStudent: %v", err)
		return err
	}
	return nil
}

// Creates a new user account
func (s *DB) InsertAccount(user Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO accounts (email, password) values($1,$2)`

	_, err := s.DB.ExecContext(ctx, query,
		user.Email,
		user.Password,
	)
	if err != nil {
		log.Printf("Could not create a new account at school-db.InsertAStudent: %v", err)
		return err
	}

	return nil
}

// Returns a Account from the DB to json so it can be checked to see if the client credentials are correct
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
		log.Printf("Could not get a user account at school-db.GetAccount: %v", err)
		return nil, err
	}

	return &user, nil
}
