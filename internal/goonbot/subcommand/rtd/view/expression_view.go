package view

import (
	"fmt"
	"goonbot/internal/goonbot/localization"
	"goonbot/internal/goonbot/subcommand/rtd/domain"
)

func ExpressionToView(c domain.Config, exp domain.ExpressionResult) (*domain.ExpressionResultView, error) {
	res := domain.ExpressionResultView{
		Title: localization.KeyToText(localization.RtdSuccessTitle, c.Language),
		Body:  "",
	}

	totalFmt := localization.KeyToText(localization.RollTotal, c.Language)
	res.Body += fmt.Sprintf(totalFmt, exp.Total)

	if !c.Verbose {
		return &res, nil
	}

	for _, roll := range exp.RollResults {
		rollView, err := rollToViewStr(c, roll)
		if err != nil {
			return nil, fmt.Errorf("error converting roll to view: %w", err)
		}

		res.Body += rollView
	}

	return &res, nil
}
