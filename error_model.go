package customerror

import "log/slog"

var (
	logs                 *logging
	prefixLoggingMessage string
)

type CustomError struct {
	err              error
	message          string
	developerMessage string
	code             string
	op               string
}

type LevelError string

const (
	Error    LevelError = "Error"
	Debug    LevelError = "Debug"
	Info     LevelError = "Info"
	External LevelError = "External"
)

type logging struct {
	externalLog ExternalLog
	logger      slog.Logger
}

type CustomErrorInterface interface {
	Message() string
	DeveloperMessage() string
	Code() string
	Error() error
	Marshal() []byte
	As(err error) bool
	SupplementDevMessage(devMessage string) *CustomError
	SupplementDevMessageAndChangeCode(devMessage, code string) *CustomError
	SupplementDevMessageAndChangeCodeWithLogging(devMessage, code string, levelError LevelError) *CustomError
	SupplementDevMessageWithLogging(devMessage string, levelError LevelError) *CustomError
	ChangeDevMessage(devMessage string) *CustomError
	ChangeDevMessageWithLogging(devMessage string, levelError LevelError) *CustomError
	ChangeDevMessageAndCode(devMessage, code string) *CustomError
	ChangeDevMessageAndCodeWithLogging(devMessage, code string, levelError LevelError) *CustomError
}

type ExternalLog interface {
	Write(CustomError)
}
