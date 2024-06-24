package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	_constant "sg-edts.com/edts-go-boilerplate/constant"

	_model "sg-edts.com/edts-go-boilerplate/model"
	_api "sg-edts.com/edts-go-boilerplate/pkg/api"
	_auth "sg-edts.com/edts-go-boilerplate/pkg/auth"
	_message "sg-edts.com/edts-go-boilerplate/pkg/message"
	_session "sg-edts.com/edts-go-boilerplate/usecase/session"
)

type Handler struct {
	SessionUseCase _session.Usecase
}

// getOwnSession show all owned session
func (a *Handler) getOwnSession(c *fiber.Ctx) error {
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

	data, err := a.SessionUseCase.GetOwnSession(ctx, claims)
	if err != nil {
		return _api.Failed(c, http.StatusInternalServerError, err)
	}

	return _api.Success(c, http.StatusOK, data)
}

// getSession show all owned session
func (a *Handler) getSession(c *fiber.Ctx) error {
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

	data, err := a.SessionUseCase.GetOwnSession(ctx, claims)
	if err != nil {
		return _api.Failed(c, http.StatusInternalServerError, err)
	}

	return _api.Success(c, http.StatusOK, data)
}

// dropOwnSession delete own session
func (a *Handler) dropOwnSession(c *fiber.Ctx) error {
	var req _model.DropSession

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

	data, err := a.SessionUseCase.DropOwnSession(ctx, claims, &req)
	if err != nil {
		var errApi *_api.Error
		if errors.As(err, &errApi) {
			return _api.Failed(c, errApi.Status, err)
		}

		return _api.Failed(c, http.StatusInternalServerError, err)
	}

	msg := map[string]interface{}{
		"message": data,
		"code":    http.StatusOK,
	}

	return _api.SuccessOnlyMessage(c, http.StatusOK, msg)
}

// dropSession delete session
func (a *Handler) dropSession(c *fiber.Ctx) error {
	var req _model.DropSession

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

	data, err := a.SessionUseCase.DropSession(ctx, claims, &req)
	if err != nil {
		var errApi *_api.Error
		if errors.As(err, &errApi) {
			return _api.Failed(c, errApi.Status, err)
		}

		return _api.Failed(c, http.StatusInternalServerError, err)
	}

	msg := map[string]interface{}{
		"message": data,
		"code":    http.StatusOK,
	}

	return _api.SuccessOnlyMessage(c, http.StatusOK, msg)
}
