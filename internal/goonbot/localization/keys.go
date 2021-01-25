package localization

// Key is a localization key.
// A key is associated with a specific message and can be used to retrieve that message localized in any supported
// language.
type Key int

const (
	UnknownKey Key = iota

	ErrUnknownErr

	// RTD
	ErrEmptyInput
	ErrMathSeparated
	ErrInvalidDiceFmt
	ErrInvalidNumDie
	ErrInvalidNumFaces
	ErrUnknownArithmetic
	YouRolled
	Plus
	Minus
)
