package models

type Response struct {
	Content  interface{} `json:"content"`
	Message  string      `json:"message"`
	Code     int         `json:"code"`
}

// GetResponse func to Generate response
func GetResponse(result interface{}, code int, message string) Response {
	return Response{
		Code:    code,
		Content: result,
		Message: message,
	}
}