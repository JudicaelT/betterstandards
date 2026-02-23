package assert

func Must[T any](v T, err error) T {
	NoError(err)
	return v
}
