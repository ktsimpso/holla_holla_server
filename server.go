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

	e.Get("/user", getHandlerFunc(models.GetUsers))
	e.Get("/store", getHandlerFunc(models.GetStores))
	e.Get("/deal", getHandlerFunc(models.GetDeals))

	e.Run(":3000")
}

func getHandlerFunc(gf models.GetFunction) echo.HandlerFunc {
	return func(c *echo.Context) error {
		items, err := gf()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, items)
	}
}
