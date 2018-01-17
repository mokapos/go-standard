package postgresql_test

import (
	"sort"
	"testing"

	"github.com/mokapos/go-standard/postgresql"
	"github.com/mokapos/go-standard/testingutil"
)

func TestFindAllEmails(t *testing.T) {
	t.Parallel()
	db, destroy := postgresql.CreateTestDB(t)
	// TODO: still can't drop database automatically, need to do research about this
	defer destroy()

	postgresql.SetSchema(t, db, "admin")
	repo := postgresql.NewAdminRepository(db, db)

	// Test for query and execution errors
	emails, err := repo.FindAllEmail()
	testingutil.Ok(t, err)

	// Test for return value errors
	// All emails must appeared exactly once
	sort.Strings(emails)
	testingutil.Equals(t, "black@goseiger.com", emails[0])
	testingutil.Equals(t, "blue@goseiger.com", emails[1])
	testingutil.Equals(t, "dups@goseiger.com", emails[2])
	testingutil.Equals(t, "pink@goseiger.com", emails[3])
	testingutil.Equals(t, "red@goseiger.com", emails[4])
	testingutil.Equals(t, "yellow@goseiger.com", emails[5])
}
