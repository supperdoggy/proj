package pagedata

import (
	"github.com/supperdoggy/score/sctructs"
)

type LoginRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token sctructs.AuthToken `json:"token"`
	Error string `json:"error"`
	OK bool `json:"ok"`
}
