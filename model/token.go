package model

import "myapp/data_type"

type Token struct {
	// system
	AccessToken          string
	AccessTokenExpiredAt data_type.DateTime
	TokenType            string
}
