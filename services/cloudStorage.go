package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

func (s Service) StoreAvatarinCloud(objName string, avatar multipart.File) (avatarUrl string, err error) {
	pictureName := base64.StdEncoding.EncodeToString([]byte(objName))
	object := s.StorageBucket.Object(pictureName)

	wc := object.NewWriter(context.Background())
	_, err = io.Copy(wc, avatar)
	if err != nil {
		return "", err
	}
	err = wc.Close()
	if err != nil {
		return "", err
	}
	acl := object.ACL()
	if err := acl.Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}
	return fmt.Sprintf("https://storage.googleapis.com/fitcel-e2f0e.appspot.com/%s", pictureName), err
}

func (s Service) DeleteAvatarinCloud(objname string) error {
	pictureName := base64.StdEncoding.EncodeToString([]byte(objname))
	object := s.StorageBucket.Object(pictureName)
	return object.Delete(context.Background())
}
