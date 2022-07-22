package fs

import "github.com/sirupsen/logrus"

type Storage struct {
	Data   string
	logger *logrus.Logger
}
