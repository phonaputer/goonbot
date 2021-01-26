package domain

import "goonbot/internal/goonbot/localization"

type Config struct {
	Flags
	Language localization.Language
}
