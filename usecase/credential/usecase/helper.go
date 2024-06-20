package usecase

import (
	"context"
	
	_model "sg-edts.com/edts-go-boilerplate/model"
	_repository "sg-edts.com/edts-go-boilerplate/pkg/repository"
)

const (
	_CredentialFailedMsg             string = "sign in failed"
	_NotFoundMsg                     string = "user not found"
	_AccountLockedMsg                string = "account is locked"
	_CredentialSucceedMsg            string = "sign in succeed"
	_AddCredentialFailedMsg          string = "add credential failed"
	_AddCredentialSucceedMsg         string = "add credential succeed"
	_SignOutFailedMsg                string = "sign out failed"
	_SignOutSucceedMsg               string = "sign out succeed"
	_AccountBlockedMsg               string = "account blocked"
	_FailedAttempMsg                 string = "failed attempt"
	_ResetPasswordFailedMsg          string = "failed to reset password"
	_ResetPasswordSucceedMsg         string = "reset password email sent"
	_ConfirmResetPasswordFailedMsg   string = "confirmation reset password failed"
	_ConfirmResetPasswordSucceedMsg  string = "confirmation reset password succeed"
	_ValidateResetPasswordFailedMsg  string = "reset password not allowed"
	_ValidateResetPasswordSucceedMsg string = "reset password allowed"
)

func (a ucase) getInfo(ctx context.Context, conn *_repository.Use, username string) (*_model.Credential, error, int) {
	// get by username
	// cred, err := a.credentialRepo.GetByUsername(
	// 	ctx, conn, username,
	// )
	// if err != nil {
	// 	logrus.Errorf("getInfo %v", err)
	//
	// 	return nil, fmt.Errorf(_NotFoundMsg), http.StatusNotFound
	// }
	//
	// if cred.IsLock == true {
	// 	return cred, fmt.Errorf(_AccountLockedMsg), http.StatusLocked
	// }
	//
	// return cred, nil, 0
	return nil, nil, 0
}

func (a ucase) attemptLoginFailed(ctx context.Context, param *_model.SignInRequest) (error, int, int) {
	// counter := 0
	//
	// // transaction
	// err, code := _repository.WithTransaction(
	// 	a.dbConn, func(tx _repository.Transaction) (error, int) {
	// 		conn := &_repository.Use{
	// 			Trans: tx,
	// 		}
	//
	// 		// get by username
	// 		cred, err := a.credentialRepo.GetByUsername(
	// 			ctx, conn, param.Username,
	// 		)
	// 		if err != nil {
	// 			logrus.Error(err)
	// 			return err, http.StatusInternalServerError
	// 		}
	//
	// 		attempt := int(cred.FailedAttempt.Int64 + 1)
	// 		err = a.credentialRepo.Update(
	// 			ctx, conn, cred.ID, &_model.Credential{
	// 				FailedAttempt: null.IntFrom(int64(attempt)),
	// 			},
	// 		)
	// 		if err != nil {
	// 			logrus.Error(err)
	// 			return err, http.StatusInternalServerError
	// 		}
	//
	// 		counter = attempt
	//
	// 		return nil, http.StatusOK
	// 	},
	// )
	// if err != nil {
	// 	return err, code, counter
	// }
	//
	// return nil, code, counter
	return nil, 0, 0
}
