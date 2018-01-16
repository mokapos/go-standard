package postgresql

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgreSQL driver for default package 'database/sql' or the derivatives
)

// DB connection struct, extending sqlx library to be able to log some query executions
type DB struct {
	*sqlx.DB
	logger *log.Logger
}

// Options needed when creating a new database connection
type Options struct {
	Host     string
	Port     string
	Name     string
	UserName string
	Password string
	SSLMode  string
}

// New database connection, logger can be nil if no log is wanted
func New(options *Options, logger *log.Logger) (*DB, error) {
	connString := fmt.Sprintf("dbname='%s' user='%s' password='%s' host='%s' port='%s' sslmode='%s'",
		options.Name, options.UserName, options.Password, options.Host, options.Port, options.SSLMode)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		logError(logger, err)
		return nil, err
	}
	return &DB{DB: db, logger: logger}, nil
}

// Query runs a specified query string with the specified arguments, then returns the query results
func (db *DB) Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	start := time.Now()
	rows, err := db.Queryx(query, args)
	duration := time.Since(start)
	if db.logger != nil {
		logSQL(db.logger, duration, query, args)
	}
	if err != nil && db.logger != nil {
		logError(db.logger, err)
	}
	return rows, err
}

// Queryx is the same with Query
func (db *DB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return db.Query(query, args)
}

// Get runs a 'SELECT' query and immediately save the single query result in the destination
func (db *DB) Get(destination interface{}, query string, args ...interface{}) error {
	start := time.Now()
	err := db.Get(destination, query, args)
	duration := time.Since(start)
	if db.logger != nil {
		logSQL(db.logger, duration, query, args)
	}
	if err != nil && db.logger != nil {
		logError(db.logger, err)
	}
	return err
}

// Select runs a 'SELECT' query and immediately save the query results in the destination
func (db *DB) Select(destination interface{}, query string, args ...interface{}) error {
	start := time.Now()
	err := db.Select(destination, query, args)
	duration := time.Since(start)
	if db.logger != nil {
		logSQL(db.logger, duration, query, args)
	}
	if err != nil && db.logger != nil {
		logError(db.logger, err)
	}
	return err
}

func logSQL(logger *log.Logger, duration time.Duration, query string, args ...interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("QUERY (%s):\n%s\n", duration.String(), query))
	for i, v := range args {
		buffer.WriteString(fmt.Sprintf("[%d] - %s\n", i+1, v))
	}
	logger.Println(buffer.String())
}

func logError(logger *log.Logger, err error) {
	logger.Println("Error: ", err.Error())
}
