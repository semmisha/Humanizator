package router

import "github.com/sirupsen/logrus"

type Router struct {
	Data   chan string
	logger *logrus.Logger
}
