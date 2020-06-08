package mysqlstore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/user3301/grpclab/pkg/types"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStore struct {
	Db *sql.DB
}

func (m MySQLStore) CreateUser(_ context.Context, userDetails types.UserDetails) error {
	// look if exists
	rows, err := m.Db.Query(fmt.Sprintf("SELECT * FROM userdetails WHERE username=%s", userDetails.Username))
	if err != nil {
		return err
	}
	if rows.Next() {
		return fmt.Errorf("username already exists %s", userDetails.Username)
	} else {
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
}

func (m MySQLStore) Verify(ctx context.Context, userDetails types.UserDetails) (bool, string) {
	rows, err := m.Db.Query(fmt.Sprintf("SELECT username, password FROM userdetails WHERE username=%s", userDetails.Username))
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
	if username != userDetails.Username && password != userDetails.Password {
		return false, fmt.Sprintf("user not exist or password incorrect")
	} else {
		return true, "ok"
	}
}

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
