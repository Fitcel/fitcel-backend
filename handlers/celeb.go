package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AddCeleb(c echo.Context) error {
	jsonInput := c.FormValue("json")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}
	err = h.Controller.AddCeleb(jsonInput, avatar)
	if err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}
	return c.HTML(http.StatusOK, "success")
}

func (h *Handler) GetCelebs(c echo.Context) error {
	celebs, err := h.Controller.GetCelebs()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, celebs)

}
func (h *Handler) GetCeleb(c echo.Context) error {
	id := c.QueryParam("id")
	celeb, err := h.Controller.GetCeleb(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, celeb)
}
func (h *Handler) GetCelebByDietID(c echo.Context) error {
	id := c.QueryParam("dietID")
	celeb, err := h.Controller.GetCelebByDietID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, celeb)
}
