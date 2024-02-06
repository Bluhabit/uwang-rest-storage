package models

type BaseResponse[T any] struct {
	StatusCode uint16 `json:"status_code"`
	Data       T      `json:"data"`
	Message    string `json:"message"`
}

func (BaseResponse[T]) Success(data T, message string) BaseResponse[T] {
	return BaseResponse[T]{
		StatusCode: 200,
		Data:       data,
		Message:    message,
	}
}

func (BaseResponse[T]) BadRequest(data T, message string) BaseResponse[T] {
	return BaseResponse[T]{
		StatusCode: 400,
		Data:       data,
		Message:    message,
	}
}
