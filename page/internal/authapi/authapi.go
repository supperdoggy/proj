package authapi

import (
	"encoding/json"
	"github.com/supperdoggy/score/sctructs"
	"github.com/supperdoggy/score/sctructs/communication"
	authdata "github.com/supperdoggy/score/sctructs/service/auth"
)

// ApiV1 - sends request to apiv1 auth
func ApiV1(path, method string, data interface{}) (result []byte, err error) {
	marshaled, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	result, err = communication.MakeHttpRequest(authdata.AuthRoute+sctructs.ApiV1+path, method, marshaled)
	if err != nil {
		return nil, err
	}
	return
}
