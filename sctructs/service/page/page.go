package pagedata

import (
	"github.com/supperdoggy/score/sctructs"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token sctructs.AuthToken `json:"token"`
	Error string             `json:"error"`
	OK    bool               `json:"ok"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Token sctructs.AuthToken `json:"token"`
	Error string             `json:"error"`
	OK    bool               `json:"ok"`
}
