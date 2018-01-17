package postgresql

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/mokapos/go-standard/testingutil"
)

const (
	dbHost     = "localhost"
	dbPort     = "5432"
	dbUser     = "postgres"
	dbPassword = ""
	dbName     = "library"
)

// CreateTestDB creates a database for testing
// Also, it returns the database object, database name, and a function to delete that database
func CreateTestDB(t *testing.T) (*DB, func()) {
	db, err := New(&Options{
		Host:     dbHost,
		Port:     dbPort,
		Name:     dbName,
		UserName: dbUser,
		Password: dbPassword,
		SSLMode:  "disable",
	}, nil)
	if err != nil {
		// On fatal exception during testing, call Fatal from testing package to stop the execution
		// Otherwise, call asserts function from testingutil.go
		t.Fatalf("Fail to connect database. %s", err.Error())
	}

	rand.Seed(time.Now().UnixNano())
	testDBName := "test" + strconv.FormatInt(rand.Int63(), 10)

	_, err = db.Exec("CREATE DATABASE " + testDBName)
	if err != nil {
		t.Fatalf("Fail to create database %s. %s", testDBName, err.Error())
	}

	testDB, err := New(&Options{
		Host:     dbHost,
		Port:     dbPort,
		Name:     testDBName,
		UserName: dbUser,
		Password: dbPassword,
		SSLMode:  "disable",
	}, nil)
	db.Exec("ALTER DATABASE " + testDBName + " CONNECTION LIMIT 1;")
	if err != nil {
		t.Fatalf("Fail to connect database. %s", err.Error())
	}

	return testDB, func() {
		_, err := db.Exec("DROP DATABASE " + testDBName)
		if err != nil {
			t.Fatalf("Fail to drop database %s. %s", testDBName, err.Error())
		}
	}
}

// SetSchema sets the schema of a test database using specified SQL file
func SetSchema(t *testing.T, db *DB, SQLFile string) {
	content, err := ioutil.ReadFile(fmt.Sprintf("./testdata/%s.sql", SQLFile))
	testingutil.Ok(t, err)

	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {
			_, err := db.Exec(query)
			testingutil.Ok(t, err)
		}
	}
}
