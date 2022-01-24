package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.WithField("username", "budi").WithField("name", "pranoto budi").Info("1")
}
