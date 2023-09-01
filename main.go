package main

import (
	"dp_test/controllers"
	"dp_test/logger"
	"dp_test/shared"
	"fmt"
)

func main() {
	fmt.Println("Dep Injection Test")
	defer logger.CloseLogFile()

	logger1 := logger.WithWrappedModel(shared.Wrapped{
		Name:        "L1",
		Description: "this is logger 1",
	})

	logger2 := logger.WithWrappedModel(shared.Wrapped{
		Name:        "L2",
		Description: "this is logger 2",
	})

	c1 := controllers.NewController1(logger1)
	c1.ServeController1()

	c2 := controllers.NewController2(logger2)
	c2.ServeController2()

}
