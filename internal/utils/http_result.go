package utils

type HttpResult[T any] struct {
	HasError *bool   `json:"has_error"`
	Error    *string `json:"error"`
	Value    *T      `json:"value"`
}

func ErrorResult[T any](err error) HttpResult[T] {
	result := HttpResult[T]{
		HasError: PointerTo(true),
		Error:    PointerTo(err.Error()),
		Value:    nil,
	}

	return result
}

func ValueResult[T any](value T) HttpResult[T] {
	result := HttpResult[T]{
		HasError: PointerTo(false),
		Error:    nil,
		Value:    PointerTo(value),
	}

	return result
}
