package web

import (
	"Humanizator/input"
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func (t *Web) ReadData(s string) input.Input {
	t.Data = s
	return t

}

func (t *Web) ReadKB(kbpath string) input.Input {
	var (
		logger = t.logger
	)

	//if valid := fs.ValidPath(kbpath); !valid {
	//	logger.Fatal(fmt.Errorf("\nUnable to find file on provided path: %v\n", kbpath))
	//}

	file, err := os.Open(kbpath)
	if err != nil {
		logger.Fatalf("\nUnable to open table comaprsion file, error: %v\n", err)
	}
	defer file.Close()

	KBData, err := godotenv.Parse(file)

	if err != nil {
		logger.Fatalf("\nUnable to read table comparsion file body, errof:%v\n", err)
	}
	if len(KBData) == 0 {
		logger.Fatal(errors.New("\nTable comparsion file is empty\n"))
	}
	t.KBData = KBData
	return t
}

func (t *Web) Write(s string) input.Input {
	//TODO implement me
	panic("implement me")
}

func (t *Web) Export() (string, map[string]string) {
	var (
		logger = t.logger
	)
	if len(t.KBData) == 0 || len(t.Data) == 0 {
		logger.Fatalf("\nt.Data len: %v  t.KBData len: %v\n", len(t.Data), len(t.KBData))
	}

	return t.Data, t.KBData

}
