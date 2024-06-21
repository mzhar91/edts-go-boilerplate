package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	_credentialRepo "sg-edts.com/edts-go-boilerplate/repository/credential/psql"
	_sessionRepo "sg-edts.com/edts-go-boilerplate/repository/session/psql"

	_config "sg-edts.com/edts-go-boilerplate/config"
	_apiHelper "sg-edts.com/edts-go-boilerplate/helper/api"
	_model "sg-edts.com/edts-go-boilerplate/model"
	_api "sg-edts.com/edts-go-boilerplate/pkg/api"
	_auth "sg-edts.com/edts-go-boilerplate/pkg/auth"
	_repository "sg-edts.com/edts-go-boilerplate/pkg/repository"
	_psql "sg-edts.com/edts-go-boilerplate/pkg/repository/psql"
	_session "sg-edts.com/edts-go-boilerplate/usecase/session"
)

type ucase struct {
	credentialRepo _credentialRepo.PsqlRepository
	sessionRepo    _sessionRepo.PsqlRepository
	contextTimeout time.Duration
	dbConn         *sql.DB
	debug          bool
	apiLibs        _apiHelper.Libs
}

func NewUcase(psql *_psql.Repository, connection *_config.Connection, timeout time.Duration, apiLibs _apiHelper.Libs) _session.Usecase {
	return &ucase{
		sessionRepo:    psql.Session,
		dbConn:         connection.Database,
		contextTimeout: timeout,
		debug:          _config.Cfg.Debug,
		apiLibs:        apiLibs,
	}
}

func (a *ucase) GetOwnSession(ctx context.Context, claims *_auth.Claims) ([]*_model.SessionResponse, error) {
	list := make([]*_model.SessionResponse, 0)
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	conn := &_repository.Use{
		Db: a.dbConn,
	}
	paramList := []*_model.ParamSession{
		{
			Key:   "username",
			Value: claims.Username,
		},
	}

	sessionList, err := a.sessionRepo.ReadBy(
		ctx, conn, paramList,
	)
	if err != nil {
		log.Printf(err.Error())

		return nil, err
	}

	for _, value := range sessionList {
		list = append(
			list, &_model.SessionResponse{
				ID:             value.ID,
				Username:       value.Username,
				AccessToken:    value.AccessToken,
				RefreshToken:   value.RefreshToken,
				Scope:          value.Scope,
				DeviceID:       value.DeviceID,
				Ip:             value.Ip,
				Host:           value.Host,
				UserAgent:      value.UserAgent,
				CreatedAt:      value.CreatedAt,
				CreatedBy:      value.CreatedBy,
				LastModifiedBy: value.LastModifiedBy,
				LastModifiedAt: value.LastModifiedAt,
			},
		)
	}

	return list, nil
}

func (a *ucase) GetSession(ctx context.Context, param *_model.QuerySession) ([]*_model.SessionResponse, int, error) {
	list := make([]*_model.SessionResponse, 0)
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	conn := &_repository.Use{
		Db: a.dbConn,
	}

	sessionList, err := a.sessionRepo.Read(
		ctx, conn, param,
	)
	if err != nil {
		log.Printf(err.Error())

		return nil, 0, err
	}

	total, err := a.sessionRepo.Total(ctx, conn, param)
	if err != nil {
		log.Printf(err.Error())

		return nil, 0, err
	}

	if param.Page == 0 {
		param.Page = 1
	}
	if param.Limit == 0 {
		param.Limit = int64(total)
	}

	for _, value := range sessionList {
		list = append(
			list, &_model.SessionResponse{
				ID:             value.ID,
				Username:       value.Username,
				AccessToken:    value.AccessToken,
				RefreshToken:   value.RefreshToken,
				Scope:          value.Scope,
				DeviceID:       value.DeviceID,
				Ip:             value.Ip,
				Host:           value.Host,
				UserAgent:      value.UserAgent,
				CreatedAt:      value.CreatedAt,
				CreatedBy:      value.CreatedBy,
				LastModifiedBy: value.LastModifiedBy,
				LastModifiedAt: value.LastModifiedAt,
			},
		)
	}

	return list, total, nil
}

func (a *ucase) DropOwnSession(ctx context.Context, claims *_auth.Claims, param *_model.DropSession) (res string, err error) {
	var result string

	err, code := _repository.WithTransaction(
		a.dbConn, func(tx _repository.Transaction) (error, int) {
			conn := &_repository.Use{
				Trans: tx,
			}
			paramList := []*_model.ParamSession{
				{
					Key:   "username",
					Value: claims.Username,
				},
				{
					Key:   "id",
					Value: param.SessionID,
				},
			}

			// get session
			// start
			session, err := a.sessionRepo.ReadBy(
				ctx, conn, paramList,
			)
			if err != nil {
				logrus.Error(err)
				return err, http.StatusInternalServerError
			}
			// get session
			// end

			for _, loopSession := range session {
				// delete session
				// start
				err := a.sessionRepo.Delete(
					ctx, conn, loopSession.ID,
				)
				if err != nil {
					logrus.Error(err)
					return err, http.StatusInternalServerError
				}
				// delete session
				// end
			}

			return nil, http.StatusOK
		},
	)

	if err != nil {
		if a.debug || code == http.StatusUnprocessableEntity {
			return result, _api.WithMessage(
				0, fmt.Sprintf("Remove Order from Account failed caused: %v", err.Error()), code,
			)
		}

		return result, _api.WithMessage(0, "Remove Order from Account failed", code)
	}

	return result, nil
}

func (a *ucase) DropSession(ctx context.Context, claims *_auth.Claims, param *_model.DropSession) (res string, err error) {
	var result string

	err, code := _repository.WithTransaction(
		a.dbConn, func(tx _repository.Transaction) (error, int) {
			conn := &_repository.Use{
				Trans: tx,
			}
			paramList := []*_model.ParamSession{
				{
					Key:   "id",
					Value: param.SessionID,
				},
			}

			// get session
			// start
			session, err := a.sessionRepo.ReadBy(
				ctx, conn, paramList,
			)
			if err != nil {
				logrus.Error(err)
				return err, http.StatusInternalServerError
			}
			// get session
			// end

			for _, loopSession := range session {
				// delete session
				// start
				err := a.sessionRepo.Delete(
					ctx, conn, loopSession.ID,
				)
				if err != nil {
					logrus.Error(err)
					return err, http.StatusInternalServerError
				}
				// delete session
				// end
			}

			return nil, http.StatusOK
		},
	)

	if err != nil {
		if a.debug || code == http.StatusUnprocessableEntity {
			return result, _api.WithMessage(
				0, fmt.Sprintf("Remove Order from Account failed caused: %v", err.Error()), code,
			)
		}

		return result, _api.WithMessage(0, "Remove Order from Account failed", code)
	}

	return result, nil
}
