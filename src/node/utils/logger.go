package utils

import "go.uber.org/zap"

func InitLogger() (func()) {
	logger := zap.NewExample()

	undo := zap.RedirectStdLog(logger)

	return func () {
		undo()
	}
}