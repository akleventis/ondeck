package lib

import "errors"

var (
	ErrInvalidID          = errors.New("ERROR_INVALID_ID")
	ErrInvalidArgJSONBody = errors.New("ERROR_INVALID_JSON")
	ErrDrinkNotFound      = errors.New("ERROR_DRINK_NOT_FOUND")
	ErrPersonNotFound     = errors.New("ERROR_PERSON_NOT_FOUND")
)
