package localization

import "errors"

type localizedErr struct {
	Err    error
	LocKey Key
}

func (l *localizedErr) Error() string {
	return l.Err.Error()
}

func (l *localizedErr) Unwrap() error {
	return l.Err
}

func NewWithUserMsg(err string, key Key) error {
	return &localizedErr{
		Err:    errors.New(err),
		LocKey: key,
	}
}

func WithUserMsg(err error, key Key) error {
	return &localizedErr{
		Err:    err,
		LocKey: key,
	}
}

func ErrToText(err error, lang Language) string {
	return KeyToText(getErrLocKey(err), lang)
}

func getErrLocKey(err error) Key {
	var uMsgErr *localizedErr
	if errors.As(err, &uMsgErr) {
		return uMsgErr.LocKey
	}

	return ErrUnknownErr
}
