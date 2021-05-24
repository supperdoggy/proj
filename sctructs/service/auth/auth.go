package authdata

import (
	"github.com/supperdoggy/score/sctructs"
	"time"
)

type CheckTokenReq struct {
	Token string `json:"token"`
}

type CheckTokenRes struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type RegisterReq struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
}

type RegisterRes struct {
	Error string             `json:"error"`
	Token sctructs.AuthToken `json:"token"`
}

type LoginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token sctructs.AuthToken `json:"token"`
	OK    bool               `json:"ok"`
	Error string             `json:"error"`
}
