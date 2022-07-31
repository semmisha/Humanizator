package web

import (
	"Humanizator/input"
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func NewWebVar(env map[string]string, logger *logrus.Logger) *Web {
	return &Web{
		Env:    env,
		logger: logger,
	}
}

func (t *Web) ReadData(s string) input.Input {
	var (
		logger = t.logger
	)

	if len(s) == 0 {

		logger.Fatalf("\n input string is empty\n")

	}

	t.Data = strings.ToLower(s)

	return t

}

func (t *Web) ReadKB(kbpath string) input.Input {
	var (
		logger = t.logger
	)

	file, err := os.Open(kbpath)
	if err != nil {
		logger.Fatalf("\nUnable to open table comaprsion file, error: %v\n", err)
	}
	defer file.Close()

	var KBData = make(map[string]string)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		scnToLower := strings.ToLower(scan.Text())
		scnSlice := strings.Split(scnToLower, "=")
		if len(scnSlice) == 2 {
			KBData[scnSlice[0]] = scnSlice[1]

		} else {
			logger.Errorf("\nWrong format of KB line, line content:%v\n", scnSlice)
		}

	}
	//fmt.Println(KBData)
	t.KBData = KBData
	return t
}

func (t *Web) ReadVRDKB(kbpath string) input.Input {
	var (
		logger = t.logger
	)

	file, err := os.Open(kbpath)
	if err != nil {
		logger.Fatalf("\nUnable to open table comaprsion file, error: %v\n", err)
	}
	defer file.Close()

	var KBData = make(map[string]string)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		scnToLower := strings.ToLower(scan.Text())
		scnSlice := strings.Split(scnToLower, "=")
		if len(scnSlice) == 2 {
			KBData[scnSlice[0]] = scnSlice[1]

		} else {
			logger.Errorf("\nWrong format of KB line, line content:%v\n", scnSlice)
		}

	}
	//fmt.Println(KBData)
	t.KBVRDData = KBData
	return t
}

func (t *Web) Write(s string) input.Input {
	//TODO implement me
	panic("implement me")

}

func (t *Web) Export() (string, map[string]string, map[string]string) {
	var (
		logger = t.logger
	)
	if len(t.KBData) == 0 || len(t.Data) == 0 || len(t.KBVRDData) == 0 {
		logger.Fatalf("\nt.Data len: %v  t.KBData len: %v t.KBVRDData len:%v\n", len(t.Data), len(t.KBData), len(t.KBVRDData))
	}

	return t.Data, t.KBData, t.KBVRDData

}
