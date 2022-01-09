package responses

import "net/http"

// function response bad request
func BadRequestResponse(message string) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusBadRequest,
		"Message": message,
	}
	return result
}

// function response success dengan return value data
func SuccessResponseData(message string, data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Data":    data,
		"Message": message,
	}
	return result
}
