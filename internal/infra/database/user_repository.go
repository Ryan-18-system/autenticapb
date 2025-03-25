package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/Ryan-18-system/autenticaPB/internal/dto"
	"github.com/Ryan-18-system/autenticaPB/pkg/handlerException"
)

type UserRepositoryInterface interface {
	FindUserByCPF(cpf string) (*dto.UserImputDTO, error)
	Create(entidade dto.UserImputDTO) error
	Update(entidade dto.UserImputDTO) error
	Delete(entidade dto.UserImputDTO) error
	FindAll() ([]dto.UserImputDTO, error)
	FindByID(id string) (*dto.UserImputDTO, error)
}
type UserRepository struct {
	DataBase *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{DataBase: GetDB()}
}

func (u *UserRepository) FindUserByCPF(cpf string) (*dto.UserImputDTO, error) {
	return nil, nil
}
func (u *UserRepository) Create(entidade dto.UserImputDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	tx, err := u.DataBase.BeginTx(ctx, nil)
	if err != nil {
		return handlerException.HandlerError(err)
	}
	// própria doc do go indica - https://go.dev/doc/database/execute-transactions
	defer tx.Rollback()
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO user (cpf, username, nome, password, email) VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return handlerException.HandlerError(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, entidade.CPF, entidade.Username, entidade.Nome, entidade.Password, entidade.Email)
	if err != nil {
		return handlerException.HandlerError(err)
	}
	err = tx.Commit()
	if err != nil {
		return handlerException.HandlerError(err)
	}
	return nil
}
func (u *UserRepository) Update(entidade dto.UserImputDTO) error {
	return nil
}
func (u *UserRepository) Delete(entidade dto.UserImputDTO) error {
	return nil
}
func (u *UserRepository) FindAll() ([]dto.UserImputDTO, error) {
	return nil, nil
}
func (u *UserRepository) FindByID(id string) (*dto.UserImputDTO, error) {
	return nil, nil
}
