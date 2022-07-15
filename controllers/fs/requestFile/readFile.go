package requestFile

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
	"reflect"
	"strings"
)

func ReadSQLFile(fspath string, logger *logrus.Logger) string {

	if valid := fs.ValidPath(fspath); !valid {
		logger.Fatal(fmt.Errorf("\nUnable to find file on provided path: %v\n", fspath))
	}

	file, err := os.Open(fspath)
	if err != nil {
		logger.Fatalf("\nUnable to open sql script file, error: %v\n", err)
	}
	defer file.Close()
	var stringBytes = bytes.NewBufferString("")
	aliasMap := make(map[string]string)
	fileBody := bufio.NewScanner(file)
	for fileBody.Scan() {
		words := strings.Fields(fileBody.Text())
		for i := range words {

			if i < len(words)-3 && reflect.DeepEqual(strings.ToLower(words[i]), "join") {

				aliasMap[fmt.Sprint(words[i+2], ".")] = fmt.Sprint(words[i+1], ".")

				words[i+2] = ""
			} else if i < len(words)-3 && reflect.DeepEqual(strings.ToLower(words[i]), "from") {
				aliasMap[fmt.Sprint(words[i+2], ".")] = fmt.Sprint(words[i+1], ".")

				words[i+2] = ""

			}
			if i == len(words)-1 {
				stringBytes.WriteString(fmt.Sprintf("%v\n", words[i]))

			} else {
				stringBytes.WriteString(fmt.Sprintf("%v ", words[i]))
			}
		}

	}

	if err != nil {

		logger.Fatalf("\nUnable to read SQL file body, error: %v\n", err)
	}
	newString := strings.ToLower(stringBytes.String())
	for i, b := range aliasMap {
		newString = strings.ReplaceAll(newString, i, b)
	}
	fmt.Println(aliasMap)
	return newString
}
