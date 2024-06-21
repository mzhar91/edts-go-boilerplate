package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	_credentialRepo "sg-edts.com/edts-go-boilerplate/repository/credential/psql"
	_sessionRepo "sg-edts.com/edts-go-boilerplate/repository/session/psql"

	_config "sg-edts.com/edts-go-boilerplate/config"
	_constant "sg-edts.com/edts-go-boilerplate/constant"
	_apiHelper "sg-edts.com/edts-go-boilerplate/helper/api"
	_model "sg-edts.com/edts-go-boilerplate/model"
	_api "sg-edts.com/edts-go-boilerplate/pkg/api"
	_auth "sg-edts.com/edts-go-boilerplate/pkg/auth"
	_repository "sg-edts.com/edts-go-boilerplate/pkg/repository"
	_psql "sg-edts.com/edts-go-boilerplate/pkg/repository/psql"
	_security "sg-edts.com/edts-go-boilerplate/pkg/security"
	_credential "sg-edts.com/edts-go-boilerplate/usecase/credential"
)

type ucase struct {
	credentialRepo _credentialRepo.PsqlRepository
	sessionRepo    _sessionRepo.PsqlRepository
	contextTimeout time.Duration
	dbConn         *sql.DB
	debug          bool
	apiLibs        _apiHelper.Libs
}

func NewUcase(psql *_psql.Repository, connection *_config.Connection, timeout time.Duration, apiLibs _apiHelper.Libs) _credential.Usecase {
	return &ucase{
		credentialRepo: psql.Credential,
		sessionRepo:    psql.Session,
		dbConn:         connection.Database,
		contextTimeout: timeout,
		debug:          _config.Cfg.Debug,
		apiLibs:        apiLibs,
	}
}

func (a *ucase) RefreshToken(ctx context.Context, app string, cCtx *_auth.ClaimsContext) (*_model.SignInResponse, error, map[string]interface{}) {
	response := new(_model.SignInResponse)
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	// transaction
	// start
	err, code := _repository.WithTransaction(
		a.dbConn, func(tx _repository.Transaction) (error, int) {
			var duration int64

			conn := &_repository.Use{
				Trans: tx,
			}

			claims, _, err, _ := cCtx.Claims()
			if err != nil {
				logrus.Errorf("RefreshToken %v", err)

				return err, http.StatusUnauthorized
			}

			// get by username
			// start
			cred, err, errCode := a.getInfo(ctx, conn, claims.Username)
			if err != nil {
				logrus.Errorf("RefreshToken %v", err)

				return err, errCode
			}
			// get by username
			// end

			// generate temporary access token
			// start
			accessToken, err, _ := cCtx.GenerateAccessToken(
				cred.UserID.String(),
				cred.Username,
				app,
			)
			if err != nil {
				logrus.Errorf("RefreshToken %v", err)

				return err, http.StatusInternalServerError
			}
			// generate temporary access token
			// end

			if cred.Username != _constant.SystemUsername {
				// send email
				// start
				// _, err := a.apiLibs.Notification.SendEmail()
				// if err != nil {
				// 	logrus.Errorf("RefreshToken %v", err)
				//
				// 	return err, http.StatusInternalServerError
				// }
				// send email
				// end

				// generate access token
				// start
				accessToken, err, _ = cCtx.GenerateAccessToken(
					cred.UserID.String(),
					cred.Username,
					app,
				)
				if err != nil {
					logrus.Errorf("RefreshToken %v", err)

					return err, http.StatusInternalServerError
				}
				// generate access token
				// end
			}

			// generate refresh token
			// start
			refreshToken, err, _ := cCtx.GenerateRefreshToken(cred.UserID.String(), cred.Username, app)
			if err != nil {
				logrus.Errorf("RefreshToken %v", err)

				return err, http.StatusInternalServerError
			}
			// generate refresh token
			// end

			if app == "mobile" {
				duration = _config.Cfg.Jwt.RefreshPeriodMobile
			} else if app == "bo" {
				duration = _config.Cfg.Jwt.RefreshPeriodBo
			} else {
				return fmt.Errorf("app was not declared"), http.StatusInternalServerError
			}

			response.AccessToken = accessToken
			response.RefreshToken = refreshToken
			response.TokenExpiration = time.Now().Add(time.Duration(duration) * time.Minute).Unix()

			return nil, http.StatusOK
		},
	)
	if err != nil {
		if a.debug {
			return nil, _api.WithMessage(code, err.Error(), code), nil
		}

		return nil, _api.WithMessage(code, _CredentialFailedMsg, code), nil
	}
	// transaction
	// end

	return response, nil, map[string]interface{}{"message": _CredentialSucceedMsg, "code": http.StatusOK}
}

