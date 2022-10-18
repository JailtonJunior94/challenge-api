package dtos

type HttpResponse struct {
	StatusCode int
	Data       interface{} `json:"data,omitempty"`
}

func NewHttpResponse(statusCode int, data interface{}) *HttpResponse {
	return &HttpResponse{StatusCode: statusCode, Data: data}
}
