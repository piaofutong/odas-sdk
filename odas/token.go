package odas

import (
	"encoding/json"
	"net/http"
)

type TokenRequest struct {
	GrantType string `json:"grantType"`
	AccessId  string `json:"accessId"`
	AccessKey string `json:"accessKey"`
}

func (o *TokenRequest) Api() string {
	return "/token"
}

func (o *TokenRequest) Body() []byte {
	body := TokenRequest{
		GrantType: "client_credential",
		AccessId:  o.AccessId,
		AccessKey: o.AccessKey,
	}
	b, _ := json.Marshal(body)
	return b
}

func (o *TokenRequest) Method() string {
	return http.MethodPost
}

func (o *TokenRequest) ContentType() string {
	return "application/json"
}

func NewTokenRequest(accessId, accessKey string) *TokenRequest {
	return &TokenRequest{AccessId: accessId, AccessKey: accessKey}
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
	Trial       bool   `json:"trial"`
}
