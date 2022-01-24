package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("1")
	logrus.Error("2")
	logrus.Warning("3")
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
