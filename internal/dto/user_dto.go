package dto

type UserImputDTO struct {
	Username string `json:"username"`
	Nome     string `json:"nome"`
	Password string `json:"password"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
}
