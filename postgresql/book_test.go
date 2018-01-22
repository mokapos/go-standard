package postgresql_test

import (
	"testing"

	"github.com/mokapos/go-standard/postgresql"
	"github.com/mokapos/go-standard/testingutil"
)

func TestBookFindAll(t *testing.T) {
	db := postgresql.NewTestDBConnection(t)
	defer db.Close()
	postgresql.RunSQLFile(t, db, "book")
	repo := postgresql.NewBookRepository(db, db)

	// All books must be appeared
	// TODO: unfinished
	_, err := repo.FindAll()
	testingutil.Ok(t, err)
}
