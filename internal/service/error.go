package service

import "errors"

var (
	ErrNotFound              = errors.New("not found")
	ErrEntityExists          = errors.New("entity exists")
	ErrCantHandleCredentials = errors.New("can't handle password conversion")
	ErrCredentialsMismatch   = errors.New("credentials mismatch")
)
