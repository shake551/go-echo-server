package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var textFlag = flag.String("text", "Hello World!", "")

func main() {
	flag.Parse()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	if *textFlag == "foo" {
		for i := 0; i < 1000000; i++ {
			fmt.Println(i)
		}
	}
	return c.String(http.StatusOK, *textFlag+"\n")
}
