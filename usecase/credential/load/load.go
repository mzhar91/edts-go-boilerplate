package load

import (
	"time"

	"github.com/gofiber/fiber/v2"
	_http "sg-edts.com/edts-go-boilerplate/handler/auth/http"
	_userApi "sg-edts.com/edts-go-boilerplate/helper/api/user"

	_notifApi "sg-edts.com/edts-go-boilerplate/helper/api/notification"

	_config "sg-edts.com/edts-go-boilerplate/config"
	_credentialPsql "sg-edts.com/edts-go-boilerplate/handler/auth/repository/psql"
	_usecase "sg-edts.com/edts-go-boilerplate/handler/auth/usecase"
	_sessionPsql "sg-edts.com/edts-go-boilerplate/handler/session/repository/psql"
	_api "sg-edts.com/edts-go-boilerplate/helper/api"
	_psql "sg-edts.com/edts-go-boilerplate/helper/repository/psql"
)

func Load(e *fiber.Ctx, connection *_config.Connection, timeoutContext time.Duration) {
	repo := &_psql.Repository{
		Credential: _credentialPsql.NewPsqlRepository(),
		Session:    _sessionPsql.NewPsqlRepository(),
	}
	apiLib := _api.Libs{
		Notification: _notifApi.NewNotification(
			_config.Cfg.ServiceNotifURL,
			_config.Cfg.ServiceNotifAuthHeader,
			_config.Cfg.Debug,
		),
		User: _userApi.NewUser(
			_config.Cfg.ServiceUserURL,
			_config.Cfg.ServiceUserAuthHeader,
			_config.Cfg.Debug,
		),
	}

	ucase := _usecase.NewUcase(repo, connection, timeoutContext, apiLib)

	_http.NewHandler(e, ucase)
}
