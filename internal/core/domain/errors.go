package domain

import "errors"

var (
	UserNotFoundErr        = errors.New("user not found")
	UserAlreadyExistErr    = errors.New("user DNI or email already exist")
	AdminRoleNotAllowedErr = errors.New("admin role not allowed")
)
