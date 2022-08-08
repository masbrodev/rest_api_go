package controllers

import (
	"net/http"
	"rest_api/models"
	"rest_api/vendor/github.com/labstack/echo/v4"

	_ "github.com/labstack/echo/v4"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
