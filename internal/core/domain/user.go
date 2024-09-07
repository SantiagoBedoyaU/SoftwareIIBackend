package domain

type Role int

const (
	Admin Role = iota
	Medic
	Patient
)

type TypeDNI int

const (
	CC TypeDNI = iota
	TI
	TP // passport
)

type User struct {
	ID        string  `json:"id"`
	TypeDNI   TypeDNI `json:"type_dni"`
	DNI       string  `json:"dni"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Role      Role    `json:"role"`
}
