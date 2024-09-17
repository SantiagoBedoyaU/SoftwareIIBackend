package port

import "softwareIIbackend/internal/core/domain"

type AuthService interface {
	GetAuthToken(dni string, role domain.UserRole) (string, error)
}
