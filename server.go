package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/codegangsta/martini-contrib/cors"
	"github.com/ktsimpso/holla_holla_server/models"
	"encoding/json"
	"net/http"
)

func main() {
	models.CreateTables()

	m := martini.New()
	m.Use(martini.Logger())
	m.Use(gzip.All())
	m.Use(martini.Recovery())
	m.Use(cors.Allow(&cors.Options{
		AllowAllOrigins: true,
	}))

	m.Use(func(c martini.Context, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	r := martini.NewRouter()

	r.Get("/user", func () (int, string) {
		users, err := packIntoJson(models.GetUsers)
		if err != nil {
			return 500, "An error occured!"
		}

		return 200, users
	})

	r.Get("/store", func () (int, string) {
		stores, err := packIntoJson(models.GetStores)
		if err != nil {
			return 500, "An error occured!"
		}

		return 200, stores
	})

	r.Get("/deal", func () (int, string) {
		deals, err := packIntoJson(models.GetDeals)
		if err != nil {
			return 500, "An error occured"
		}

		return 200, deals
	})

	m.Action(r.Handle)
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
