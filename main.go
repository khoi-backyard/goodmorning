package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/khoi/goodmorning/api"
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

	v1 := e.Group("/api/v1")

	v1.GET("/health", api.GetHealth())

	e.Logger.Fatal(e.Start(":8888"))
}