func (a *ucase) SignIn(ctx context.Context, app string, cCtx *_auth.ClaimsContext, req *_model.SignInRequest) (*_model.SignInResponse, error, map[string]interface{}) {
	response := new(_model.SignInResponse)
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	req.Username = strings.TrimSpace(strings.ToLower(req.Username))

	logrus.Infof("SignIn ucase %v", req)

	// transaction
	// start
	err, code := _repository.WithTransaction(
		a.dbConn, func(tx _repository.Transaction) (error, int) {
			var duration int64

			conn := &_repository.Use{
				Trans: tx,
			}

			// verify password
			// start
			salt := _security.GenerateSalt(req.Username)
			hash := _security.Hash(req.Password, []byte(salt))
			cred, err := a.credentialRepo.VerifyPassword(
				ctx, conn, req.Username, hash,
			)
			if err != nil {
				logrus.Errorf("SignIn %v", err)

				return err, http.StatusUnauthorized
			}
			// verify password
			// end

			// update logged in
			// start
			err = a.credentialRepo.Update(
				ctx, conn, cred.ID, &_model.CredentialSetLogIn{
					LastLogin:    time.Now().String(),
					ModifiedDate: time.Now().String(),
				},
			)
			if err != nil {
				logrus.Errorf("SignIn %v", err)

				return err, http.StatusInternalServerError
			}
			// update logged in
			// end

			// generate temporary access token
			// start
			accessToken, err, _ := cCtx.GenerateAccessToken(
				cred.UserID.String(),
				cred.Username,
				app,
			)
			if err != nil {
				logrus.Errorf("SignIn %v", err)

				return err, http.StatusInternalServerError
			}
			// generate temporary access token
			// end

			if cred.Username != _constant.SystemUsername {
				// generate access token
				// start
				accessToken, err, _ = cCtx.GenerateAccessToken(
					cred.UserID.String(),
					cred.Username,
					app,
				)
				if err != nil {
					logrus.Errorf("SignIn %v", err)

					return err, http.StatusInternalServerError
				}
				// generate access token
				// end
			}

			// generate refresh token
			// start
			refreshToken, err, _ := cCtx.GenerateRefreshToken(cred.UserID.String(), cred.Username, app)
			if err != nil {
				logrus.Errorf("SignIn %v", err)

				return err, http.StatusInternalServerError
			}
			// generate refresh token
			// end

			// create session
			// start
			err = a.sessionRepo.Create(
				ctx, conn, &_model.SessionSetLogin{
					Username:     req.Username,
					Ip:           req.RequestInfo.IpNumber,
					UserAgent:    req.RequestInfo.UserAgent,
					Host:         req.RequestInfo.Host,
					DeviceID:     req.DeviceId,
					Scope:        app,
					AccessToken:  accessToken,
					RefreshToken: refreshToken,
					CreatedAt:    time.Now().Unix(),
					CreatedBy:    req.Username,
				},
			)
			if err != nil {
				logrus.Errorf("SignIn %v", err)

				return err, http.StatusInternalServerError
			}
			// create session
			// end

			if app == "mobile" {
				duration = _config.Cfg.Jwt.AccessPeriodMobile
			} else if app == "bo" {
				duration = _config.Cfg.Jwt.AccessPeriodBo
			} else {
				return fmt.Errorf("app was not declared"), http.StatusInternalServerError
			}

			response.AccessToken = accessToken
			response.RefreshToken = refreshToken
			response.TokenExpiration = time.Now().Add(time.Duration(duration) * time.Minute).Unix()

			return nil, http.StatusOK
		},
	)
	if err != nil {
		// if error code unauthorized
		if code == http.StatusUnauthorized {
			err, code, attempt := a.attemptLoginFailed(ctx, req)
			if err == nil {
				return nil, _api.WithMessage(code, fmt.Sprintf("%v, %v", _FailedAttempMsg, attempt), code), nil
			}
		}

		if a.debug {
			return nil, _api.WithMessage(code, err.Error(), code), nil
		}

		return nil, _api.WithMessage(code, _CredentialFailedMsg, code), nil
	}
	// transaction
	// end

	return response, nil, map[string]interface{}{"message": _CredentialSucceedMsg, "code": http.StatusOK}
}

