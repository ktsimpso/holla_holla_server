package models

import (
	_ "github.com/lib/pq"
	"github.com/coocood/qbs"
	"time"
)

type GetFunction func() (interface{}, error)

type User struct {
	Id int64 `json:"id"`
	Name string `qbs:"size:64,index" json:"name"`
}

type Store struct {
	Id int64 `json:"id"`
	Name string `qbs:"size:64,index" json:"name"`
}

type Deal struct {
	Id int64 `json:"id"`

	UserId int64 `json:"user_id"`
	User *User `json:"user"`

	StoreId int64 `json:"store_id"`
	Store *Store `json:"store"`

	Created time.Time `json:"date"`
}

//TODO: handle errors
func CreateTables() error {
	qbs.Register("postgres", "user=holla dbname=hollaholla password=gimmiechocolate sslmode=disable", "hollaholla", qbs.NewPostgres())
	createTable(new(User))
	createTable(new(Store))
	createTable(new(Deal))
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

func GetDeals() (interface{}, error) {
	var deals []*Deal
	return getModels(&deals)
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
