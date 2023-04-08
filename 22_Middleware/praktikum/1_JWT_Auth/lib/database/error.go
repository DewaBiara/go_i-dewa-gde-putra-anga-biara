package database

import "errors"

var (
	// ErrInvalidID is returned when an invalid ID is provided to a method like Delete.
	ErrInvalidID = errors.New("provided id was invalid")
)
