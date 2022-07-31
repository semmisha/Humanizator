package fs

import "github.com/sirupsen/logrus"

func NewStorageVar(env map[string]string, logger *logrus.Logger) *Storage {
	return &Storage{

		Env:    env,
		logger: logger}
}

func (t *Storage) Write(url string) {
	panic("Impement me")
}
