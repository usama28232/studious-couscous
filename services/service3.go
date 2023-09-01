package services

import "go.uber.org/zap"

type Service3 struct {
	logger *zap.SugaredLogger
}

func (s *Service3) ServeService3() {
	s.logger.Info("Service3")
}

func NewService3(logger *zap.SugaredLogger) *Service3 {
	return &Service3{
		logger: logger,
	}
}
