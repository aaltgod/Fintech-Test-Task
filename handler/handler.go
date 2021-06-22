package handler

import (
	"net/http"

	"github.com/alyaskastorm/Fintech-Test-Task/storage"
	"github.com/alyaskastorm/Fintech-Test-Task/tools"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	storage storage.Storage
}

func NewHandler(storage storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) Short(c echo.Context) error {
	var url storage.URL

	if err := c.Bind(&url); err != nil {
		return err
	}

	// Check if long URL exists
	shortURL := h.storage.GetShort(url.Name)
	if shortURL != "" {
		url.Name = shortURL
		return c.JSON(http.StatusOK, url)
	}

	ID := tools.RandomID(8)
	shortURL = "https://mysite.com/" + ID

	if err := h.storage.Create(ID, url.Name, shortURL); err != nil {
		return err
	}

	url.Name = shortURL

	return c.JSON(http.StatusOK, url)
}

func (h *Handler) Long(c echo.Context) error {
	var url storage.URL

	if err := c.Bind(&url); err != nil {
		return err
	}

	longURL := h.storage.GetLong(url.Name)
	url.Name = longURL

	return c.JSON(http.StatusOK, url)
}