package credential

import (
	"time"

	"github.com/gofiber/fiber/v2"
	_credentialPsql "sg-edts.com/edts-go-boilerplate/repository/credential/psql"
	_sessionPsql "sg-edts.com/edts-go-boilerplate/repository/session/psql"

	_config "sg-edts.com/edts-go-boilerplate/config"
	_http "sg-edts.com/edts-go-boilerplate/handler/auth/http"
	_apiHelper "sg-edts.com/edts-go-boilerplate/helper/api"
	_notifApi "sg-edts.com/edts-go-boilerplate/helper/api/notification"
	_psql "sg-edts.com/edts-go-boilerplate/pkg/repository/psql"
	_usecase "sg-edts.com/edts-go-boilerplate/usecase/credential/usecase"
)

func Load(e *fiber.App, connection *_config.Connection, timeoutContext time.Duration) {
	repo := &_psql.Repository{
		Credential: _credentialPsql.NewPsqlRepository(),
		Session:    _sessionPsql.NewPsqlRepository(),
	}
	apiLib := _apiHelper.Libs{
		Notification: _notifApi.NewNotification(
			"",
			"",
			_config.Cfg.Debug,
		),
	}

	ucase := _usecase.NewUcase(repo, connection, timeoutContext, apiLib)

	_http.NewHandler(e, ucase)
}