func (a *ucase) SignOut(ctx context.Context, app string, username string, req *_model.SignOutRequest) (error, map[string]interface{}) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	// transaction
	// start
	err, code := _repository.WithTransaction(
		a.dbConn, func(tx _repository.Transaction) (error, int) {
			conn := &_repository.Use{
				Trans: tx,
			}
			paramList := []*_model.ParamSession{
				{
					Key:   "username",
					Value: username,
				},
			}

			// get by username
			// start
			cred, err := a.credentialRepo.GetByUsername(
				ctx, conn, username,
			)
			if err != nil {
				logrus.Error(err)
				return err, http.StatusInternalServerError
			}
			// get by username
			// end

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

			if len(session) == 1 {
				err = a.credentialRepo.Update(
					ctx, conn, cred.ID, &_model.CredentialSetLogOut{
						ModifiedDate: time.Now().String(),
					},
				)
				if err != nil {
					logrus.Error(err)
					return err, http.StatusInternalServerError
				}
			}

			for _, loopSession := range session {
				if loopSession.DeviceID == req.DeviceId {
					// delete session based on device id
					// start
					err := a.sessionRepo.Delete(
						ctx, conn, loopSession.ID,
					)
					if err != nil {
						logrus.Error(err)
						return err, http.StatusInternalServerError
					}
					// delete session based on device id
					// end
				}
			}

			return nil, http.StatusOK
		},
	)
	if err != nil {
		if a.debug {
			return _api.WithMessage(code, err.Error(), code), nil
		}

		return _api.WithMessage(code, _SignOutFailedMsg, code), nil
	}
	// transaction
	// end

	return nil, map[string]interface{}{"message": _SignOutSucceedMsg, "code": http.StatusOK}
}

func (a *ucase) AddCredential(ctx context.Context, app string, claims *_auth.Claims, req *_model.AddCredentialRequest) (*_model.AddCredentialResponse, error, map[string]interface{}) {
	response := new(_model.AddCredentialResponse)
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	logrus.Infof("AddCredential ucase %v", req)

	// transaction
	// start
	err, code := _repository.WithTransaction(
		a.dbConn, func(tx _repository.Transaction) (error, int) {
			conn := &_repository.Use{
				Trans: tx,
			}

			// generate password
			// start
			salt := _security.GenerateSalt(req.Username)
			hash := _security.Hash(req.Password, salt)
			// generate password
			// end

			// update confirm reset password
			// start
			err := a.credentialRepo.Create(
				ctx, conn, &_model.Credential{
					Username:     req.Username,
					PasswordHash: hash,
					CreatedDate:  time.Now().String(),
				},
			)
			if err != nil {
				logrus.Errorf("AddCredential %v", err)

				return err, http.StatusInternalServerError
			}
			// update confirm reset password
			// end

			response.Username = req.Username

			return nil, http.StatusOK
		},
	)
	if err != nil {
		if a.debug {
			return nil, _api.WithMessage(code, err.Error(), code), nil
		}

		return nil, _api.WithMessage(code, _AddCredentialFailedMsg, code), nil
	}
	// transaction
	// end

	return response, nil, map[string]interface{}{"message": _AddCredentialSucceedMsg, "code": http.StatusOK}
}
