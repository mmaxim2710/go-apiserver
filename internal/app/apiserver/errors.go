package apiserver

import "errors"

var (
	ErrEnvVariableNotFound      = errors.New("env variable not found")
	ErrIncorrectEmailOrPassword = errors.New("incorrect email or password")
)
