package pagedata

import (
	"github.com/supperdoggy/score/sctructs"
	authdata "github.com/supperdoggy/score/sctructs/service/auth"
)

type LoginRequest struct {
	authdata.LoginReq
}

type LoginResponse struct {
	Token sctructs.AuthToken `json:"token"`
	Error string `json:"error"`
	OK bool `json:"ok"`
}
