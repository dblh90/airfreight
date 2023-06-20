package testutils

import (
	"database/sql"
	"fmt"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"testing"
)

func TestDB(t *testing.M) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
		postgres.WithQueriesFile("/var/project/migrations/schema.sql"),
	)
	container, _ := gnomock.Start(p)
	//t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)
	db, _ := sql.Open("postgres", connStr)
	db.Ping()
	// migrations has the required schema and data, and is ready to use
}
