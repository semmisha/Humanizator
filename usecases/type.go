package usecases

import (
	"github.com/sirupsen/logrus"
)

type UCProcess struct {
	Data   string
	KBData map[string]string
	Logger *logrus.Logger
}
