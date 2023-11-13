package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type AuthTokenResponse struct {
	AccessToken          string             `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	AccessTokenExpiredAt data_type.DateTime `json:"access_token_expired_at" format:"YYYY-MM-DDThh:mm:ssZ" example:"2006-01-02T03:04:05+07:00"`
	TokenType            string             `json:"token_type" example:"Bearer"`
} // @name AuthTokenResponse

func NewAuthTokenResponse(token model.Token) AuthTokenResponse {
	return AuthTokenResponse{
		AccessToken:          token.AccessToken,
		AccessTokenExpiredAt: token.AccessTokenExpiredAt,
		TokenType:            token.TokenType,
	}
}

func NewAuthTokenResponseP(token model.Token) *AuthTokenResponse {
	r := NewAuthTokenResponse(token)
	return &r
}
