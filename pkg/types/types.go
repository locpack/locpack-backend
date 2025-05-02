package types

type PackStatus = string

type AccessToken struct {
	Value        string
	RefreshToken string
	ExpiresIn    float64
}

type Token struct {
	Valid     bool
	Username  string
	Email     string
	ExpiresIn float64
	Roles     []string
}
