package dto

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ResponseSucsess(code int, message string) Response[any] {
	return Response[any]{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func ResponseError(code int, message string) Response[any] {
	return Response[any]{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func ResponseSucsessData(code int, message string, data any) Response[any] {
	return Response[any]{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
