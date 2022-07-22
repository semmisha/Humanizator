package fs

import "github.com/sirupsen/logrus"

type FS struct {
	Data   string
	KBData map[string]string
	logger *logrus.Logger
}
