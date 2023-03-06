package api

// InternalError is returned in case of internal error of app
type InternalError struct {
	Error       string
	Description string
}
