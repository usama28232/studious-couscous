package controllers

import (
	"dp_test/services"

	"go.uber.org/zap"
)

type Controller2 struct {
	logger  *zap.SugaredLogger
	service *services.Service2
}

func (c *Controller2) ServeController2() {
	c.service.ServeService2()
}

func NewController2(logger *zap.SugaredLogger) *Controller2 {
	return &Controller2{
		logger:  logger,
		service: services.NewService2(logger),
	}
}
