package activity

import "errors"

var (
	ErrorInternalServer = errors.New("Internal Server Error")
	ErrorNotFound       = errors.New("Your requested Activity is not found")
	ErrorConflict       = errors.New("Your Activity already exists")
)
