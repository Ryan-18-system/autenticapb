package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/Ryan-18-system/autenticaPB/internal/dto"
	"github.com/Ryan-18-system/autenticaPB/internal/entity"
	"github.com/Ryan-18-system/autenticaPB/pkg/handlerException"
)

type UserRepositoryInterface interface {
	FindUserByCPF(cpf string) (*dto.UserImputDTO, error)
	Create(entidada *entity.User) error
	Update(entidade *entity.User) error
	DeleteByCPF(cpf string) error
	FindAll() ([]entity.User, error)
	FindByID(id string) (*entity.User, error)
	FindByCpf(cpf string) (*entity.User, error)
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
func (u *UserRepository) Create(entidade *entity.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
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
func (u *UserRepository) Update(entidade *entity.User) error {
	return nil
}
func (u *UserRepository) DeleteByCPF(cpf string) error {
	return nil
}
func (u *UserRepository) FindAll() ([]entity.User, error) {
	return nil, nil
}
func (u *UserRepository) FindByID(id string) (*entity.User, error) {
	return nil, nil
}
func (u *UserRepository) FindByCpf(cpf string) (*entity.User, error) {
	return nil, nil
}
