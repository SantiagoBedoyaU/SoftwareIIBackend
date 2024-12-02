package domain

import "errors"

var (
	ErrUserNotFound             = errors.New("user not found")
	ErrUserAlreadyExist         = errors.New("user DNI or email already exist")
	ErrUserEmailAlreadyInUse    = errors.New("user email already in use")
	ErrAdminRoleNotAllowed      = errors.New("admin role not allowed")
	ErrNotAnAdminRole           = errors.New("user role is not admin")
	ErrNotAMedicRole            = errors.New("doctorId don't belong to a medic user")
	ErrAlreadyHaveAnAppointment = errors.New("already have an appointment")
	ErrNotValidDates            = errors.New("end date must be greater than start date")
	ErrNotValidEndDate			    = errors.New("end date must be before the actual date")
	ErrAppointmentNotFound      = errors.New("appointment not found")
	ErrInvalidIDFormat          = errors.New("id is invalid")
	ErrUnavailableTimeNotFound  = errors.New("unavailable time not found")
)
