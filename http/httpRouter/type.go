package httpRouter

import "github.com/sirupsen/logrus"

type Router struct {
	Data   string
	Env    map[string]string
	logger *logrus.Logger
}
