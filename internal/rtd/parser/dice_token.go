package parser

import (
	"math/rand"
	"time"
)

type DiceToken struct {
	NumDie    int
	NumFaces  int
	Operation Operation
}

func (r DiceToken) Apply(currentTotal int) int {
	return ExecOperation(currentTotal, r.Operation, rollDice(r.NumDie, r.NumFaces))
}

func rollDice(numDie, numFaces int) int {
	rand.Seed(time.Now().UnixNano())

	sum := 0
	for i := 0; i < numDie; i++ {
		sum += rand.Intn(numFaces) + 1
	}

	return sum
}
