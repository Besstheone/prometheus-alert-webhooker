package telegram

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/krpn/prometheus-alert-webhooker/executor"
	"github.com/krpn/prometheus-alert-webhooker/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

const (
	paramToken   = "bot_token"
	paramChatID  = "chat_id"
	paramMessage = "message"

	defaultBuffer = 100
)

var requiredStringParameters = []string{
	paramToken,
	paramChatID,
	paramMessage,
}

//go:generate mockgen -source=telegram.go -destination=telegram_mocks.go -package=telegram doc github.com/golang/mock/gomock

// Telegram is the interface of Telegram client.
type Telegram interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

type task struct {
	executor.TaskBase
	chatID   int64
	message  string
	telegram Telegram
}

func (task *task) ExecutorName() string {
	return "telegram"
}

func (task *task) ExecutorDetails() interface{} {
	return map[string]interface{}{"chatID": task.chatID, "message": task.message}
}

func (task *task) Fingerprint() string {
	base := fmt.Sprintf("%v|%v", strconv.FormatInt(task.chatID, 10), task.message)
	return utils.MD5Hash(base)
}

func (task *task) Exec(logger *logrus.Logger) error {
	message := tgbotapi.NewMessage(task.chatID, task.message)
	_, err := task.telegram.Send(message)
	return err
}

type taskExecutor struct {
	client *http.Client
}

// NewExecutor creates TaskExecutor for Telegram tasks.
func NewExecutor(client *http.Client) executor.TaskExecutor {
	return taskExecutor{client: client}
}

func (executor taskExecutor) ValidateParameters(parameters map[string]interface{}) error {
	for _, reqParam := range requiredStringParameters {
		_, ok := parameters[reqParam]
		if !ok {
			return fmt.Errorf("required parameter %v is missing", reqParam)
		}
	}

	_, ok := parameters[paramChatID].(int)
	if !ok {
		return fmt.Errorf("%v parameter value is not int", paramChatID)
	}

	return nil
}

func (executor taskExecutor) NewTask(eventID, rule, alert string, blockTTL time.Duration, preparedParameters map[string]interface{}) executor.Task {
	task := &task{
		chatID:  int64(preparedParameters[paramChatID].(int)),
		message: preparedParameters[paramMessage].(string),
		telegram: &tgbotapi.BotAPI{
			Token:  preparedParameters[paramToken].(string),
			Buffer: defaultBuffer,
			Client: executor.client,
		},
	}

	task.SetBase(eventID, rule, alert, blockTTL)
	return task
}
