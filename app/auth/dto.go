package auth

type RegisterRequestBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role string `json:"role"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}