package localization

type Language string

const (
	Default Language = English
	English Language = "en"
)

var SupportedLanguages = map[Language]struct{}{
	English: {},
}
