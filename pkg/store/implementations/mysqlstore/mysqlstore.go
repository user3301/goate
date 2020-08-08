package mysqlstore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/user3301/grpclab/pkg/types"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// MySQLStore mysql database implementation
type MySQLStore struct {
	Db *sql.DB
}

// CreateUser creates user record
func (m MySQLStore) CreateUser(_ context.Context, userDetails types.UserDetails) error {
	// look if exists
	const base = `SELECT * FROM userdetails WHERE username=%s`
	rows, err := m.Db.Query(base, userDetails.Username)
	if err != nil {
		return err
	}
	for rows.Next() {
		return fmt.Errorf("username already exists %s", userDetails.Username)
	}
	if rows.Err() != nil {
		return rows.Err()
	}
	stmt, err := m.Db.Prepare("INSERT userdetails SET username=?, password=?")
	if err != nil {
		return err
	}
	result, err := stmt.Exec(userDetails.Username, userDetails.Password)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return fmt.Errorf("query issue")
	}
	return nil
}

// Verify check if username and password matches database record
func (m MySQLStore) Verify(ctx context.Context, userDetails types.UserDetails) (verified bool, reason string) {
	const base = `SELECT username, password FROM userdetails WHERE username=%s`
	rows, err := m.Db.Query(base, userDetails.Username)
	if err != nil {
		return false, err.Error()
	}
	var username, password string
	for rows.Next() {
		err = rows.Scan(&username, &password)
		if err != nil {
			return false, err.Error()
		}
	}
	if rows.Err() != nil {
		return false, rows.Err().Error()
	}
	if username != userDetails.Username && password != userDetails.Password {
		return false, fmt.Sprintf("user not exist or password incorrect")
	}
	return true, "ok"
}

// UpdateUser updates user recprd
func (m MySQLStore) UpdateUser(ctx context.Context, details types.UserDetails) (bool, error) {
	stmt, err := m.Db.Prepare("UPDATE userdetails SET password=? WHERE username=?")
	if err != nil {
		return false, fmt.Errorf("prepare statement errored")
	}
	_, err = stmt.Exec(details.Password, details.Username)
	if err != nil {
		return false, fmt.Errorf("query execution errored")
	}
	return true, nil
}
