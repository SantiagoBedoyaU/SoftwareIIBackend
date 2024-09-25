package domain

type Auth struct {
	DNI      string `json:"dni"`
	Password string `json:"password"`
}

type RecoverPassword struct {
	DNI string `json:"dni"`
}

type ResetPassword struct {
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}
