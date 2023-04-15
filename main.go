package main

import (
	"fitcel-backend/configuration"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	config := configuration.ConfigurationInit()
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/addCeleb", config.Handler.AddCeleb)
	e.GET("getCelebs", config.Handler.GetCelebs)
	e.GET("getCeleb", config.Handler.GetCeleb)

	e.GET("getCelebDiet", config.Handler.GetCelebDiet)

	e.POST("/addUser", config.Handler.AddUser)
	e.PUT("updateUser", config.Handler.UpdateUser)
	runmode := viper.GetString("runmode")
	PORT := os.Getenv("PORT")
	if runmode == "dev" {
		PORT = "8080"
	}
	e.Logger.Fatal(e.Start(":" + PORT))
}
