package adapters

import (
	"go-bot/internal/usecase"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type TgAdapter struct {
	bot *tgbotapi.BotAPI
}

func InitTg(Token string) (*TgAdapter, error) {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		return nil, err
	}

	return &TgAdapter{bot: bot}, nil
}

func (tg *TgAdapter) SendMessage(text string, userID int64) error {
	msg := tgbotapi.NewMessage(userID, text)
	msg.ParseMode = "markdown"
	_, err := tg.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (tg *TgAdapter) TgUpdates(authId int64) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tg.bot.GetUpdatesChan(u)

	logrus.Info("Start listening for tg messages")

	for update := range updates {
		tg.ProcessingMessages(update, authId)
	}
}

func (t *TgAdapter) ProcessingMessages(update tgbotapi.Update, authId int64) {
	if update.Message != nil && update.Message.From.ID == authId {
		if update.Message.Document != nil && update.Message.Document.FileID != "" {
			logrus.Printf("[%s] %s", update.Message.From.UserName, update.Message.Document.FileID)
			href, err := t.bot.GetFileDirectURL(update.Message.Document.FileID)
			if err != nil {
				logrus.Panic(err)
			}

			errYandex := usecase.YandexDisk(href, strconv.Itoa(update.Message.MessageID))
			if errYandex != "ok" {
				t.SendMessage(errYandex, authId)
			}
		}
	}
}
