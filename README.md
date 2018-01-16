# go-standard
A N-Tier architecture of go project to be used as a template for next development in go.



### References

<u>*General Guidance and Convention*</u>

* [Official Golang CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)
* [Effective Go](https://golang.org/doc/effective_go.html)
* [What's in a name by Andrew Gerrand](https://talks.golang.org/2014/names.slide#1)

<u>*Project Structure*</u>

* [Structuring Application by Ben Jhonson](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091)

<u>*Design Pattern*</u>

* [Mencoba Clean Architecture pada golang Oleh Iman Tumorang](https://medium.com/golangid/mencoba-golang-clean-architecture-c2462f355f41)
* [Clean Architecture by Uncle Bob](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)




## Table Of Contents

1. Naming Convention
2. Comment
3. Database
4. Error Handling
5. Imports
6. Project Structure
7. Packaging
8. Development Environment



### 2. Comment

 `go` provides **C-style** block comments (`/* */`) and line comments (`// here lies the comment`). Every package in `go` must have a package comments which defined preceeding the `package` clause.

```go
/*
Package postgresql will implement Repository interface which will be interact
with database connection to create / read / update / delete data from database.

No business logic will be handled in this level, business logic will be handled
in interactor / usecase level.
*/
package postgresql

type UserRepository struct { ... }

func (ur *UserRepository) GetServices() ([]*Service, error) { ... }
func (ur *UserRepository) GetServiceByID(id uint64) (*Service, error) { ... }
```



Comments for declaration *variable* / *struct* / *constant* / *interface* must be defined in full sentence. Comments must begin with the "thing" being described and end with a period. Ex.

```go
// GroupRepository will used as interface to be implemented in repository level.
type GroupRepository interface { ... }

// Lower will set text value to lowercase - will return error when conversion failed.
func Lower(text string) (string, error) { ... }
```

