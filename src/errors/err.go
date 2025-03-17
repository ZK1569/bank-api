package errs

import "errors"

var (
	BadId         = errors.New("Incorrect ID")
	NotFound      = errors.New("Not found")
	InternalError = errors.New("Internal error")
	BadRequest    = errors.New("Bad request")
)
