package helpers
/**
 * @author : Donald Trieu
 * @created : 4/23/21, Friday
**/
import "strings"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObject struct {
}

func BuildResponse(success bool, message string, data interface{}) Response {
	res := Response{
		Success: success,
		Message: message,
		Data:    data,
	}

	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitError := strings.Split(err, "\n")
	res := Response{
		Success: false,
		Message: message,
		Error:   splitError,
		Data:    data,
	}

	return res
}

