package dto

type ApiResponse[T any] struct {
	Key     string `json:"key"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ApiResponseList struct {
	Key      string      `json:"key"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Page     int         `json:"page"`
	Total    int64       `json:"total"`
	LastPage int         `json:"last_page"`
}
