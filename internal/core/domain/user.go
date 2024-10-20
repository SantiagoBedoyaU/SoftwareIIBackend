package domain

type UserRole int

const (
	AdminRole UserRole = iota
	MedicRole
	PatientRole
)

type UserTypeDNI int

const (
	TypeDniCC UserTypeDNI = iota
	TypeDniTI
	TypeDniTP // passport
)

type User struct {
	ID        string      `json:"id" bson:"_id,omitempty"`
	TypeDNI   UserTypeDNI `json:"type_dni" bson:"type_dni"`
	DNI       string      `json:"dni" bson:"dni"`
	FirstName string      `json:"first_name" bson:"first_name"`
	LastName  string      `json:"last_name" bson:"last_name"`
	Email     string      `json:"email" bson:"email"`
	Password  string      `json:"password,omitempty" bson:"password"`
	Role      UserRole    `json:"role" bson:"role"`
}

type UpdatePassword struct {
	NewPassword string `json:"new_password"`
}

type UpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UpdateRole struct {
	DNI string  `json:"dni"`
	NewRole int `json:"new_role"`
}
