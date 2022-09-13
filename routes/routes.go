package routes

import (
	"net/http"
	"rest_api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	g := e.Group("/v1")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World22!")
	})

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	}))

	g.GET("/pegawai", controllers.FetchAllPegawai)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai", controllers.UpdatePegawai)
	e.DELETE("/pegawai", controllers.DeletePegwai)

	e.POST("/login", controllers.CheckLogin)
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	// e.GET("/validation-test", controllers.TestStructValidation)
	return e
}
