package main

import (
	"github.com/go-martini/martini"
	"github.com/ktsimpso/holla_holla_server/models"
	"encoding/json"
	"net/http"
)

func main() {
	models.RegisterDb()
	models.CreateTables()

	m := martini.Classic()

	m.Use(func(c martini.Context, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	m.Get("/user", func () (int, string) {
		users, err := packIntoJson(models.GetUsers)
		if err != nil {
			return 500, "An error occured!"
		}

		return 200, users
	})

	m.Get("/store", func () (int, string) {
		stores, err := packIntoJson(models.GetStores)
		if err != nil {
			return 500, "An error occured!"
		}

		return 200, stores
	})

	m.Run()
}

func packIntoJson(get models.GetFunction) (string, error) {
	items, err := get()

	if err != nil {
		return "", err
	}

	data, err := json.Marshal(items)

	if err != nil {
		return "", err
	}

	stringData := string(data)
	if stringData == "null" {
		stringData = "[]"
	}

	return stringData, nil
}
