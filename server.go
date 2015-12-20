package main

import (
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

	e.Get("/user", func(c *echo.Context) error {
		users, err := models.GetUsers()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, users)
	})

	e.Get("/store", func(c *echo.Context) error {
		stores, err := models.GetStores()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, stores)
	})

	e.Get("/deal", func(c *echo.Context) error {
		deals, err := models.GetDeals()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, deals)
	})

	e.Run(":3000")
}
