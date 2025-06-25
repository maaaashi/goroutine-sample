package main

import (
	"goroutine-sample/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/sample", handler.SampleHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
