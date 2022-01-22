package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

/**
* Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
* A: Considering we are developing an application not a common package (no dependencies on current project) - then we should
* wrap the error and return for constructing better debugging info like stack trace.
**/

type User struct {
	UserId string
	Name string
}

func main () {

}

// DAO level - ORM example
func findAll() {
	// Sql connections/select query
	var err error
	_, err = sqlFindAll()
	if err != nil {
		return nil, errors.Wrap(err, "findAll error")
	}
}

// DAO implementation
func sqlFindAll() ([]User, error) {
	return nil, sql.ErrNoRows
}