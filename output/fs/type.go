package fs

import "github.com/sirupsen/logrus"

type Storage struct {
	Env    map[string]string
	Data   string
	logger *logrus.Logger
}
