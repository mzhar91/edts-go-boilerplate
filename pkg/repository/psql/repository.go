package psql

import (
	_credential "sg-edts.com/edts-go-boilerplate/data/psql/credential"
	_session "sg-edts.com/edts-go-boilerplate/data/psql/session"
)

type Repository struct {
	Credential _credential.PsqlRepository
	Session    _session.PsqlRepository
}
