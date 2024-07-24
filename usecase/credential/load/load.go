package load

import (
	"time"

	"github.com/gofiber/fiber/v2"
	_http "sg-edts.com/edts-go-boilerplate/handler/auth/http"

	_notifApi "sg-edts.com/edts-go-boilerplate/helper/api/notification"

	_config "sg-edts.com/edts-go-boilerplate/config"
	_api "sg-edts.com/edts-go-boilerplate/helper/api"
	_psql "sg-edts.com/edts-go-boilerplate/pkg/repository/psql"
	_credentialPsql "sg-edts.com/edts-go-boilerplate/repository/credential/psql"
	_sessionPsql "sg-edts.com/edts-go-boilerplate/repository/session/psql"
	_usecase "sg-edts.com/edts-go-boilerplate/usecase/credential/usecase"
)

func Load(e *fiber.App, connection *_config.Connection, timeoutContext time.Duration) {
	repo := &_psql.Repository{
		Credential: _credentialPsql.NewPsqlRepository(),
		Session:    _sessionPsql.NewPsqlRepository(),
	}
	apiLib := _api.Libs{
		Notification: _notifApi.NewNotification(
			_config.Cfg.Services.KlikPromo.Url,
			_config.Cfg.Services.KlikPromo.Header,
			_config.Cfg.Debug,
		),
	}

	ucase := _usecase.NewUcase(repo, connection, timeoutContext, apiLib)

	_http.NewHandler(e, ucase)
}
