package person_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eminetto/post-sqlc/person"
	"github.com/eminetto/post-sqlc/person/db"
	"github.com/stretchr/testify/assert"
)

func TestService_Get(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer d.Close()
	queries := db.New(d)
	t.Run("usuário encontrado", func(t *testing.T) {
		// fase: Arrange
		rows := sqlmock.NewRows([]string{"id", "name", "lastname", "created_at", "updated_at"}).
			AddRow(1, "Ozzy", "Osbourne", time.Now(), time.Now())
		mock.ExpectQuery("[A-Za-z]?select id, first_name, last_name, created_at, updated_at from person where id").
			WillReturnRows(rows)

		service := person.NewService(queries)
		// fase: Act
		found, err := service.Get(context.TODO(), person.ID(1))

		// fase: Assert
		p := &person.Person{
			ID:       1,
			Name:     "Ozzy",
			LastName: "Osbourne",
		}
		assert.Nil(t, err)
		assert.Equal(t, p, found)
	})
	t.Run("usuário não encontrado", func(t *testing.T) {
		mock.ExpectQuery("[A-Za-z]?select id, first_name, last_name, created_at, updated_at from person where id").WillReturnError(errors.New(""))
		service := person.NewService(queries)
		found, err := service.Get(context.TODO(), person.ID(1))
		assert.Nil(t, found)
		assert.Errorf(t, err, "erro lendo person do repositório: %w")
	})
}

func TestService_Search(t *testing.T) {
	d, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer d.Close()
	queries := db.New(d)

	p1 := &person.Person{
		ID:       1,
		Name:     "Ozzy",
		LastName: "Osbourne",
	}
	p2 := &person.Person{
		ID:       2,
		Name:     "Ronnie",
		LastName: "Dio",
	}

	tests := []struct {
		query       string
		result      *sqlmock.Rows
		expectedErr error
		expected    []*person.Person
	}{
		{
			query:       "ozzy",
			result:      sqlmock.NewRows([]string{"id", "name", "lastname"}).AddRow(p1.ID, p1.Name, p1.LastName),
			expectedErr: nil,
			expected:    []*person.Person{p1},
		},
		{
			query:       "Ozzy",
			result:      sqlmock.NewRows([]string{"id", "name", "lastname"}).AddRow(p1.ID, p1.Name, p1.LastName),
			expectedErr: nil,
			expected:    []*person.Person{p1},
		},
		{
			query:       "osbourne",
			result:      sqlmock.NewRows([]string{"id", "name", "lastname"}).AddRow(p1.ID, p1.Name, p1.LastName),
			expectedErr: nil,
			expected:    []*person.Person{p1},
		},
		{
			query:       "Osbourne",
			result:      sqlmock.NewRows([]string{"id", "name", "lastname"}).AddRow(p1.ID, p1.Name, p1.LastName),
			expectedErr: nil,
			expected:    []*person.Person{p1},
		},
		{
			query:       "Dio",
			result:      sqlmock.NewRows([]string{"id", "name", "lastname"}).AddRow(p2.ID, p2.Name, p2.LastName),
			expectedErr: nil,
			expected:    []*person.Person{p2},
		},
		{
			query:       "dio",
			result:      sqlmock.NewRows([]string{"id", "name", "lastname"}).AddRow(p2.ID, p2.Name, p2.LastName),
			expectedErr: nil,
			expected:    []*person.Person{p2},
		},
		{
			query:       "ronnie",
			result:      sqlmock.NewRows([]string{"id", "name", "lastname"}).AddRow(p2.ID, p2.Name, p2.LastName),
			expectedErr: nil,
			expected:    []*person.Person{p2},
		},
		{
			query:       "Ronnie",
			result:      sqlmock.NewRows([]string{"id", "name", "lastname"}).AddRow(p2.ID, p2.Name, p2.LastName),
			expectedErr: nil,
			expected:    []*person.Person{p2},
		},
		{
			query:       "Tony",
			result:      nil,
			expectedErr: fmt.Errorf("error searching from database: %w", fmt.Errorf("not found")),
			expected:    nil,
		},
		{
			query:       "martin",
			result:      nil,
			expectedErr: fmt.Errorf("error searching from database: %w", fmt.Errorf("not found")),
			expected:    nil,
		},
	}
	for _, test := range tests {
		q := "-- name: Search :many select id, first_name, last_name from person where first_name like ? or last_name like ?"

		if test.expectedErr != nil {
			mock.ExpectQuery(q).
				WithArgs(test.query, test.query).
				WillReturnError(errors.New("not found"))
		}
		if test.result != nil {
			mock.ExpectQuery(q).
				WithArgs(test.query, test.query).
				WillReturnRows(test.result)
		}
		service := person.NewService(queries)
		found, err := service.Search(context.TODO(), test.query)
		assert.Equal(t, test.expectedErr, err)
		assert.Equal(t, test.expected, found)
	}
}

func TestCreate(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer d.Close()
	firstName := "Ozzy"
	lastName := "Osbourne"
	// mock.ExpectBegin()
	mock.ExpectExec("[A-Za-z]?insert into person").
		WithArgs(firstName, lastName).
		WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectCommit()
	queries := db.New(d)
	service := person.NewService(queries)
	id, err := service.Create(context.TODO(), firstName, lastName)
	assert.Nil(t, err)
	assert.Equal(t, person.ID(1), id)
}
