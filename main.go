package main

import (
	h "Fintech-Test-Task/handler"
	"Fintech-Test-Task/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	db := storage.CreateConnection()
	storage.PrepareStorage(db)
	db.Close()

	storage := storage.NewURLStorage()
	handler := h.NewHandler(storage)

	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)

	e.Logger.Print(e.POST("/short", handler.Short))
	e.Logger.Print(e.POST("/long", handler.Long))
	e.Logger.Fatal(e.Start(":1453"))
}

