package main

import (
	router2 "Humanizator/http/router"
	"Humanizator/logging"
)

func main() {
	var (
		logger = logging.Logger
		router = router2.NewRouter(logger())
	)
	router.Setup()

	// TODO ----- controller ----- //
	//controller.ReadData("testData/sql.txt")
	//controller.ReadKB("testData/sql_formated3.txt")
	//
	////TODO ----- UseCases -----//
	//
	//ucData.ProcessTableAlias()
	//ucData.Process()

	//logger.Printf("\n\n\n%v\n\n\n", ucData.Data)
}
