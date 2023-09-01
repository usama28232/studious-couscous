package models

import (
	"dp_test/helpers"

	"go.uber.org/zap"
)

type Model1 struct {
	logger *zap.SugaredLogger
}

func (m *Model1) RepoFunc() {
	m.logger.Info("Repo Request Model 1")
	h := helpers.NewHelper(*m.logger)
	h.HelpMe("req")
}

func NewModel1(logger *zap.SugaredLogger) *Model1 {
	return &Model1{
		logger: logger,
	}
}
