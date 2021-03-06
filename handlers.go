package main

import (
	"io"
	"fmt"
	"errors"

	"github.com/labstack/echo"
	"github.com/AidHamza/optimizers-api/storage"
	"github.com/AidHamza/optimizers-api/messaging"

	"github.com/AidHamza/optimizers-api/pkg/helpers"
	op "github.com/AidHamza/optimizers-api/pkg/operation"
)

func cruncher(c echo.Context) error {
	//Process #1: Welcome the file
	file, err := c.FormFile("file")
	if err != nil {
		return throwHTTPError(c, badImageRequest, "FILE_UPLOAD", err)
	}

	//Process #2: Store the bytes into src
	src, err := file.Open()
	if err != nil {
		return throwHTTPError(c, inputFileOpen, "FILE_OPEN", err)
	}
	defer src.Close()

	//Process #3: Try to guess the type and reject if not acceptable
	imageType := guessImageMimeTypes(src)
	allowedImages := []string{"image/jpeg", "image/png"}
	isImageAllowed, _ := helpers.InArray(imageType, allowedImages)

	if isImageAllowed == false {
		return throwHTTPError(c, invalidImageType, "FILE_TYPE", errors.New("FILE_TYPE: invalid file type"))
	}

	var imageBucket string = "jpeg"
	if imageType == "image/png" {
		imageBucket = "png"
	}

	fileSize, err := src.Seek(0, io.SeekEnd) //2 = from end
	if err != nil {
		return throwHTTPError(c, invalidImageType, "FILE_GET_SIZE", err)
	}

	//Return to the head of the file
	src.Seek(0, io.SeekStart)

	// Generate an Unique Operation ID
	opId := helpers.RandomID()
	fileName := fmt.Sprintf("%d/%s", opId, file.Filename)

	//Process #4: Queue for processing
	err = storage.PutObject(src, imageBucket, fileName, imageType)
	if err != nil {
		return throwHTTPError(c, failedStoreFile, "FILE_STORAGE_FAILED", err)
	}

	producer, err := messaging.NewProducer()
	if err != nil {
		return throwHTTPError(c, failedQueueFile, "OP_QUEUE_FAILED", err)
	}
	
	operation, err := op.NewOperation(opId, file.Filename, imageType)
	if err != nil {
		return throwHTTPError(c, failedQueueFile, "OP_QUEUE_FAILED", err)
	}

	err = producer.PublishMessage(operation)
	if err != nil {
		return throwHTTPError(c, failedQueueFile, "OP_QUEUE_PUBLISH_FAILED", err)
	}


	result := &compressSuccess{
		Filename: file.Filename,
		Size:     fileSize,
		Id:     opId,
	}

	return c.JSON(200, result)
}
