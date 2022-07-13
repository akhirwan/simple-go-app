package model

type EmptyResponse struct {
}

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type WebResponseError struct {
	Error ErrorResponse `json:"error"`
}
