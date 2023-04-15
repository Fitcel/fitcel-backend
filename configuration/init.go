package configuration

import (
	"context"
	"fitcel-backend/controllers"
	"fitcel-backend/handlers"
	"fitcel-backend/models"
	"fitcel-backend/services"
	"log"
)

type Configuration struct {
	Handler handlers.Handler
}

func ConfigurationInit() Configuration {
	db := dbConnect()
	firbaseApp := firebaseInitialization()
	firbaseStorage, err := firbaseApp.Storage(context.Background())
	if err != nil {
		log.Panic("error while creating firebaase storage client", err)
	}
	storgageBucket, err := firbaseStorage.DefaultBucket()
	if err != nil {
		log.Panic("error while accessing storage bucket", err)
	}

	return Configuration{
		Handler: handlers.Handler{
			Controller: controllers.Controller{
				Model: models.Model{DB: db},
				Services: services.Service{
					Food_Api_KEY:  "IAkaKHtBzs/uJ81Ezscwkg==dZlQxOQBqn4LbADY",
					Food_Api_URL:  "https://api.api-ninjas.com/v1/nutrition?query=",
					StorageBucket: storgageBucket,
				},
			},
		},
	}
}
