package processor

import (
	"Humanizator/input"
	"Humanizator/input/web"
	"Humanizator/output/fs"
	"Humanizator/usecases"
	"bytes"
	"github.com/sirupsen/logrus"
)

func ProcessWeb(str string, env map[string]string, logger *logrus.Logger) []byte {
	var (
		controller = web.NewWebVar(env, logger)
		ucData     = usecases.NewUCProcessVar(env, logger)
		stData     = fs.NewStorageVar(env, logger)
	)

	// TODO ----- Controller Logic -----//
	input.Input(controller).ReadKB(env["kbpath"]).ReadVRDKB(env["kbvrdpath"]).ReadData(str)

	ucData.Data, ucData.KBData, ucData.KBVRDData = input.Input(controller).Export()

	// TODO ----- UsesCase Logic -----//

	usecases.UC(ucData).ProcessTableAlias().ProcessVRD().Process()

	stData.Data = usecases.UC(ucData).Export()

	//TODO ----- Return -----//

	return bytes.NewBufferString(stData.Data).Bytes()
}
