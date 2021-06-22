package main

import (
	h "github.com/alyaskastorm/Fintech-Test-Task/handler"
	"github.com/alyaskastorm/Fintech-Test-Task/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	st := storage.CreateConnection()
	storage.PrepareStorage(st)
	st.Close()

	storage := storage.NewURLStorage()
	handler := h.NewHandler(storage)

	e := echo.New()

	e.Logger.Print(e.POST("/short", handler.Short))
	e.Logger.Print(e.POST("/long", handler.Long))
	e.Logger.Fatal(e.Start(":1453"))
}

