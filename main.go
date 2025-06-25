package main

import (
	"goroutine-sample/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/sample", handler.SampleHandler)
	e.GET("/sample2", handler.Sample2Handler)

	e.Logger.Fatal(e.Start(":1323"))
}
