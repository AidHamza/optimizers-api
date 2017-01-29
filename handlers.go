package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func cruncher(c echo.Context, glob *globals) error {
	file, err := c.FormFile("file")
	if err != nil {
		return throwHTTPError(c, badImageRequest, glob, "FILE_UPLOAD", err)
	}
	src, err := file.Open()
	if err != nil {
		return throwHTTPError(c, inputFileOpen, glob, "FILE_OPEN", err)
	}
	defer src.Close()

	imageType := guessImageMimeTypes(src)
	allowedImages := []string{"image/jpeg", "image/png"}
	isImageAllowed, _ := inArray(imageType, allowedImages)

	if isImageAllowed == false {
		return throwHTTPError(c, invalidImageType, glob, "FILE_TYPE", err)
	}

	return c.JSON(http.StatusBadRequest, err)
}
