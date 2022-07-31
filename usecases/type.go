package usecases

import (
	"github.com/sirupsen/logrus"
)

type UCProcess struct {
	Envdata   map[string]string
	Data      string
	KBData    map[string]string
	KBVRDData map[string]string
	Logger    *logrus.Logger
}
