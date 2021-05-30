package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/page/internal/authapi"
	"github.com/supperdoggy/score/sctructs"
	authdata "github.com/supperdoggy/score/sctructs/service/auth"
)

// getAuthToken - make request to auth service and get token struct
func GetAuthToken(c *gin.Context) (sctructs.AuthToken, error) {
	token, err := c.Cookie("t")
	if err != nil {
		return sctructs.AuthToken{}, err
	}
	var req authdata.GetTokenByValueReq
	req.Token = token
	data, err := authapi.ApiV1(authdata.GetTokenByValuePath, "POST", req)
	if err != nil {
		return sctructs.AuthToken{}, err
	}
	var resp authdata.GetTokenByValueResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return sctructs.AuthToken{}, err
	}
	return resp.Token, nil
}
