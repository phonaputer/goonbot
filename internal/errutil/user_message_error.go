package errutil

import "errors"

type userMsgErr struct {
	Err     error
	UserMsg string
}

func (u *userMsgErr) Error() string {
	return u.Err.Error()
}

func (u *userMsgErr) Unwrap() error {
	return u.Err
}

func NewWithUserMsg(msg string) error {
	return &userMsgErr{
		Err:     errors.New(msg),
		UserMsg: msg,
	}
}

func WithUserMsg(err error, msg string) error {
	return &userMsgErr{
		Err:     err,
		UserMsg: msg,
	}
}

func GetUserMsg(err error) string {
	var uMsgErr *userMsgErr
	if errors.As(err, &uMsgErr) {
		return uMsgErr.UserMsg
	}

	return "unknown error occurred"
}
