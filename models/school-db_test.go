package models

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var u = &Student{
	ID:         int(uuid.New().ID()),
	First_Name: "jon",
	Last_Name:  "smith",
	Comments:   "nice",
	Behavior:   "good",
	Grade:      "A",
	Average:    94,
}

var b = Student{
	ID:         0,
	First_Name: "jon",
	Last_Name:  "smith",
	Comments:   "nice",
	Behavior:   "good",
	Grade:      "A",
	Average:    94,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
	repo := &DB{DB: db}
	defer func() {
		repo.DB.Close()
	}()

	query := `SELECT \* FROM mrsmith_class WHERE id \= \$1`

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "comments", "behavior", "grade", "average"}).AddRow(u.ID, u.First_Name, u.Last_Name, u.Comments, u.Behavior, u.Grade, u.Average)

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	user, err := repo.GetById(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetAll(t *testing.T) {
	db, mock := NewMock()
	repo := &DB{DB: db}
	defer func() {
		repo.DB.Close()
	}()

	query := `SELECT \* from mrsmith_class `

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "comments", "behavior", "grade", "average"}).AddRow(1, "jon", "smith", "nice", "good", "b", "80").AddRow(2, "bob", "wayne", "bad", "rude", "f", "0")

	mock.ExpectQuery(query).WillReturnRows(rows)

	user, err := repo.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetByName(t *testing.T) {
	db, mock := NewMock()
	repo := &DB{DB: db}
	defer func() {
		repo.DB.Close()
	}()

	query := `SELECT \* from mrsmith_class where first_name\=\$1`

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "comments", "behavior", "grade", "average"}).AddRow(u.ID, u.First_Name, u.Last_Name, u.Comments, u.Behavior, u.Grade, u.Average)

	mock.ExpectQuery(query).WithArgs(u.First_Name).WillReturnRows(rows)

	user, err := repo.GetByName(u.First_Name)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestDeleteStudent(t *testing.T) {
	db, mock := NewMock()
	repo := &DB{DB: db}
	defer func() {
		repo.DB.Close()
	}()

	stmt := `DELETE from mrsmith_class WHERE id\=\$1`

	prep := mock.ExpectPrepare(stmt)
	prep.ExpectExec().WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteStudent(u.ID)
	assert.NoError(t, err)

}

func TestInsertAStudent(t *testing.T) {
	db, mock := NewMock()
	repo := &DB{DB: db}
	defer func() {
		repo.DB.Close()
	}()

	query := `INSERT INTO mrsmith_class \(first_name, last_name, comments, behavior, grade, average\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(b.First_Name, b.Last_Name, b.Comments, b.Behavior, b.Grade, b.Average).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.InsertAStudent(b)
	assert.NoError(t, err)
}

func TestUpdateStudent(t *testing.T) {
	db, mock := NewMock()
	repo := &DB{DB: db}
	defer func() {
		repo.DB.Close()
	}()

	query := `UPDATE mrsmith_class SET first_name \= \$1, last_name \= \$2, comments \= \$3, behavior \= \$4, grade \= \$5, average \= \$6  WHERE id \= \$7`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(b.First_Name, b.Last_Name, b.Comments, b.Behavior, b.Grade, b.Average, b.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateStudent(b)
	assert.NoError(t, err)

}
