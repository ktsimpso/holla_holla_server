package models

import (
	_ "github.com/lib/pq"
	"github.com/coocood/qbs"
)

func RegisterDb() {
	qbs.Register("postgres", "user=holla dbname=hollaholla password=gimmiechocolate sslmode=disable", "hollaholla", qbs.NewPostgres())	
}

type User struct {
	Id int64
	Name string `qbs:"size:64,index"`
}

func CreateUserTable() error {
	migration, err := qbs.GetMigration()
	if err != nil {
		return err
	}

	defer migration.Close()
	return migration.CreateTableIfNotExists(new(User))
}

func GetUsers() ([]*User, error) {
	var users []*User;

	q, err := qbs.GetQbs()
	if err != nil {
		return users, err
	}
	defer q.Close()

	err = q.FindAll(&users)
	return users, err
}
