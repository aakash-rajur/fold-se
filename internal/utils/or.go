package utils

func Or[T any](value *T, def *T) *T {
	if value == nil {
		return def
	}

	return value
}
