package common

import "errors"

var (
	InvalidRequestBody = errors.New("Invalid request body")
	SignUpFailed       = errors.New("Unable to sign up")
)
