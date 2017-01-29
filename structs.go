package main

import "gopkg.in/inconshreveable/log15.v2"

type globals struct {
	Name string
	Log  log15.Logger
}

const (
	invalidImageType = 100
	badImageRequest  = 101
	inputFileOpen    = 102
)

var errorText = map[int]string{
	invalidImageType: "Invalid image type",
	badImageRequest:  "Bad image payload",
	inputFileOpen:    "Cannot open the input file",
}

var errorHTTP = map[int]int{
	invalidImageType: 415,
	badImageRequest:  400,
	inputFileOpen:    406,
}
