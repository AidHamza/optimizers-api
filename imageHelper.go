package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime"

	_ "golang.org/x/image/webp"
)

// Guess image format from gif/jpeg/png/webp
func guessImageFormat(r io.Reader) (format string, err error) {
	_, format, err = image.DecodeConfig(r)
	return
}

// Guess image mime types from gif/jpeg/png/webp
func guessImageMimeTypes(r io.Reader) string {
	format, _ := guessImageFormat(r)
	if format == "" {
		return ""
	}
	return mime.TypeByExtension("." + format)
}
