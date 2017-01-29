package main

import (
	"github.com/labstack/echo"
)

type httpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Trace   string `json:"trace_id"`
}

func getErrorText(errorCode int) string {
	return errorText[errorCode]
}

func getError(errorCode int) *httpError {
	return &httpError{
		Code:    errorCode,
		Message: getErrorText(errorCode),
		Trace:   "h6d7vfd",
	}
}

func throwHTTPError(c echo.Context, errorCode int, glob *globals, errContext string, appError error) error {
	glob.Log.Error(getErrorText(errorCode), errContext, appError.Error())
	return c.JSON(errorHTTP[errorCode], getError(errorCode))
}
