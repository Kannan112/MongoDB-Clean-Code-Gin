package unit

import "strings"

// Response is used for static shape json return
type Response struct {
	Status  bool        `josn:"status" bson:"json"`
	Message string      `josn:"message"`
	Errors  interface{} `josn:"error,omitemty"`
	Data    interface{} `json:"data,omitemty"`
}

// ErrorResponse method is used to inject data value to dynamic failed response
func ErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

// SuccessResponse method is used to inject data value to dynamic success response
func SuccessResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}
