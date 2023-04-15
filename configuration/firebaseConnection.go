package configuration

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func firebaseInitialization() *firebase.App {
	opt := option.WithCredentialsFile("./conf/fitcel.json")
	config := &firebase.Config{
		StorageBucket: "fitcel-e2f0e.appspot.com",
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Panic("error initializing firebase app", err)
	}
	return app
}
