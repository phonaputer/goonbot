package rtd

import (
	"goonbot/internal/goonbot/localization"
	"goonbot/internal/goonbot/subcommand/rtd/domain"
	"goonbot/internal/goonbot/subcommand/rtd/executer"
	"goonbot/internal/goonbot/subcommand/rtd/parser"
	"goonbot/internal/goonbot/subcommand/rtd/view"
)

// Settings is a list of options related to the kind of result RollTheDice will return.
type Settings struct {

	// Language is the human language in which you want your result.
	Language localization.Language
}

// Result is a view model containing information about the outcome of your roll
type Result struct {
	Title string
	Body  string
}

// RollTheDice executes a DnD style dice roll using the provided input.
//
// The input parameter is expected to be a list of rolls separated by arithmetic operators, together forming a
// "roll expression." The expression is evaluated and resulting total returned.
//
// If the input arguments are malformed or otherwise incorrect, an error message will be returned explaining what went
// wrong.
//
// Some example expressions:
//		// roll 2 6-sided die
//		2d6
//
//		// roll 2 6-sided die then roll 1 20-sided die and add the results
//		2d6 + 1d20
//
//		// roll 2 6-sided die then roll 1 20-sided die and sum the results. Subtract 5 from the total.
//		2d6 + 1d20 - 5
//
// Rolls may have one of the following formats:
//		1. dice roll (XdY). X = the number of die to roll, Y = the number of sides on each die
// 		2. fixed number (X). X must be an integer.
//
// The following operations are supported:
//		1. Addition: +
// 		2. Subtraction: -
//
// Supported flags:
//		1. Verbose: --v . Returns a detail explanation of what you rolled.
func RollTheDice(s Settings, input []string) (*Result, error) {
	f, rollInput := parser.CheckFlags(input)
	c := domain.Config{Language: s.Language, Flags: f}

	exp, err := parser.Parse(rollInput)
	if err != nil {
		return nil, err
	}

	expRes, err := executer.ExecuteExpression(*exp)
	if err != nil {
		return nil, err
	}

	res, err := view.ExpressionToView(c, expRes)
	if err != nil {
		return nil, err
	}

	return &Result{Body: res.Body, Title: res.Title}, nil
}
