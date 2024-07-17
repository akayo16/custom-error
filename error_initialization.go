package customerror

import "log/slog"

func InitializingLogging(
	external ExternalLog,
	logger slog.Logger,
	prefixForLoggingMessage *string,
) {
	logs = &logging{
		externalLog: external,
		logger:      logger,
	}

	prefixLoggingMessage = "Create New Custom Error With Dev Message:"

	if prefixForLoggingMessage != nil {
		prefixLoggingMessage = *prefixForLoggingMessage
	}
}
