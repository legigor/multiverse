package main

import (
	"github.com/legigor/multiverse/controller"
	"github.com/sirupsen/logrus"
)

func main() {
	err := controller.Execute()
	if err != nil {
		logrus.Error(err)
	}
}
