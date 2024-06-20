package credential

import (
	"time"
	
	"github.com/labstack/echo/v4"
	
	_config "sg-edts.com/edts-go-boilerplate/config"
	_credentialPsql "sg-edts.com/edts-go-boilerplate/data/psql/credential"
	_sessionPsql "sg-edts.com/edts-go-boilerplate/data/psql/session"
	_http "sg-edts.com/edts-go-boilerplate/handler/auth/http"
	_apiHelper "sg-edts.com/edts-go-boilerplate/helper/api"
	_notifApi "sg-edts.com/edts-go-boilerplate/helper/api/notification"
	_psql "sg-edts.com/edts-go-boilerplate/pkg/repository/psql"
	_usecase "sg-edts.com/edts-go-boilerplate/usecase/credential/usecase"
)

func Load(e *echo.Echo, connection *_config.Connection, timeoutContext time.Duration) {
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
