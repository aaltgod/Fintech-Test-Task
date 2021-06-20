package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type URL struct {
	Name string `json:"url"`
}

func Short(c echo.Context) error {
	url := new(URL)
	if err := c.Bind(url); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, url)
}