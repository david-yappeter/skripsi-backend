package validator

import "regexp"

var (
	alphaNumDashRegexString         = "^[a-zA-Z0-9\\-]+$"
	alphaNumDotRegexString          = "^[a-zA-Z0-9\\.]+$"
	alphaNumDashDotSlashRegexString = "^[a-zA-Z0-9\\-\\./]+$"
	splitParamsRegexString          = `'[^']*'|\S+`
)

var (
	alphaNumDashRegex         = regexp.MustCompile(alphaNumDashRegexString)
	alphaNumDotRegex          = regexp.MustCompile(alphaNumDotRegexString)
	alphaNumDashDotSlashRegex = regexp.MustCompile(alphaNumDashDotSlashRegexString)
	splitParamsRegex          = regexp.MustCompile(splitParamsRegexString)
)
