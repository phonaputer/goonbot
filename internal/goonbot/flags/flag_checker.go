package flags

import "goonbot/internal/goonbot/localization"

const (
	languageFlag = "--lang"
)

type flagChecker struct {
	aFlagWasMatched bool
	err             error
}

func newFlagChecker() *flagChecker {
	return &flagChecker{
		aFlagWasMatched: false,
		err:             nil,
	}
}

func (f *flagChecker) shouldShortCircuit() bool {
	return f.aFlagWasMatched || f.err != nil
}

func (f *flagChecker) language(token string, gf *GlobalFlags) {
	if f.shouldShortCircuit() {
		return
	}

	isLangFlag, val := checkFlagForValue(token, languageFlag)
	if !isLangFlag {
		return
	}

	f.aFlagWasMatched = true
	valLang := localization.Language(val)

	_, isSupported := localization.SupportedLanguages[valLang]
	if !isSupported {
		f.err = localization.NewWithUserMsg("unsupported language", localization.ErrUnsupportedLanguage)
		return
	}

	gf.Language = valLang
}

func checkFlagForValue(token string, f string) (isFlag bool, value string) {
	flagPrefix := f + "="
	lenFlagPrefix := len(flagPrefix)

	if len(token) < lenFlagPrefix || token[:lenFlagPrefix] != flagPrefix {
		return false, ""
	}

	if len(token) == lenFlagPrefix {
		return true, ""
	}

	return true, token[lenFlagPrefix:]
}
