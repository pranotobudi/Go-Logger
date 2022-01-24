package main

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Trace("1")
	logger.Debug("2")
	logger.Info("3")
	logger.Warning("4")
	// logger.Panic("5")
	// logger.Fatal("6")
}
