package main

import (
	"io"
	"os"
	"errors"

	"github.com/labstack/echo"
)

func cruncher(c echo.Context, glob *globals) error {
	//Process #1: Welcome the file
	file, err := c.FormFile("file")
	if err != nil {
		return throwHTTPError(c, badImageRequest, glob, "FILE_UPLOAD", err)
	}

	//Process #2: Store the bytes into src
	src, err := file.Open()
	if err != nil {
		return throwHTTPError(c, inputFileOpen, glob, "FILE_OPEN", err)
	}
	defer src.Close()

	//Process #3: Try to guess the type and reject if not acceptable
	imageType := guessImageMimeTypes(src)
	allowedImages := []string{"image/jpeg", "image/png"}
	isImageAllowed, _ := inArray(imageType, allowedImages)

	if isImageAllowed == false {
		return throwHTTPError(c, invalidImageType, glob, "FILE_TYPE", errors.New("FILE_TYPE: invalid file type"))
	}

	fileSize, err := src.Seek(0, 2) //2 = from end
	if err != nil {
		return throwHTTPError(c, invalidImageType, glob, "FILE_GET_SIZE", err)
	}

	//Return to the head of the file
	src.Seek(0, 0)

	//Process #4: Store + Queue for processing
	dst, err := os.Create(uploadPath + file.Filename)
	if err != nil {
		return throwHTTPError(c, failedSaveFile, glob, "FILE_TMP_SAVE", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return throwHTTPError(c, failedCopyFile, glob, "FILE_COPY", err)
	}

	compressOutput := compressImage(file.Filename, imageType)

	result := &compressSuccess{
		Filename: file.Filename,
		Size:     fileSize,
		Output:   compressOutput,
	}

	return c.JSON(200, result)
}
