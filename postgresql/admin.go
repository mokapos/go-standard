package postgresql

import (
	"github.com/mokapos/go-standard"
)

// AdminRepository is a repository for admin table
type AdminRepository struct {
	Master *DB
	Slave  *DB
}

// NewAdminRepository creates a new repository for admin table
func NewAdminRepository(master *DB, slave *DB) *AdminRepository {
	return &AdminRepository{
		Master: master,
		Slave:  slave,
	}
}

// FindByID searches for any admin with specified id in the database
func (repo *AdminRepository) FindByID(id uint64) (*standard.Admin, error) {
	var admin standard.Admin
	err := repo.Slave.Get(&admin, "SELECT * FROM admins WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// FindByNameAndPassword searches for any admin with specified name and password
func (repo *AdminRepository) FindByNameAndPassword(name string, password string) (*standard.Admin, error) {
	var admin standard.Admin
	err := repo.Slave.Get(&admin, "SELECT * FROM admins WHERE name=$1 AND password=$2", name, password)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// FindAllEmail returns all admin emails
func (repo *AdminRepository) FindAllEmail() ([]string, error) {
	rows, err := repo.Slave.Query("SELECT DISTINCT email FROM admins")
	defer rows.Close()
	// Check for error during query executions
	if err != nil {
		return nil, err
	}
	// Manually scan for query result row by row, the other way is to use 'db.Select'
	// but it needs to construct multiple full structs.
	// FYI, golang has no way to nil a struct, it only sets default values to all struct members
	var emails []string
	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}
	// Check for error during iterations
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return emails, nil
}
