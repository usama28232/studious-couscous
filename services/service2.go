package services

import (
	"dp_test/models"

	"go.uber.org/zap"
)

type Service2 struct {
	logger   *zap.SugaredLogger
	service3 *Service3
	model    *models.Model2
}

func (s *Service2) ServeService2() {
	s.model.RepoFunc()
	s.service3.ServeService3()
}

func NewService2(logger *zap.SugaredLogger) *Service2 {
	return &Service2{
		logger:   logger,
		service3: NewService3(logger),
		model:    models.NewModel2(logger),
	}
}
