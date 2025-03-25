package entity

import (
	"github.com/Ryan-18-system/autenticaPB/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Username string    `json:"username"`
	CPF      string    `json:"cpf"`
	Nome     string    `json:"nome"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
}

func NewUser(username, password, email, nome, cpf string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Username: username,
		CPF:      cpf,
		Nome:     nome,
		Password: string(hash),
		Email:    email,
	}, nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password),
		[]byte(password))
	return err == nil
}
