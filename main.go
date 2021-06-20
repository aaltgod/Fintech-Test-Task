package main

import (
	h "Fintech-Test-Task/handler"
	"Fintech-Test-Task/postgresql"
	"github.com/labstack/echo/v4"
)

func main() {
	postgresql.CreateConnection()

	storage := postgresql.NewURLStorage()
	handler := h.NewHandler(storage)

	e := echo.New()

	e.POST("/short", handler.Short)
	e.POST("/long", handler.Long)
	e.Logger.Fatal(e.Start(":1453"))
}

