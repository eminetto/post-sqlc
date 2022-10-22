package person

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MysqlDBContainer struct {
	testcontainers.Container
	URI string
}

const (
	dbUser         = "post-sqlc"
	dbPassword     = "post-sqlc"
	database       = "post-sqlc"
	dbRootPassword = "db-root-password"
)

func SetupMysql(ctx context.Context) (*MysqlDBContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mariadb:10.9.3-jammy",
		ExposedPorts: []string{"3306/tcp"},
		WaitingFor:   wait.ForLog("Version: '10.9.3-MariaDB-1:10.9.3+maria~ubu2204'  socket: '/run/mysqld/mysqld.sock'  port: 3306  mariadb.org binary distribution"),
		Env: map[string]string{
			"MARIADB_USER":          dbUser,
			"MARIADB_PASSWORD":      dbPassword,
			"MARIADB_ROOT_PASSWORD": dbRootPassword,
			"MARIADB_DATABASE":      database,
		},
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	mappedPort, err := container.MappedPort(ctx, "3306")
	if err != nil {
		return nil, err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", dbRootPassword, hostIP, mappedPort.Port(), database)

	return &MysqlDBContainer{Container: container, URI: uri}, nil
}

func InitMySQL(ctx context.Context, db *sql.DB) error {
	file, err := os.Open("schema.sql")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		_, err := db.ExecContext(ctx, scanner.Text())
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func TruncateMySQL(ctx context.Context, db *sql.DB) error {
	query := []string{
		fmt.Sprintf("use %s;", database),
		"truncate table person",
	}
	for _, q := range query {
		_, err := db.ExecContext(ctx, q)
		if err != nil {
			return err
		}
	}
	return nil
}
