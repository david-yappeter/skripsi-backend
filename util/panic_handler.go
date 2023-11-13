package util

import (
	"fmt"
	"log"
	"runtime/debug"
)

func FormatRecover(r interface{}) string {
	logMessage := ""
	switch v := r.(type) {
	case error:
		logMessage = fmt.Sprintf("Captured error: %s", v.Error())
	default:
		logMessage = fmt.Sprintf("Unhandled err type %T, Content: %+v", v, v)
	}

	if len(logMessage) != 0 {
		logMessage += "\n"
	}
	logMessage += string(debug.Stack())

	return logMessage
}

func PanicHandler() {
	if r := recover(); r != nil {
		log.Println(FormatRecover(r))
	}
}
