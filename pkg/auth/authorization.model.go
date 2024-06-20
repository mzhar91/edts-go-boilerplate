package auth

import (
	"crypto/rsa"
	"encoding/json"
	"strings"
	
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type auth struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	debug      bool
}

type ClaimsContext struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	debug      bool
	echo.Context
}

type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type multiString string

func (ms *multiString) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		switch data[0] {
		case '"':
			var s string
			if err := json.Unmarshal(data, &s); err != nil {
				return err
			}
			*ms = multiString(s)
		case '[':
			var s []string
			if err := json.Unmarshal(data, &s); err != nil {
				return err
			}
			*ms = multiString(strings.Join(s, ","))
		}
	}
	
	return nil
}
