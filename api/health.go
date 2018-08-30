package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetHealth() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, struct {
			status string
		}{status: "OK"})
	}
}