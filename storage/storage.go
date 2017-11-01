package storage

import (
	"fmt"
	"io"
	"github.com/minio/minio-go"
	"github.com/AidHamza/optimizers-api/config"
	"github.com/AidHamza/optimizers-api/log"
)

func PutObject(fileReader io.Reader, bucket string, fileName string, fileType string) error {
	s3Client, err := minio.New(
		fmt.Sprintf("%s:%d", config.App.Storage.Host, config.App.Storage.Port),
		config.App.Storage.AccessKey,
		config.App.Storage.SecretKey,
		config.App.Storage.TLS)

	if err != nil {
		return err
	}

	n, err := s3Client.PutObject(bucket, fileName, fileReader, fileType)
	if err != nil {
		return err
	}

	log.Logger.Info("File stored", "FILENAME", fileName, "SIZE", n)

	return nil
}