package executer

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type DiceRoll struct {
	NumDie    int
	NumFaces  int
	Operation Operation
}

func (r DiceRoll) Apply(current RollingResult) RollingResult {
	totalRolled, rollsStr := rollDice(r.NumDie, r.NumFaces)

	textLine := fmt.Sprintf("\n%s %v\t=>\t %vd%v: %s",
		GetOperationStr(r.Operation), totalRolled, r.NumDie, r.NumFaces, rollsStr)

	current.Sum = ExecOperation(current.Sum, r.Operation, totalRolled)
	current.Text += textLine

	return current
}

func rollDice(numDie, numFaces int) (int, string) {
	rand.Seed(time.Now().UnixNano())

	sum := 0
	var rolls []string

	for i := 0; i < numDie; i++ {
		randRoll := rand.Intn(numFaces) + 1
		sum += randRoll
		rolls = append(rolls, strconv.Itoa(randRoll))
	}

	return sum, strings.Join(rolls, ", ")
}
