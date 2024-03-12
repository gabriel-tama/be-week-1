package helper

import (
	"strings"

	"github.com/gabriel-tama/be-week-1/types"
)



func BuildResponse(status bool, message string, data interface{}) types.Response {
	res := types.Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) types.Response {
	splittedError := strings.Split(err, "\n")
	res := types.Response{
		Status:  false,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}
	return res
}