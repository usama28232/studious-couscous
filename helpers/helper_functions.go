package helpers

import (
	"go.uber.org/zap"
)

type Helpers struct {
	logger zap.SugaredLogger
}

func (h *Helpers) HelpMe(s string) {
	h.logger.Infof("Helper is helping > %s", s)
}

func NewHelper(logger zap.SugaredLogger) *Helpers {
	return &Helpers{
		logger: logger,
	}
}
