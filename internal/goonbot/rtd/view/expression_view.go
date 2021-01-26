package view

import (
	"fmt"
	"goonbot/internal/goonbot/localization"
	"goonbot/internal/goonbot/rtd/domain"
)

func ExpressionToView(f domain.RTDFlags, exp domain.ExpressionResult) (string, error) {
	var resView string

	totalFmt := localization.KeyToText(localization.RollTotal, localization.English)
	resView += fmt.Sprintf(totalFmt, exp.Total)

	if !f.Verbose {
		return resView, nil
	}

	for _, roll := range exp.RollResults {
		rollView, err := rollToViewStr(roll)
		if err != nil {
			return "", fmt.Errorf("error converting roll to view: %w", err)
		}

		resView += rollView
	}

	return resView, nil
}
