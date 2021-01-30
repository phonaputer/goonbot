package localization

const rtdErrSuffix = "\n\nget help: !goonbot help rtd"

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
		Example Usage: !goonbot help rtd

	rtd : Roll dice. Baby needs a new pair of shoes!
		Example Usage: !goonbot rtd 1d6 + 10
		Supported Flags:
			--v : Verbose output`,

	// RTD
	ErrEmptyInput:        "input must not be empty" + rtdErrSuffix,
	ErrMathSeparated:     "input must be a list of +/- separated rolls" + rtdErrSuffix,
	ErrInvalidDiceFmt:    "invalid dice roll format" + rtdErrSuffix,
	ErrInvalidNumDie:     "dice roll xdy: x must be > 0" + rtdErrSuffix,
	ErrInvalidNumFaces:   "dice roll xdy: y must be > 1" + rtdErrSuffix,
	ErrUnknownArithmetic: "unknown arithmetic operation" + rtdErrSuffix,
	RollTotal:            "total: %v",
	DiceRollResult:       "\n%s %v\t=>\t %vd%v: %s",
	SimpleRollResult:     "\n%s %v",
	RtdSuccessTitle:      "you rolled:",
	HelpRtdTitle:         "!goonbot rtd : Roll some dice",
	HelpRtdDetails: `!goonbot rtd executes DnD style dice rolls.

RTD requires input in the form of a list of rolls separated by arithmetic operators, together forming a "roll expression." The expression is evaluated and resulting total returned.

Some example expressions:
	* Roll 2 6-sided die: 2d6
	* Roll 2 6-sided die then roll 1 20-sided die and add the results: 2d6 + 1d20
	* Roll 2 6-sided die then roll 1 20-sided die and sum the results. Subtract 5 from the total: 2d6 + 1d20 - 5

Rolls may have one of the following formats:
	1. Dice roll (XdY). X = the number of die to roll, Y = the number of sides on each die.
	2. Fixed number (X). X must be an integer.

The following operations are supported:
	1. Addition: +
	2. Subtraction: -

Supported flags:
	1. Verbose: --v . Returns a detailed breakdown of what you rolled.`,
}
