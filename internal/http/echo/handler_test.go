package echo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eminetto/post-sqlc/internal/http/echo"
	"github.com/eminetto/post-sqlc/person"
	person_mock "github.com/eminetto/post-sqlc/person/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHello(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := echo.Handlers(nil).NewContext(req, rec)
	c.SetPath("/hello")
	err := echo.Hello(c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())
}

func TestGetUser(t *testing.T) {
	t.Run("status ok", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		p := []*person.Person{
			{
				ID:       1,
				Name:     "Ronnie",
				LastName: "Dio",
			},
		}
		s := person_mock.NewUseCase(t)
		s.On("Search", mock.Anything, "dio").
			Return(p, nil).
			Once()
		c := echo.Handlers(nil).NewContext(req, rec)
		c.SetPath("/hello/:lastname")
		c.SetParamNames("lastname")
		c.SetParamValues("dio")
		h := echo.GetUser(s)
		err := h(c)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello Ronnie Dio", rec.Body.String())
	})
	t.Run("status not found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		s := person_mock.NewUseCase(t)
		s.On("Search", mock.Anything, "dio").
			Return([]*person.Person{}, nil).
			Once()
		c := echo.Handlers(nil).NewContext(req, rec)
		c.SetPath("/hello/:lastname")
		c.SetParamNames("lastname")
		c.SetParamValues("dio")
		h := echo.GetUser(s)
		err := h(c)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
