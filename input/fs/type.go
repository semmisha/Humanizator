package fs

import "github.com/sirupsen/logrus"

type FS struct {
	Env       map[string]string
	Data      string
	KBData    map[string]string
	KBVRDData map[string]string
	logger    *logrus.Logger
}
