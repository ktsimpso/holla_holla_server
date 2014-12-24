package models

import (
	"github.com/coocood/qbs"
	_ "github.com/lib/pq"
	"time"
)

type GetFunction func() (interface{}, error)

type User struct {
	Id            int64  `json:"id"`
	Name          string `qbs:"size:64,index" json:"name"`
	Email         string `qbs:"size:256,index,unique" json:"-"`
	EmailVerified bool   `json:"-"`
	Password      []byte `json:"-"`
	CanLogin      bool   `json:"-"`
}

type Store struct {
	Id   int64  `json:"id"`
	Name string `qbs:"size:64,index" json:"name"`
}

type Deal struct {
	Id int64 `json:"id"`

	UserId int64 `json:"user_id"`
	User   *User `json:"user"`

	StoreId int64  `json:"store_id"`
	Store   *Store `json:"store"`

	Created time.Time `json:"date"`
}

func init() {
	createTables()

	// Insure there is an anonymous user
	anonymousUser := User{
		Id:            1,
		Name:          "Anonymous",
		Email:         "",
		EmailVerified: false,
		Password:      []byte{},
		CanLogin:      false,
	}
	q, err := qbs.GetQbs()
	if err != nil {
		panic(err)
	}
	defer q.Close()

	_, err = q.Save(&anonymousUser)
	if err != nil {
		panic(err)
	}
}

func createTables() {
	qbs.Register("postgres", "user=holla dbname=hollaholla password=gimmiechocolate sslmode=disable", "hollaholla", qbs.NewPostgres())
	err := createTable(new(User))
	if err != nil {
		panic(err)
	}
	err = createTable(new(Store))
	if err != nil {
		panic(err)
	}
	err = createTable(new(Deal))
	if err != nil {
		panic(err)
	}
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
