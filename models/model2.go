package models

import (
	"dp_test/helpers"

	"go.uber.org/zap"
)

type Model2 struct {
	logger *zap.SugaredLogger
}

func (m *Model2) RepoFunc() {
	m.logger.Info("Repo Request Model 2")
	h := helpers.NewHelper(*m.logger)
	h.HelpMe("req")
}

func NewModel2(logger *zap.SugaredLogger) *Model2 {
	return &Model2{
		logger: logger,
	}
}
