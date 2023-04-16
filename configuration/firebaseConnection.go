package configuration

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func firebaseInitialization(runmode string) *firebase.App {
	credentialFile := viper.GetString(runmode + ".services.firebase.service_file")
	opt := option.WithCredentialsFile(credentialFile)
	bucketName := viper.GetString(runmode + ".services.firebase.bucket")
	config := &firebase.Config{
		StorageBucket: bucketName,
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Panic("error initializing firebase app", err)
	}
	return app
}
