package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

func say(text string, duration int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(duration) * time.Millisecond)
	println(text)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		var wg sync.WaitGroup
		wg.Add(2)

		println("starting...")
		go say("World", 2000, &wg)
		go say("Hello", 1000, &wg)

		wg.Wait()
		println("end...")

		return c.String(http.StatusOK, "Done")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
