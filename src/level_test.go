package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.Trace("1")
	logger.Debug("2")
	logger.Info("3")
	logger.Warn("4")
	// logger.Error("5")
	// logger.Fatal("6")
}
