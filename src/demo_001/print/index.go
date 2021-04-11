package print

const JAPANESE = "JAPANESE"
const CHINESE = "CHINESE"

func PrintStr(name, language string) string {
	if name == "" {
		name = "world"
	}
	return getLanguageHello(language) + name
}

func getLanguageHello(languageType string) string {
	switch languageType {
	case CHINESE:
		return "你好！"
	case JAPANESE:
		return "こんにちは～"
	default:
		return "Hello "
	}
}
