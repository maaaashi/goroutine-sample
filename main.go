package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type UserInfo struct {
	id      string
	name    string
	address string
}

type Todo struct {
	id      string
	title   string
	content string
}

type User struct {
	id      string
	name    string
	address string
	todo    Todo
}

func GetInfo(duration int, wg *sync.WaitGroup, chanel chan<- UserInfo) {
	defer wg.Done()
	time.Sleep(time.Duration(duration) * time.Millisecond)

	chanel <- UserInfo{
		id:      "f516154d-28d6-41f3-bde7-cc1ad9a88e2e",
		name:    "山田太郎",
		address: "〇〇県〇〇市",
	}

	println("INFOを読み込みました")
}

func GetTodo(duration int, wg *sync.WaitGroup, channel chan<- Todo) {
	defer wg.Done()
	time.Sleep(time.Duration(duration) * time.Millisecond)

	channel <- Todo{
		id:      "97d2f4eb-3052-4f0d-8bbb-19d12d559933",
		title:   "TODOその1",
		content: "TODOその1の内容",
	}

	println("TODOを読み込みました")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		var wg sync.WaitGroup
		wg.Add(2)

		println("starting...")

		infoChannel := make(chan UserInfo)
		go GetInfo(0, &wg, infoChannel)

		todoChannel := make(chan Todo)
		go GetTodo(0, &wg, todoChannel)

		wg.Wait()
		println("end...")

		userInfo := <-infoChannel
		todo := <-todoChannel

		user := User{
			id:      userInfo.id,
			name:    userInfo.name,
			address: userInfo.address,
			todo:    todo,
		}

		return c.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
