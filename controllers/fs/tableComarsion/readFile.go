package tableComarsion

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
)

func ReadTableComparsionFile(path string, logger *logrus.Logger) map[string]string {
	if valid := fs.ValidPath(path); !valid {
		logger.Fatal(fmt.Errorf("\nUnable to find file on provided path: %v\n", path))
	}

	file, err := os.Open(path)
	if err != nil {
		logger.Fatalf("\nUnable to open table comaprsion file, error: %v\n", err)
	}
	defer file.Close()

	compareMap, err := godotenv.Parse(file)
	if err != nil {
		logger.Fatalf("\nUnable to read table comparsion file body, errof:%v\n", err)
	}
	if len(compareMap) == 0 {
		logger.Fatal(errors.New("\nTable comparsion file is empty\n"))
	}
	//fmt.Println(compareMap)
	return compareMap
}
