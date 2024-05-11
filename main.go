package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserInfo struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Todo struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Todo    Todo   `json:"todo"`
}

func GetInfo(channel chan<- UserInfo) {
	defer close(channel)
	time.Sleep(time.Duration(2000) * time.Millisecond)

	info := UserInfo{
		Id:      "f516154d-28d6-41f3-bde7-cc1ad9a88e2e",
		Name:    "山田太郎",
		Address: "〇〇県〇〇市",
	}

	channel <- info

	println("INFOを読み込みました")
}

func GetTodo(userId string, channel chan<- Todo) {
	defer close(channel)
	time.Sleep(time.Duration(2000) * time.Millisecond)

	println(userId)

	channel <- Todo{
		Id:      "97d2f4eb-3052-4f0d-8bbb-19d12d559933",
		Title:   "TODOその1",
		Content: "TODOその1の内容",
	}

	println("TODOを読み込みました")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		println("starting...")

		infoChannel := make(chan UserInfo)
		go GetInfo(infoChannel)

		userInfo := <-infoChannel
		todoChannel := make(chan Todo)
		go GetTodo(userInfo.Id, todoChannel)

		todo := <-todoChannel

		user := User{
			Id:      userInfo.Id,
			Name:    userInfo.Name,
			Address: userInfo.Address,
			Todo:    todo,
		}

		println("end...")
		return c.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
