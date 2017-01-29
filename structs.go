package main

import "gopkg.in/inconshreveable/log15.v2"

type globals struct {
	Name string
	Log  log15.Logger
}

type compressSuccess struct {
	Filename string `json:"file"`
	Size     int64  `json:"size"`
	Output   string `json:"output"`
}

const (
	invalidImageType = 100
	badImageRequest  = 101
	inputFileOpen    = 102
	failedSaveFile   = 103
	failedCopyFile   = 104
	failedFileInfo   = 105

	//Some settings
	uploadPath   = "./uploads/"
	downloadPath = "./downloads/"
)

var errorText = map[int]string{
	invalidImageType: "Invalid image type",
	badImageRequest:  "Bad image payload",
	inputFileOpen:    "Cannot open the input file",
	failedSaveFile:   "Failed to save the temporary file",
	failedCopyFile:   "Failed to copy the file",
	failedFileInfo:   "Failed to get the file info",
}

var errorHTTP = map[int]int{
	invalidImageType: 415,
	badImageRequest:  400,
	inputFileOpen:    406,
	failedSaveFile:   500,
	failedCopyFile:   500,
	failedFileInfo:   500,
}
