package controllers

import (
	"net/http"
	"rest_api/models"

	"github.com/labstack/echo/v4"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	result, err := models.StorePegawai(nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
