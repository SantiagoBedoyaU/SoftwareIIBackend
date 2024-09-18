package domain

type Auth struct {
	DNI      string `json:"dni"`
	Password string `json:"password"`
}

type RecoverPassword struct {
	Email string `json:"email"`
}
