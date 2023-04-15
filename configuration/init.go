package configuration

import (
	"context"
	"fitcel-backend/controllers"
	"fitcel-backend/handlers"
	"fitcel-backend/models"
	"fitcel-backend/services"
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	Handler handlers.Handler
}

func getRunmode() string {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s", err.Error())
	}
	viper.SetEnvPrefix("global")
	runmode := viper.GetString("runmode")
	return runmode
}

func ConfigurationInit() Configuration {
	runmode := getRunmode()
	db := dbConnect(runmode)
	firbaseApp := firebaseInitialization(runmode)
	firbaseStorage, err := firbaseApp.Storage(context.Background())
	if err != nil {
		log.Panic("error while creating firebaase storage client", err)
	}
	storgageBucket, err := firbaseStorage.DefaultBucket()
	if err != nil {
		log.Panic("error while accessing storage bucket", err)
	}
	foodApiKey := viper.GetString(runmode + ".services.apininja.food.apiKey")
	foodApiUrl := viper.GetString(runmode + ".services.apininja.food.apiUrl")
	return Configuration{
		Handler: handlers.Handler{
			Controller: controllers.Controller{
				Model: models.Model{DB: db},
				Services: services.Service{
					Food_Api_KEY:  foodApiKey,
					Food_Api_URL:  foodApiUrl,
					StorageBucket: storgageBucket,
				},
			},
		},
	}
}
