package common

type appResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(message string, data interface{}) *appResponse {
	return &appResponse{
		Message: message,
		Data:    data,
	}
}
