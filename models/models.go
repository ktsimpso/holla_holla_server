package models

import (
	_ "github.com/lib/pq"
	"github.com/coocood/qbs"
)

type GetFunction func() (interface{}, error)

func RegisterDb() {
	qbs.Register("postgres", "user=holla dbname=hollaholla password=gimmiechocolate sslmode=disable", "hollaholla", qbs.NewPostgres())	
}

type User struct {
	Id int64 `json:"id"`
	Name string `qbs:"size:64,index" json:"name"`
}

type Store struct {
	Id int64 `json:"id"`
	Name string `qbs:"size:64,index" json:"name"`
}

//TODO: handle errors
func CreateTables() error {
	createTable(new(User))
	createTable(new(Store))
	return nil
}

func GetUsers() (interface{}, error) {
	var users []*User
	return getModels(&users)
}

func GetStores() (interface{}, error) {
	var stores []*Store
	return getModels(&stores)
}

func createTable(t interface{}) error {
	migration, err := qbs.GetMigration()
	if err != nil {
		return err
	}

	defer migration.Close()
	return migration.CreateTableIfNotExists(t)
}

func getModels(items interface{}) (interface{}, error) {
	q, err := qbs.GetQbs()
	if err != nil {
		return nil, err
	}
	defer q.Close()

	err = q.FindAll(items)
	return items, err
}
