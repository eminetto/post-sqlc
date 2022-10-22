package echo

import (
	"fmt"
	"net/http"

	"github.com/eminetto/post-sqlc/person"
	"github.com/labstack/echo/v4"
)

func Handlers(pService person.UseCase) *echo.Echo {
	e := echo.New()
	e.GET("/hello", Hello)
	e.GET("/hello/:lastname", GetUser(pService))
	return e
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetUser(s person.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		lastname := c.Param("lastname")
		people, err := s.Search(c.Request().Context(), lastname)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		if len(people) == 0 {
			return c.String(http.StatusNotFound, "not found")
		}
		return c.String(http.StatusOK, fmt.Sprintf("Hello %s %s", people[0].Name, people[0].LastName))
	}
}
