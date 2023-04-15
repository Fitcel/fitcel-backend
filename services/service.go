package services

import "cloud.google.com/go/storage"

type Service struct {
	Food_Api_KEY string
	Food_Api_URL string

	StorageBucket *storage.BucketHandle
}
