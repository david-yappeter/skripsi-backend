package logger

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type telegramBotLogger struct {
	botApi *tgbotapi.BotAPI

	chatId int64
}

func (l *telegramBotLogger) Write(format string, v ...interface{}) {
	msg := tgbotapi.NewMessage(l.chatId, "<pre>"+fmt.Sprintf(format, v...)+"</pre>")
	msg.ParseMode = "HTML"

	go l.botApi.Send(msg)
}

func NewTelegramBotLogger(token string, chatId int64) (Logger, error) {
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	logger := &telegramBotLogger{
		botApi: botApi,

		chatId: chatId,
	}

	return logger, nil
}
