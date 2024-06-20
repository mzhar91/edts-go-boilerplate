package http

import (
	"context"
	"errors"
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	
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
func (a *Handler) getOwnSession(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	
	cCtx := c.(*_auth.ClaimsContext)
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
func (a *Handler) getSession(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	
	cCtx := c.(*_auth.ClaimsContext)
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
func (a *Handler) dropOwnSession(c echo.Context) error {
	var req _model.DropSession
	
	if err := c.Bind(&req); err != nil {
		return _api.Failed(c, http.StatusBadRequest, _api.FromErrorCode(_message.IncorrectFormat))
	}
	
	if okBind, err := _api.IsRequestValid(&req); !okBind {
		return _api.Failed(c, http.StatusBadRequest, err)
	}
	
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	
	cCtx := c.(*_auth.ClaimsContext)
	claims, _, err, authStatus := cCtx.Claims()
	if err != nil {
		logrus.Error(err)
		return _api.Failed(c, authStatus, err)
	}
	
	req.RequestInfo.IpNumber = c.RealIP()
	req.RequestInfo.UserAgent = c.Request().UserAgent()
	req.RequestInfo.Host = c.Request().Host
	
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
func (a *Handler) dropSession(c echo.Context) error {
	var req _model.DropSession
	
	if err := c.Bind(&req); err != nil {
		return _api.Failed(c, http.StatusBadRequest, _api.FromErrorCode(_message.IncorrectFormat))
	}
	
	if okBind, err := _api.IsRequestValid(&req); !okBind {
		return _api.Failed(c, http.StatusBadRequest, err)
	}
	
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	
	cCtx := c.(*_auth.ClaimsContext)
	claims, _, err, authStatus := cCtx.Claims()
	if err != nil {
		logrus.Error(err)
		return _api.Failed(c, authStatus, err)
	}
	
	req.RequestInfo.IpNumber = c.RealIP()
	req.RequestInfo.UserAgent = c.Request().UserAgent()
	req.RequestInfo.Host = c.Request().Host
	
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
