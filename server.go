package main

import (
	"github.com/go-martini/martini"
	"github.com/ktsimpso/holla_holla_server/models"
	"encoding/json"
)

func main() {
	models.RegisterDb()
	models.CreateUserTable()

	m := martini.Classic()
	m.Get("/", func () string {
		return "Hello World!"
	})

	m.Get("/user", func () (int, string) {
		users, err := models.GetUsers()
		if err != nil {
			return 500, "An error occured!"
		}

		data, err := json.Marshal(users)

		if err != nil {
			return 500, "An error occured!!"
		}

		return 200, string(data)
	})

	m.Run()
}
