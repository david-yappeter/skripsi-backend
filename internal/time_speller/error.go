package time_speller

import "errors"

var ErrNegativeValue = errors.New("invalid value (negative)")
var ErrInvalidLocale = errors.New("invalid locale")
