package services

import (
	"dp_test/models"

	"go.uber.org/zap"
)

type Service1 struct {
	logger *zap.SugaredLogger
	model  *models.Model1
}

func (s *Service1) ServeService1() {
	s.model.RepoFunc()
}

func NewService1(logger *zap.SugaredLogger) *Service1 {
	return &Service1{
		logger: logger,
		model:  models.NewModel1(logger),
	}
}
