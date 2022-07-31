package web

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Web struct {
	Env       map[string]string
	Data      string
	KBData    map[string]string
	KBVRDData map[string]string
	client    *http.Client
	logger    *logrus.Logger
}
