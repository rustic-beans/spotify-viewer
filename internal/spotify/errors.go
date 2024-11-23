package spotify

type NotAuthenticatedError struct {
}

func (e NotAuthenticatedError) Error() string {
	return "client is not authenticated yet"
}
