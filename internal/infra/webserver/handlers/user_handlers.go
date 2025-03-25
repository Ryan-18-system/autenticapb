package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ryan-18-system/autenticaPB/internal/dto"
	"github.com/Ryan-18-system/autenticaPB/internal/infra/database"
	"github.com/Ryan-18-system/autenticaPB/internal/service"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	service := service.NewUserService(database.NewUserRepository())
	user := dto.UserImputDTO{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}
	err = service.CreateUser(user)
	if err != nil {
		http.Error(w, "Erro ao criar o usuário", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
