package fs

import "github.com/sirupsen/logrus"

func NewStorage(logger *logrus.Logger) *Storage {
	return &Storage{logger: logger}
}

func (t *Storage) Write(url string) {

}
