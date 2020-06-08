package mysqlstore

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/user3301/grpclab/pkg/store"
)

type Config struct {
	ConnStr string `yaml:"conn"`
	Host    string `yaml:"host, omitempty"`
	DDLPath string `yaml:"ddl, omitempty"`
}

func (c Config) MySQLStore() (store.UserStorer, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/userdetails?charset=utf8", c.Host, c.ConnStr))
	if err != nil {
		return nil, err
	}
	mySqlStore := MySQLStore{db}
	if c.DDLPath != "" {
		err := initDB(db, c.DDLPath)
		if err != nil {
			return nil, err
		}
	}
	return mySqlStore, nil
}

func initDB(db *sql.DB, f string) error {
	bytes, err := ioutil.ReadFile(filepath.Clean(f))
	if err != nil {
		return err
	}
	_, err = db.Exec(string(bytes))
	if err != nil {
		return err
	}
	return nil
}
