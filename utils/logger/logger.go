package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	baseLogger *zap.Logger
}

func (l *Logger) Info(messages ...any) {
	for _, msg := range messages {
		l.baseLogger.Info(fmt.Sprintf("%v", msg))
	}
}
