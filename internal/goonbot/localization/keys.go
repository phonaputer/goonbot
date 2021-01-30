package localization

// Key is a localization key.
// A key is associated with a specific message and can be used to retrieve that message localized in any supported
// language.
type Key int

const (
	UnknownKey Key = iota

	ErrUnknownCommandTitle
	ErrUnknownErr
	ErrUnsupportedLanguage

	// Error
	ErrorTitle
	ErrorFieldName

	// Help
	HelpTitle
	HelpAvailableCommands
	HelpCommandDetails

	// RTD
	ErrEmptyInput
	ErrMathSeparated
	ErrInvalidDiceFmt
	ErrInvalidNumDie
	ErrInvalidNumFaces
	ErrUnknownArithmetic
	RollTotal
	DiceRollResult
	SimpleRollResult
	RtdSuccessTitle
	HelpRtdTitle
	HelpRtdDetails
)
