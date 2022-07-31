package fs

import (
	"Humanizator/input"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"os"
	"strings"
)

func NewFSVar(env map[string]string, logger *logrus.Logger) *FS {
	return &FS{
		Env:    env,
		logger: logger,
	}
}
func (t *FS) Write(s string) input.Input {
	return t
}
func (t *FS) ReadData(fspath string) input.Input {

	var (
		logger = t.logger
	)

	if valid := fs.ValidPath(fspath); !valid {
		logger.Fatal(fmt.Errorf("\nUnable to find file on provided path: %v\n", fspath))
	}

	file, err := os.Open(fspath)
	if err != nil {
		logger.Fatalf("\nUnable to open sql script file, error: %v\n", err)
	}
	defer file.Close()

	fileBody, err := io.ReadAll(file)
	if err != nil {
		logger.Fatalf("\nUnable to read body, error: %v\n", err)
	}
	fileBodyString := strings.ToLower(string(fileBody))

	t.Data = fileBodyString

	return t
}
func (t *FS) ReadKB(kbpath string) input.Input {
	var (
		logger = t.logger
	)

	if valid := fs.ValidPath(kbpath); !valid {
		logger.Fatal(fmt.Errorf("\nUnable to find file on provided path: %v\n", kbpath))
	}

	file, err := os.Open(kbpath)
	if err != nil {
		logger.Fatalf("\nUnable to open table comaprsion file, error: %v\n", err)
	}
	defer file.Close()

	t.KBData, err = godotenv.Parse(file)
	if err != nil {
		logger.Fatalf("\nUnable to read table comparsion file body, errof:%v\n", err)
	}
	if len(t.KBData) == 0 {
		logger.Fatal(errors.New("\nTable comparsion file is empty\n"))
	}
	//fmt.Println(compareMap)
	return t
}
func (t *FS) ReadVRDKB(kbpath string) input.Input {
	var (
		logger = t.logger
	)

	if valid := fs.ValidPath(kbpath); !valid {
		logger.Fatal(fmt.Errorf("\nUnable to find file on provided path: %v\n", kbpath))
	}

	file, err := os.Open(kbpath)
	if err != nil {
		logger.Fatalf("\nUnable to open table comaprsion file, error: %v\n", err)
	}
	defer file.Close()

	t.KBVRDData, err = godotenv.Parse(file)
	if err != nil {
		logger.Fatalf("\nUnable to read table comparsion file body, errof:%v\n", err)
	}
	if len(t.KBVRDData) == 0 {
		logger.Fatal(errors.New("\nTable comparsion file is empty\n"))
	}
	//fmt.Println(compareMap)
	return t
}
func (t *FS) Export() (string, map[string]string, map[string]string) {

	return t.Data, t.KBData, t.KBVRDData
}
