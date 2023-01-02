package lib

import "errors"

var (
	ErrInvalidID          = errors.New("ERROR_INVALID_ID")
	ErrInvalidArgJSONBody = errors.New("ERROR_INVALID_JSON")
)
