package handler

import (
	"Fintech-Test-Task/postgresql"
	"Fintech-Test-Task/tools"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	storage postgresql.Storage
}

func NewHandler(storage postgresql.Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) Short(c echo.Context) error {
	var url postgresql.URL

	if err := c.Bind(&url); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, url)
}

func (h *Handler) Long(c echo.Context) error {
	var url postgresql.URL

	if err := c.Bind(url); err != nil {
		return err
	}

	shortURL, err := tools.ShortURL(url.Name)
	if err != nil {
		return err
	}

	err = h.storage.Create(url.Name, shortURL)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, shortURL)
}