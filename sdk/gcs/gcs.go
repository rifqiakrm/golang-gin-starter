package gcs

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"google.golang.org/api/option"

	"gin-starter/config"
)

// Upload uploads file to bucket
func Upload(c *gin.Context, config *config.Google, filename string) string {
	bucket := config.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("gcp-sa.json"))

	if err != nil {
		return errors.Wrap(err, "[CloudStorageService-Upload] error get config json").Error()
	}

	f, err := c.FormFile("file")
	if err != nil {
		return errors.Wrap(err, "[CloudStorageService-Upload] error get file").Error()
	}

	src, err := f.Open()
	if err != nil {
		return errors.Wrap(err, "[CloudStorageService-Upload] error open file").Error()
	}

	defer func() {
		if err := src.Close(); err != nil {
			fmt.Println("error while closing gin context form file :", err)
		}
	}()

	fileName := filename

	folder := c.PostForm("folder") + "/"

	fileStored := folder + fileName

	sw := storageClient.Bucket(bucket).Object(fileStored).NewWriter(ctx)

	if _, err = io.Copy(sw, src); err != nil {
		return errors.Wrap(err, "[CloudStorageService-Upload] error copy file").Error()
	}

	if err := sw.Close(); err != nil {
		return errors.Wrap(err, "[CloudStorageService-Upload] error close file").Error()
	}

	u, err := url.Parse(config.StorageEndpoint + "/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return errors.Wrap(err, "[CloudStorageService-Upload] error parse url").Error()
	}

	return u.String()
}

func Delete(config *config.Google, path string) error {
	bucket := config.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("gcp-sa.json"))

	if err != nil {
		return errors.Wrap(err, "[CloudStorageService-Delete] error get config json")
	}

	if err := storageClient.Bucket(bucket).Object(url.QueryEscape(path)).Delete(context.Background()); err != nil {
		return errors.Wrap(err, fmt.Sprintf("[CloudStorageService-Delete] unable to delete bucket %q, file %q", bucket, url.QueryEscape(path)))
	}

	return nil
}
