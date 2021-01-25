package localization

import "github.com/sirupsen/logrus"

func KeyToText(key Key, lang Language) string {
	var res string
	var ok bool

	switch lang {
	case English:
		res, ok = englishMap[key]
	default:
		res, ok = englishMap[key]
	}

	if !ok {
		logrus.Errorf("Missing localization: key: %v, lang: %v", key, lang)
		return ""
	}

	return res
}
