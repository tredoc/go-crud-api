package service

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrEntityExists = errors.New("entity exists")
)
