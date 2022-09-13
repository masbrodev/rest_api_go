package controllers

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type Customer struct {
	Nama    string `validate:"required"`
	Alamat  string `validate:"required"`
	Telepon string `validate:"required"`
}

func TestVal(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Nama:    "Bamas",
		Alamat:  "Situbondo",
		Telepon: "0918728768",
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}
