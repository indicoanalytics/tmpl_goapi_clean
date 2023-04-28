package constants

import "errors"

var (
	ErrInvalidUserAgent = errors.New("invalid user agent")
	ErrMissingUserAgent = errors.New("user agent is missing")
)
