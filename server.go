package main

import (
	"net/http"

	h "github.com/alyaskastorm/Fintech-Test-Task/handler"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.POST("/short", h.Short)
	e.Logger.Fatal(e.Start(":1453"))
}

