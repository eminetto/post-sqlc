package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eminetto/post-sqlc/internal/api"
	"github.com/eminetto/post-sqlc/internal/http/echo"
	"github.com/eminetto/post-sqlc/person"
	"github.com/eminetto/post-sqlc/person/db"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser         = "post-sqlc"
	dbPassword     = "post-sqlc"
	database       = "post-sqlc"
	dbRootPassword = "db-root-password"
)

func main() {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, "localhost", "3306", database)
	d, err := sql.Open("mysql", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(d)
	pService := person.NewService(queries)
	h := echo.Handlers(pService)
	err = api.Start("8000", h)
	if err != nil {
		log.Fatal("error running api", err)
	}
}
