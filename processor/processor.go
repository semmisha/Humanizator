package processor

import (
	"Humanizator/input"
	"Humanizator/input/web"
	"Humanizator/output/fs"
	"Humanizator/usecases"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

func ProcessWeb(str string, env map[string]string, logger *logrus.Logger) []byte {
	var (
		controller = web.NewWeb(logger)
		ucData     = usecases.NewUCProcess(logger)
		stData     = fs.NewStorage(logger)
	)

	ucData.Data, ucData.KBData = input.Input(controller).ReadKB("./testData/sql.txt").ReadData(str).Export()
	stData.Data = usecases.UC(ucData).ProcessTableAlias().Process().Export()
	fmt.Println(stData.Data)
	return bytes.NewBufferString(stData.Data).Bytes()
}
