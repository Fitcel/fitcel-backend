package handlers

import (
	"fitcel-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) AddUser(c echo.Context) error {
	var params models.User

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := h.Controller.AddUser(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h Handler) GetUserByUUID(c echo.Context) error {
	uuid := c.QueryParam("uuid")
	user, err := h.Controller.GetUserByUUID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (h Handler) UpdateUser(c echo.Context) error {
	userID := c.QueryParam("UUID")
	dietID := c.QueryParam("dietID")
	if err := h.Controller.UpdateUser(userID, dietID); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "success")
}
