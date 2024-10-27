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
	DNI       string      `json:"dni" bson:"dni" binding:"required"`
	FirstName string      `json:"first_name" bson:"first_name"  binding:"required"`
	LastName  string      `json:"last_name" bson:"last_name"`
	Email     string      `json:"email" bson:"email"  binding:"required,email"`
	Password  string      `json:"password,omitempty" bson:"password"`
	Role      UserRole    `json:"role" bson:"role"`
	Address   string      `json:"address" bson:"address"`
	Phone     string      `json:"phone" bson:"phone"`
	IsActive  bool        `json:"is_active" bson:"is_active" default:"true"`
}

type Admin struct {
	User
}

type Medic struct {
	User
	Salary      float64 `json:"salary" bson:"salary"`
	Especiality string  `json:"especiality" bson:"especiality"`
}

type Patient struct {
	User
}

type UpdatePassword struct {
	NewPassword string `json:"new_password"`
}

type UpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}

type UpdateRole struct {
	DNI     string   `json:"dni"`
	NewRole UserRole `json:"new_role"`
}
