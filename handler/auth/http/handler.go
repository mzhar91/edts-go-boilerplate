package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	_config "sg-edts.com/edts-go-boilerplate/config"
	_constant "sg-edts.com/edts-go-boilerplate/constant"
	_model "sg-edts.com/edts-go-boilerplate/model"
	_api "sg-edts.com/edts-go-boilerplate/pkg/api"
	_auth "sg-edts.com/edts-go-boilerplate/pkg/auth"
	_message "sg-edts.com/edts-go-boilerplate/pkg/message"
	_credential "sg-edts.com/edts-go-boilerplate/usecase/credential"
)

type Handler struct {
	CredentialUseCase _credential.Usecase
}

// @Tags Auth
// @Summary Add Credential
// @Description Add a new credential for user
// @Accept  json
// @Produce  json
// @Param req body _model.AddCredentialRequest
// @Router /auth [post]
func (a *Handler) addCredential(c *fiber.Ctx) error {
	var req _model.AddCredentialRequest

	if err := c.BodyParser(&req); err != nil {
		return _api.Failed(c, http.StatusBadRequest, _api.FromErrorCode(_message.IncorrectFormat))
	}

	if okBind, err := _api.IsRequestValid(&req); !okBind {
		return _api.Failed(c, http.StatusBadRequest, err)
	}

	ctx := c.UserContext()
	if ctx == nil {
		ctx = context.Background()
	}

	cCtx := c.Locals(_constant.ContextAuth).(*_auth.ClaimsContext)
	claims, _, err, authStatus := cCtx.Claims()
	if err != nil {
		logrus.Error(err)
		return _api.Failed(c, authStatus, err)
	}

	req.RequestInfo.IpNumber = c.IP()
	req.RequestInfo.Host = c.Hostname()

	data, err, msg := a.CredentialUseCase.AddCredential(ctx, _constant.ScopeMobile, claims, &req)
	if err != nil {
		var errApi *_api.Error
		if errors.As(err, &errApi) {
			return _api.Failed(c, errApi.Status, err)
		}

		return _api.Failed(c, http.StatusInternalServerError, err)
	}

	return _api.SuccessWithMessage(c, http.StatusOK, data, msg)
}

// @Tags Auth
// @Summary Sign in
// @Description Sign in user
// @Accept  json
// @Produce  json
// @Param req body _model.SignInRequest
// @Router /auth [post]
func (a *Handler) signIn(c *fiber.Ctx) error {
	var req _model.SignInRequest

	if err := c.BodyParser(&req); err != nil {
		return _api.Failed(c, http.StatusBadRequest, _api.FromErrorCode(_message.IncorrectFormat))
	}

	if okBind, err := _api.IsRequestValid(&req); !okBind {
		return _api.Failed(c, http.StatusBadRequest, err)
	}

	ctx := c.UserContext()
	if ctx == nil {
		ctx = context.Background()
	}

	cCtx := c.Locals(_constant.ContextAuth).(*_auth.ClaimsContext)

	req.RequestInfo.IpNumber = c.IP()
	req.RequestInfo.Host = c.Hostname()

	data, err, msg := a.CredentialUseCase.SignIn(ctx, _constant.ScopeMobile, cCtx, &req)
	if err != nil {
		var errApi *_api.Error
		if errors.As(err, &errApi) {
			return _api.Failed(c, errApi.Status, err)
		}

		return _api.Failed(c, http.StatusInternalServerError, err)
	}

	return _api.SuccessWithMessage(c, http.StatusOK, data, msg)
}

func (a *Handler) signOut(c *fiber.Ctx) error {
	var req _model.SignOutRequest

	if err := c.BodyParser(&req); err != nil {
		return _api.Failed(c, http.StatusBadRequest, _api.FromErrorCode(_message.IncorrectFormat))
	}

	if okBind, err := _api.IsRequestValid(&req); !okBind {
		return _api.Failed(c, http.StatusBadRequest, err)
	}

	ctx := c.UserContext()
	if ctx == nil {
		ctx = context.Background()
	}

	cCtx := c.Locals(_constant.ContextAuth).(*_auth.ClaimsContext)
	claims, _, err, authStatus := cCtx.Claims()
	if err != nil {
		logrus.Error(err)
		return _api.Failed(c, authStatus, err)
	}

	req.RequestInfo.IpNumber = c.IP()
	req.RequestInfo.Host = c.Hostname()

	err, msg := a.CredentialUseCase.SignOut(ctx, _constant.ScopeMobile, claims.Username, &req)
	if err != nil {
		var errApi *_api.Error
		if errors.As(err, &errApi) {
			return _api.Failed(c, errApi.Status, err)
		}

		return _api.Failed(c, http.StatusInternalServerError, err)
	}

	return _api.SuccessWithMessage(c, http.StatusOK, nil, msg)
}

func (a *Handler) refreshToken(c *fiber.Ctx) error {
	ctx := c.UserContext()
	if ctx == nil {
		ctx = context.Background()
	}

	cCtx := c.Locals(_constant.ContextAuth).(*_auth.ClaimsContext)
	_, _, err, authStatus := cCtx.Claims()
	if err != nil {
		logrus.Error(err)
		return _api.Failed(c, authStatus, err)
	}

	data, err, msg := a.CredentialUseCase.RefreshToken(ctx, _constant.ScopeMobile, cCtx)
	if err != nil {
		var errApi *_api.Error
		if errors.As(err, &errApi) {
			return _api.Failed(c, errApi.Status, err)
		}

		return _api.Failed(c, http.StatusInternalServerError, err)
	}

	return _api.SuccessWithMessage(c, http.StatusOK, data, msg)
}

func (a *Handler) checkTokenAvailability(c *fiber.Ctx) error {
	ctx := c.UserContext()
	if ctx == nil {
		ctx = context.Background()
	}

	cCtx := c.Locals(_constant.ContextAuth).(*_auth.ClaimsContext)
	claims, _, err, authStatus := cCtx.Claims()
	if err != nil {
		logrus.Error(err)
		return _api.Failed(c, authStatus, err)
	}

	access, refresh, err, authStatus := cCtx.UpdateToken(claims, _constant.ScopeMobile)
	if err != nil {
		var errApi *_api.Error
		if errors.As(err, &errApi) {
			return _api.Failed(c, errApi.Status, err)
		}

		return _api.Failed(c, http.StatusInternalServerError, err)
	}

	if access != "" && refresh != "" {
		return _api.SuccessWithMessage(
			c, http.StatusOK, _model.SignInResponse{
				AccessToken:     access,
				RefreshToken:    refresh,
				TokenExpiration: time.Now().Add(time.Duration(_config.Cfg.Jwt.AccessPeriodMobile) * time.Minute).Unix(),
			}, map[string]interface{}{"message": "Token revoked", "code": http.StatusOK},
		)
	}

	return _api.SuccessWithMessage(
		c, http.StatusOK, nil, map[string]interface{}{"message": "Token available", "code": http.StatusOK},
	)
}
