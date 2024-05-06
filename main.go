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
		say("Hello", 1000, &wg)
		println("middle...")
		say("World", 2000, &wg)
		println("end...")

		wg.Wait()

		return c.String(http.StatusOK, "Done")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
