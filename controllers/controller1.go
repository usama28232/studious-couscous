package controllers

import (
	"dp_test/services"

	"go.uber.org/zap"
)

type Controller1 struct {
	logger  *zap.SugaredLogger
	service *services.Service1
}

func (c *Controller1) ServeController1() {
	c.service.ServeService1()
}

func NewController1(logger *zap.SugaredLogger) *Controller1 {
	return &Controller1{
		logger:  logger,
		service: services.NewService1(logger),
	}
}
