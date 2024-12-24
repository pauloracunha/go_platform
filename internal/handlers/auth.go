package handlers

import (
	"encoding/json"
	"net/http"
	"portal/internal/auth"
	"portal/internal/models"
	"portal/internal/services"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser registra um novo usuário
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
		return
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Senha), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao processar senha", http.StatusInternalServerError)
		return
	}
	user.Senha = string(hashedPassword)

	if err := services.DB.Create(&user).Error; err != nil {
		http.Error(w, "Erro ao salvar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login autentica um usuário e retorna o JWT
func Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := services.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
		return
	}

	// Verifica a senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(req.Password)); err != nil {
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(user.ID, 0) // 0 pode ser substituído pelo TenantID
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
