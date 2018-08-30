package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	echoMW "github.com/labstack/echo/middleware"
	"net/http"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	e := echo.New()
	e.Use(echoMW.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Good morning")
	})
	e.Logger.Fatal(e.Start(":8888"))
}
