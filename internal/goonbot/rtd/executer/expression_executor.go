package executer

import (
	"fmt"
	"goonbot/internal/goonbot/rtd/domain"
)

func ExecuteExpression(e domain.Expression) (domain.ExpressionResult, error) {
	var res domain.ExpressionResult
	for _, roll := range e.Rolls {
		addedRes, err := execOneRollAndAddToResult(roll, res)
		if err != nil {
			return res, fmt.Errorf("error executing expression: %w", err)
		}

		res = addedRes
	}

	return res, nil
}

func execOneRollAndAddToResult(roll domain.Roll, expRes domain.ExpressionResult) (domain.ExpressionResult, error) {
	rollRes, err := determineRollTypeAndExec(roll)
	if err != nil {
		return expRes, fmt.Errorf("error executing roll: %w", err)
	}

	newTotal, err := execOperation(expRes.Total, rollRes.GetOperation(), rollRes.GetTotal())
	if err != nil {
		return expRes, fmt.Errorf("error executing operation: %w", err)
	}

	return domain.ExpressionResult{
		Total:       newTotal,
		RollResults: append(expRes.RollResults, rollRes),
	}, nil
}
