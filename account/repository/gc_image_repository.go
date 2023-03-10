package repository

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/ndenisj/go_mem/account/model"
	"github.com/ndenisj/go_mem/account/model/apperrors"
)

type gcImageRepository struct {
	Storage    *storage.Client
	BucketName string
}

// NewImageRepository is a factory for initializing image repository
func NewImageRepository(gcClient *storage.Client, bucketName string) model.ImageRepository {
	return &gcImageRepository{
		Storage:    gcClient,
		BucketName: bucketName,
	}
}

func (r *gcImageRepository) UpdateProfile(ctx context.Context, objName string, imageFile multipart.File) (string, error) {
	bckt := r.Storage.Bucket(r.BucketName)
	object := bckt.Object(objName)
	wc := object.NewWriter(ctx)

	// set cache control so profile image will be served fresh by browsers
	// To do this with object handle, you first have to upload then update
	wc.ObjectAttrs.CacheControl = "Cache-Control:no-cache, max-age=0"

	// multipart.File have a reader
	if _, err := io.Copy(wc, imageFile); err != nil {
		log.Printf("unable to write file to Google Cloud Storage: %v\n", err)
		return "", apperrors.NewInternal()
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("writer.close: %v", err)
	}

	imageURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", r.BucketName, objName)

	return imageURL, nil
}

func (r *gcImageRepository) DeleteProfile(ctx context.Context, objName string) error {
	bckt := r.Storage.Bucket(r.BucketName)

	object := bckt.Object(objName)

	if err := object.Delete(ctx); err != nil {
		log.Printf("Failed to delete image object with ID: %s from GC storage\n", objName)
		return apperrors.NewInternal()
	}

	return nil
}
