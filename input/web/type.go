package web

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Web struct {
	Data   string
	KBData map[string]string
	client *http.Client
	logger *logrus.Logger
}

func NewWeb(logger *logrus.Logger) *Web {
	return &Web{

		logger: logger,
	}
}
