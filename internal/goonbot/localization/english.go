package localization

var englishMap = map[Key]string{
	ErrUnknownCommandTitle: "Command not recognized",
	ErrUnknownErr:          "unexpected error occurred",
	ErrUnsupportedLanguage: "language not supported",

	// Error
	ErrorTitle:     "Error!",
	ErrorFieldName: "ya done goofed...",

	//Help
	HelpTitle:             "Help",
	HelpAvailableCommands: "Available selections:",
	HelpCommandDetails: `== FLAGS ==
	--lang : Language selection.
		Example Usage: --lang=en

== COMMANDS ==
	help : Stop it. Get some help.
		Example Usage: !goonbot help

	rtd : Roll dice. Papa needs a new pair of shoes!
		Example Usage: !goonbot rtd 1d6 + 10
		Supported Flags:
			--v : Verbose output`,

	// RTD
	ErrEmptyInput:        "input must not be empty",
	ErrMathSeparated:     "input must be a list of +/- separated rolls",
	ErrInvalidDiceFmt:    "invalid dice roll format",
	ErrInvalidNumDie:     "dice roll xdy: x must be > 0",
	ErrInvalidNumFaces:   "dice roll xdy: y must be > 1",
	ErrUnknownArithmetic: "unknown arithmetic operation",
	RollTotal:            "total: %v",
	DiceRollResult:       "\n%s %v\t=>\t %vd%v: %s",
	SimpleRollResult:     "\n%s %v",
	RtdSuccessTitle:      "you rolled:",
}
