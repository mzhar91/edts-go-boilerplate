package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	
	_config "sg-edts.com/edts-go-boilerplate/config"
)

var authInstance *auth = nil

// auth initialization
func InitClaims(debug bool) *auth {
	if authInstance == nil {
		privateBytes, err := os.ReadFile(_config.Cfg.Jwt.PrivateKey)
		if err != nil {
			logrus.Error(err)
			panic(err)
		}
		
		privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
		if err != nil {
			logrus.Error(err)
			panic(err)
		}
		
		publicBytes, err := os.ReadFile(_config.Cfg.Jwt.PublicKey)
		if err != nil {
			panic(err)
		}
		
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
		if err != nil {
			panic(err)
		}
		
		authInstance = &auth{
			privateKey: privateKey,
			publicKey:  publicKey,
			debug:      debug,
		}
	}
	
	return authInstance
}

func (m *auth) ClaimsContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(
			&ClaimsContext{
				privateKey: m.privateKey,
				publicKey:  m.publicKey,
				debug:      m.debug,
				Context:    c,
			},
		)
	}
}

func (c *ClaimsContext) GenerateAccessToken(id string, username string, app string) (string, error, int) {
	var duration int64
	
	if app == "mobile" {
		duration = _config.Cfg.Jwt.AccessPeriodMobile
	} else if app == "bo" {
		duration = _config.Cfg.Jwt.AccessPeriodBo
	} else {
		return "", fmt.Errorf("app was not declared"), http.StatusInternalServerError
	}
	
	expirationTime := time.Now().Add(time.Duration(duration) * time.Minute)
	claims := &Claims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	
	return tokenString, nil, http.StatusOK
}

func (c *ClaimsContext) GenerateRefreshToken(id string, username string, app string) (string, error, int) {
	var duration int64
	
	if app == "mobile" {
		duration = _config.Cfg.Jwt.RefreshPeriodMobile
	} else if app == "bo" {
		duration = _config.Cfg.Jwt.RefreshPeriodBo
	} else {
		return "", fmt.Errorf("app was not declared"), http.StatusInternalServerError
	}
	
	expirationTime := time.Now().Add(time.Duration(duration) * time.Minute)
	claims := &Claims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	
	return tokenString, nil, http.StatusOK
}

func (c *ClaimsContext) Claims() (*Claims, string, error, int) {
	s := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(s, "Bearer ")
	if len(splitToken) < 2 {
		err := fmt.Errorf(fmt.Sprintf("Malformed token"))
		return nil, "", err, http.StatusBadRequest
	}
	
	token := splitToken[1]
	if token == "" {
		err := fmt.Errorf(fmt.Sprintf("Token required"))
		return nil, "", err, http.StatusUnauthorized
	}
	
	token = strings.Trim(splitToken[1], " ")
	
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(
		token, claims, func(token *jwt.Token) (interface{}, error) {
			return c.publicKey, nil
		},
	)
	if err != nil {
		if c.debug {
			return nil, "", err, http.StatusForbidden
		}
		
		err := fmt.Errorf(fmt.Sprintf("Invalid Key"))
		return nil, "", err, http.StatusForbidden
	}
	if !tkn.Valid {
		err := fmt.Errorf(fmt.Sprintf("Invalid Token"))
		return nil, "", err, http.StatusForbidden
	}
	
	return claims, token, nil, http.StatusOK
}

func (c *ClaimsContext) UpdateToken(claims *Claims, app string) (string, string, error, int) {
	var duration int64
	
	if app == "mobile" {
		duration = _config.Cfg.Jwt.RefreshPeriodMobile
	} else if app == "bo" {
		duration = _config.Cfg.Jwt.RefreshPeriodBo
	} else {
		return "", "", fmt.Errorf("app was not declared"), http.StatusInternalServerError
	}
	
	var expireOffset = duration * 60
	
	expTime, err := claims.GetExpirationTime()
	if err != nil {
		return "", "", err, http.StatusUnauthorized
	}
	
	remaining := getTokenRemaining(expTime.Unix())
	
	if remaining < expireOffset/4 {
		accessToken, err, errCode := c.GenerateAccessToken(
			claims.ID, claims.Username, app,
		)
		if err != nil {
			return "", "", err, errCode
		}
		
		refreshToken, err, errCode := c.GenerateRefreshToken(claims.ID, claims.Username, app)
		if err != nil {
			return "", "", err, errCode
		}
		
		return accessToken, refreshToken, nil, http.StatusOK
	}
	
	return "", "", nil, http.StatusOK
}

func getTokenRemaining(timestamp interface{}) int64 {
	if validity, ok := timestamp.(int64); ok {
		tm := time.Unix(validity, 0)
		remainder := tm.Sub(time.Now())
		if remainder > 0 {
			return int64(remainder.Seconds())
		}
	}
	
	return 0
}
