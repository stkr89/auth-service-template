package common

import "errors"

var (
	InvalidRequestBody = errors.New("Invalid request body")
	SomethingWentWrong = errors.New("Something went wrong")
	SignUpFailed       = errors.New("Unable to sign up")
)
