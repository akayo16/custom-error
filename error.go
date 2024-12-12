package customerror

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

func (e *CustomError) Message() string { return e.message }

func (e *CustomError) DeveloperMessage() string { return e.developerMessage }

func (e *CustomError) Code() string { return e.code }

func (e *CustomError) Error() error { return e.err }

func (e *CustomError) Op() string { return e.op }

func (e *CustomError) Marshal() []byte {
	// TODO add error handling
	marshal, _ := json.Marshal(e)

	return marshal
}

func (e *CustomError) As(err error) bool {
	return errors.As(e.Error(), &err)
}

func (e *CustomError) SupplementDevMessage(devMessage string) *CustomError {
	e.developerMessage = fmt.Sprintf("%s %s", devMessage, e.developerMessage)

	return e
}

func (e *CustomError) SupplementDevMessageAndChangeCode(devMessage, code string) *CustomError {
	e.developerMessage = fmt.Sprintf("%s %s", devMessage, e.developerMessage)
	e.code = code

	return e
}

func (e *CustomError) SupplementDevMessageAndChangeCodeWithLogging(devMessage, code string, levelError LevelError) *CustomError {
	e.SupplementDevMessageAndChangeCode(devMessage, code)

	writeToLogs(e.developerMessage, levelError, e)

	return e
}

// Logging Occurs After Message Supplement
func (e *CustomError) SupplementDevMessageWithLogging(devMessage string, levelError LevelError) *CustomError {
	e.SupplementDevMessage(devMessage)

	writeToLogs(e.developerMessage, levelError, e)

	return e
}

func (e *CustomError) ChangeDevMessage(devMessage string) *CustomError {
	e.developerMessage = devMessage

	return e
}

// Logging Occurs After Change Dev Message
func (e *CustomError) ChangeDevMessageWithLogging(devMessage string, levelError LevelError) *CustomError {
	e.ChangeDevMessage(devMessage)

	writeToLogs(e.developerMessage, levelError, e)

	return e
}

func (e *CustomError) ChangeDevMessageAndCode(devMessage, code string) *CustomError {
	e.developerMessage = devMessage
	e.code = code

	return e
}

// Logging Occurs After Change Dev Message And Code
func (e *CustomError) ChangeDevMessageAndCodeWithLogging(devMessage, code string, levelError LevelError) *CustomError {
	e.ChangeDevMessageAndCode(devMessage, code)

	writeToLogs(e.developerMessage, levelError, e)

	return e
}

func As(err, target error) bool {
	return errors.As(err, &target)
}

// Sets http.StatusInternalServerError as default CustomError.Сode And LevelError as a LevelError.Error
func ShortCreateCustomError(err error, op, bodyErr string) *CustomError {
	return NewCustomErrorWithLevelLogging(
		err,
		fmt.Sprintf("%s", bodyErr),
		fmt.Sprintf("%s: %s", op, bodyErr),
		strconv.Itoa(http.StatusInternalServerError),
		op,
		Error,
	)
}

// Sets http.StatusInternalServerError as default CustomError.Сode
func ShortCreateCustomErrorWithLevelLogging(err error, op, bodyErr string, levelError LevelError) *CustomError {
	return NewCustomErrorWithLevelLogging(
		err,
		fmt.Sprintf("%s", bodyErr),
		fmt.Sprintf("%s: %s", op, bodyErr),
		strconv.Itoa(http.StatusInternalServerError),
		op,
		levelError,
	)
}

func NewCustomError(err error, message, developerMessage, code, op string) *CustomError {

	writeToLogs(developerMessage, Error, nil)

	return &CustomError{
		err:              err,
		message:          message,
		developerMessage: developerMessage,
		code:             code,
		op:               op,
	}
}

func NewCustomErrorWithLevelLogging(err error, message, developerMessage, code, op string, levelError LevelError) *CustomError {

	customErr := &CustomError{
		err:              err,
		message:          message,
		developerMessage: developerMessage,
		code:             code,
		op:               op,
	}

	writeToLogs(developerMessage, levelError, customErr)

	return customErr
}

func NewCustomErrorWithoutLogging(err error, message, developerMessage, code, op string) *CustomError {
	return &CustomError{
		err:              err,
		message:          message,
		developerMessage: developerMessage,
		code:             code,
		op:               op,
	}
}

func writeToLogs(message string, levelError LevelError, err *CustomError) {
	switch levelError {
	case Info:
		slog.Info(fmt.Sprintf("%s %s", prefixLoggingMessage, message))
	case Debug:
		slog.Debug(fmt.Sprintf("%s %s", prefixLoggingMessage, message))
	case Error:
		slog.Error(fmt.Sprintf("%s %s", prefixLoggingMessage, message))
	case External:
		logs.externalLog.Write(*err)
	}
}
