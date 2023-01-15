package oldbackend

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
}
