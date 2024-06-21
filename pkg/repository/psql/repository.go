package psql

import (
	_credential "sg-edts.com/edts-go-boilerplate/repository/credential/psql"
	_session "sg-edts.com/edts-go-boilerplate/repository/session/psql"
)

type Repository struct {
	Credential _credential.PsqlRepository
	Session    _session.PsqlRepository
}
