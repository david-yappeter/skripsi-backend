package util

import (
	"strings"
	"time"
)

func NewCode() string {
	return time.Now().Format("060102150405")
}

func StandarizeCode(code string) string {
	return strings.ToUpper(strings.TrimSpace(code))
}

func StandarizeNonInternalCode(code string) string {
	return strings.TrimSpace(code)
}
