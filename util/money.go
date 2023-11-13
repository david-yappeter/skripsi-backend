package util

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func CurrencyFormat(amount int, language language.Tag) string {
	p := message.NewPrinter(language)
	withCommaThousandSep := p.Sprintf("%d", amount)
	return withCommaThousandSep
}
