package main

import (
	"encoding/json"
	"github.com/ktsimpso/holla_holla_server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	e := echo.New()
	e.StripTrailingSlash()
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(cors.Default().Handler)
	e.Use(func(c *echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
		return nil
	})

	e.Get("/user", func(c *echo.Context) error {
		users, err := packIntoJson(models.GetUsers)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, users)
	})

	e.Get("/store", func(c *echo.Context) error {
		stores, err := packIntoJson(models.GetStores)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, stores)
	})

	e.Get("/deal", func(c *echo.Context) error {
		deals, err := packIntoJson(models.GetDeals)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, deals)
	})

	e.Run(":3000")
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
