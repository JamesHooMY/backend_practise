package user

import "errors"

var (
	ErrUserExisted        = errors.New("user existed")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyDeleted = errors.New("user already deleted")
)
