package executer

import (
	"errors"
	"github.com/sirupsen/logrus"
	"goonbot/internal/goonbot/subcommand/rtd/domain"
	"math/rand"
	"time"
)

func determineRollTypeAndExec(roll domain.Roll) (domain.RollResult, error) {
	switch typed := roll.(type) {
	case domain.DiceRoll:
		return execDiceRoll(typed), nil
	case domain.SimpleRoll:
		return execSimpleRoll(typed), nil
	}

	logrus.Errorf("unknown roll type in expression executor: %T", roll)

	return nil, errors.New("unknown roll type")
}

func execSimpleRoll(roll domain.SimpleRoll) domain.SimpleRollResult {
	return domain.SimpleRollResult{
		Operation: roll.Operation,
		Number:    roll.Number,
	}
}

func execDiceRoll(roll domain.DiceRoll) domain.DiceRollResult {
	total, rolledInts := rollDice(roll.NumDie, roll.NumFaces)

	return domain.DiceRollResult{
		Operation:     roll.Operation,
		NumDie:        roll.NumDie,
		NumFaces:      roll.NumFaces,
		RolledNumbers: rolledInts,
		TotalRolled:   total,
	}
}

func rollDice(numDie, numFaces int) (int, []int) {
	rand.Seed(time.Now().UnixNano())

	total := 0
	var rolledInts []int

	for i := 0; i < numDie; i++ {
		randRoll := rand.Intn(numFaces) + 1
		total += randRoll
		rolledInts = append(rolledInts, randRoll)
	}

	return total, rolledInts
}
