package infrastructure

import (
	"myapp/global"
	internalLogger "myapp/internal/logger"
)

type LoggerStack interface {
	WriteAll(format string, v ...interface{})
}

type loggerStack []internalLogger.Logger

func (i loggerStack) WriteAll(format string, v ...interface{}) {
	for _, logger := range i {
		logger.Write(format, v...)
	}
}

func NewLoggerStack() LoggerStack {
	loggers := []internalLogger.Logger{}
	for _, logChannelName := range global.GetLogChannel() {
		switch logChannelName {
		case "stdout":
			loggers = append(loggers, internalLogger.NewStdoutLogger())

		case "telegram":
			telegramConfig := global.GetTelegramConfig()

			telegramBotLogger, err := internalLogger.NewTelegramBotLogger(telegramConfig.Token, telegramConfig.ChatId)
			if err != nil {
				panic(err)
			}

			loggers = append(loggers, telegramBotLogger)
		}
	}

	return loggerStack(loggers)
}
