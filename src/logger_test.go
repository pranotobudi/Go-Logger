package main

import (
	"log"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()
	logger.Println("hello world")
	log.Println("bismillah, ")
}
