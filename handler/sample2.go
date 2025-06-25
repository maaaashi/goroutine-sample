package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Sample2Handler(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")

}
