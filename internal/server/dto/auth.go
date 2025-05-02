package dto

type Register struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Refresh struct {
	Value string `json:"value"`
}

type AccessToken struct {
	Value        string
	RefreshToken string
	ExpiresIn    float64
}
