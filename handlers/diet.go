package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCelebDiet(c echo.Context) error {
	id := c.QueryParam("id")
	diet, err := h.Controller.GetCelebDiet(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, diet)

}
