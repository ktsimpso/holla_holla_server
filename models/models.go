package models

import (
	_ "github.com/lib/pq"
	"github.com/coocood/qbs"
)

type T interface{}
type GetFunction func() (T, error)

func RegisterDb() {
	qbs.Register("postgres", "user=holla dbname=hollaholla password=gimmiechocolate sslmode=disable", "hollaholla", qbs.NewPostgres())	
}

type User struct {
	Id int64 `json:"id"`
	Name string `qbs:"size:64,index"json:"name"`
}

type Store struct {
	Id int64 `json:"id"`
	Name string `qbs:"size:64,index"json:"name"`
}

//TODO: handle errors
func CreateTables() error {
	createTable(new(User))
	createTable(new(Store))
	return nil
}

func createTable(t T) error {
	migration, err := qbs.GetMigration()
	if err != nil {
		return err
	}

	defer migration.Close()
	return migration.CreateTableIfNotExists(t)
}

func GetUsers() (T, error) {
	var users []*User;

	q, err := qbs.GetQbs()
	if err != nil {
		return nil, err
	}
	defer q.Close()

	err = q.FindAll(&users)
	return users, err
}

func GetStores() (T, error) {
	var stores []*Store;

	q, err := qbs.GetQbs()
	if err != nil {
		return nil, err
	}
	defer q.Close()

	err = q.FindAll(&stores)
	return stores, err
}
