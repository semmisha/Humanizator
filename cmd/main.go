package main

import (
	tableComarsion2 "MTtoHT/controllers/fs/requestFile"
	"MTtoHT/controllers/fs/tableComarsion"
	"MTtoHT/logging"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Read comparsion file
	// Read script file
	// Compare
	// Write to new file
	logger := logging.Logger()
	comm := exec.Command("clear")
	comm.Stdout = os.Stdout

	knowbase := tableComarsion.ReadTableComparsionFile("testData/sql_formated3.txt", logger)
	request := tableComarsion2.ReadSQLFile("testData/sql.txt", logger)
	requestString := strings.ToLower(request)

	for i, _ := range knowbase {

		requestString = strings.ReplaceAll(requestString, i, knowbase[i])

	}
	fmt.Println(requestString)

}

//inner join table alias on

//from table alias
