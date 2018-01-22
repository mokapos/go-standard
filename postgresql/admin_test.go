package postgresql_test

import (
	"sort"
	"testing"

	"github.com/mokapos/go-standard/postgresql"
	"github.com/mokapos/go-standard/testingutil"
)

func TestAdminFindByID(t *testing.T) {
	db := postgresql.NewTestDBConnection(t)
	defer db.Close()
	postgresql.RunSQLFile(t, db, "admin")
	repo := postgresql.NewAdminRepository(db, db)

	// An error must not be returned on a present ID
	_, err := repo.FindByID(1)
	testingutil.Ok(t, err)

	// An error must be returned on invalid ID
	_, err = repo.FindByID(100)
	testingutil.Error(t, err, "sql: no rows in result set")
}

func TestAdminFindByNameAndPassword(t *testing.T) {
	db := postgresql.NewTestDBConnection(t)
	defer db.Close()
	postgresql.RunSQLFile(t, db, "admin")
	repo := postgresql.NewAdminRepository(db, db)

	// An instance of admin must be returned on correct name and password combination
	admin, err := repo.FindByNameAndPassword("Alata", "GoseiRed")
	testingutil.Ok(t, err)
	testingutil.Equals(t, "Alata", admin.Name)
	testingutil.Equals(t, "GoseiRed", admin.Password)
	testingutil.Equals(t, "red@goseiger.com", admin.Email)

	// An error must be encountered on wrong name and password combination
	admin, err = repo.FindByNameAndPassword("Alata", "GoseiSilver")
	testingutil.Error(t, err, "sql: no rows in result set")
}

func TestAdminFindAllEmails(t *testing.T) {
	db := postgresql.NewTestDBConnection(t)
	defer db.Close()
	postgresql.RunSQLFile(t, db, "admin")
	repo := postgresql.NewAdminRepository(db, db)

	// All emails must appeared exactly once
	emails, err := repo.FindAllEmail()
	testingutil.Ok(t, err)
	sort.Strings(emails)
	testingutil.Equals(t, "black@goseiger.com", emails[0])
	testingutil.Equals(t, "blue@goseiger.com", emails[1])
	testingutil.Equals(t, "dups@goseiger.com", emails[2])
	testingutil.Equals(t, "pink@goseiger.com", emails[3])
	testingutil.Equals(t, "red@goseiger.com", emails[4])
	testingutil.Equals(t, "yellow@goseiger.com", emails[5])
}
