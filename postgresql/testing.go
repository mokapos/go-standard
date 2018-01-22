package postgresql

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/mokapos/go-standard/testingutil"
)

const (
	dbHost     = "localhost"
	dbPort     = "5432"
	dbUserName = "postgres"
	dbPassword = ""
	dbName     = "library_test"
)

// NewTestDBConnection creates a connection to testing database.
func NewTestDBConnection(t *testing.T) *DB {
	db, err := New(&Options{
		Host:     dbHost,
		Port:     dbPort,
		Name:     dbName,
		UserName: dbUserName,
		Password: dbPassword,
		SSLMode:  "disable",
	}, nil)
	if err != nil {
		// On fatal exception during testing, call Fatal from testing package to stop the execution
		// Otherwise, call asserts function from testingutil.go
		t.Fatalf("Fail to connect database. %s", err.Error())
	}
	return db
}

// RunSQLFile runs a SQL file inside testdata folder.
func RunSQLFile(t *testing.T, db *DB, name string) {
	content, err := ioutil.ReadFile(fmt.Sprintf("./testdata/%s.sql", name))
	testingutil.Ok(t, err)
	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {
			_, err := db.Exec(query)
			testingutil.Ok(t, err)
		}
	}
}
