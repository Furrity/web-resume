package l10n

func IsSupported(code string) bool {
	switch code {
	case "ru":
		fallthrough
	case "en":
		return true
	default:
		return false
	}
}
