package customerror

import "log/slog"

type CustomError struct {
	err              error
	message          string
	developerMessage string
	code             string
}

type LevelError string

const (
	Error    LevelError = "Error"
	Debug    LevelError = "Debug"
	Info     LevelError = "Info"
	External LevelError = "External"
)

type ExternalLog interface {
	Write(CustomError)
}

type logging struct {
	externalLog ExternalLog
	logger      slog.Logger
}
