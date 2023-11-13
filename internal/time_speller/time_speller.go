package time_speller

type Interface interface {
	Locale() string
	Translate(x int) (string, error)
}

var Languages map[string]Interface
var DefaultLanguage Interface

func init() {
	DefaultLanguage = enTimeSpeller

	Languages = map[string]Interface{}
	Languages[enTimeSpeller.Locale()] = enTimeSpeller
	Languages[idTimeSpeller.Locale()] = idTimeSpeller
}

type timeSpeller struct {
	locale     string
	translator func(x int) (string, error)
}

func (ts *timeSpeller) Locale() string {
	return ts.locale
}

func (ts *timeSpeller) Translate(x int) (string, error) {
	return ts.translator(x)
}

func newTimeSpeller(locale string, translator func(x int) (string, error)) *timeSpeller {
	return &timeSpeller{
		locale:     locale,
		translator: translator,
	}
}
