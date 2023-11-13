package number_speller

type Interface interface {
	Locale() string
	Translate(x int) string
}

var Languages map[string]Interface
var DefaultLanguage Interface

func init() {
	DefaultLanguage = enNumberSpeller

	Languages = map[string]Interface{}
	Languages[enNumberSpeller.Locale()] = enNumberSpeller
	Languages[idNumberSpeller.Locale()] = idNumberSpeller
}

type numberSpeller struct {
	locale     string
	translator func(x int) string
}

func (ns *numberSpeller) Locale() string {
	return ns.locale
}

func (ns *numberSpeller) Translate(x int) string {
	return ns.translator(x)
}

func newNumberSpeller(locale string, translator func(x int) string) *numberSpeller {
	return &numberSpeller{
		locale:     locale,
		translator: translator,
	}
}
