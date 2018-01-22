package root

import "errors"

// Admin is the one who manages the library
type Admin struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Timestampable
}

// AdminInteractor is a collection of method that can be done to Book object.
type AdminInteractor struct {
	adminRepo AdminRepository
}

// AdminRepository is a collection of method that can call Book table in the database.
type AdminRepository interface {
	FindByID(id uint64) (*Admin, error)
	FindByNameAndPassword(name string, password string) (*Admin, error)
	FindAllEmail() ([]string, error)
	Create(admin *Admin) error
	Update(id uint64, admin *Admin) error
	Delete(id uint64) error
}

// NewAdminInteractor returns an instance of AdminInteractor.
func NewAdminInteractor(adminRepo AdminRepository) *AdminInteractor {
	return &AdminInteractor{adminRepo: adminRepo}
}

// TryLogin searches the database to find an admin with matching name and password and returns the admin id.
// This function returns an uncontrolled error or "invalid name or password" controlled error.
func (itrc *AdminInteractor) TryLogin(name string, password string) (uint64, error) {
	admin, err := itrc.adminRepo.FindByNameAndPassword(name, password)
	if err.Error() == "sql: no rows in result set" {
		return 0, errors.New("invalid name or password")
	}
	if err != nil {
		return 0, err
	}
	return admin.ID, nil
}

// UpdatePassword updates the password of specified admin id, given the old password for verifivation.
// This function returns an uncontrolled error or "incorrect password" controlled error.
func (itrc *AdminInteractor) UpdatePassword(id uint64, old string, new string) error {
	admin, err := itrc.adminRepo.FindByID(id)
	if err != nil {
		return err
	}
	if admin.Password != old {
		return errors.New("incorrect password")
	}
	admin.Password = new
	err = itrc.adminRepo.Update(id, admin)
	if err != nil {
		return err
	}
	return nil
}
