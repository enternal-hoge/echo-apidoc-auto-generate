package main

import (

	yaag_echo "./echo"
	"github.com/betacraft/yaag/yaag"

	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	yaag.Init(&yaag.Config{On: true, DocTitle: "Hoge Test", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(yaag_echo.Document())

	// Routes
	e.GET("/hoge", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "{ \"hoge\" : \"page\" }")
}

