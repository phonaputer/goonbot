package localization

var englishMap = map[Key]string{
	ErrUnknownErr: "unknown error occurred",

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
	YourRoll:             "your roll:",
}
