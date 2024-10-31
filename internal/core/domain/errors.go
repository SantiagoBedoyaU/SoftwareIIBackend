package domain

import "errors"

var (
	ErrUserNotFound          	= errors.New("user not found")
	ErrUserAlreadyExist      	= errors.New("user DNI or email already exist")
	ErrUserEmailAlreadyInUse 	= errors.New("user email already in use")
	ErrAdminRoleNotAllowed   	= errors.New("admin role not allowed")
	ErrNotAnAdminRole  		 	= errors.New("user role is not admin")
	ErrAlreadyHaveAnAppointment = errors.New("already have an appointment")
)
