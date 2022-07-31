package usecases

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"runtime"
	"strings"
)

func NewUCProcessVar(env map[string]string, logger *logrus.Logger) *UCProcess {
	return &UCProcess{Logger: logger}
}
func (t *UCProcess) ProcessTableAlias() UC {

	var (
		tString = bytes.NewBufferString("")
		logger  = t.Logger
	)
	if len(t.Data) != 0 && len(t.KBData) != 0 {
		rows := strings.Split(t.Data, "\n")

		for i := range rows {
			words := strings.Fields(rows[i])
			for j := range words {
				if j < len(words)-3 && (reflect.DeepEqual(words[j], "from") || reflect.DeepEqual(words[j], "join")) {
					t.KBData[fmt.Sprint(words[j+2], ".")] = fmt.Sprint(words[j+1], ".")
					words[j+2] = ""

				}
				if j == len(words)-1 {
					tString.WriteString(fmt.Sprintf("%v\n", words[j]))

				} else {
					tString.WriteString(fmt.Sprintf("%v ", words[j]))
				}
			}
		}

		t.Data = tString.String()

	} else {
		logger.Fatalf("Unable to proceed %v, t.Data or/and T.KBData is empty", runtime.Frame{}.Func)

	}
	return t
}
func (t *UCProcess) Process() UC {
	var (
		logger = t.Logger
		data   = t.KBData
	)
	if len(data) == 0 {
		logger.Fatalf("\n  t.KBData is empty \n")

	}
	for key, value := range data {

		t.Data = strings.ReplaceAll(t.Data, key, value)
	}
	return t
}
func (t *UCProcess) ProcessVRD() UC {
	var (
		logger = t.Logger
		data   = t.KBVRDData
	)
	if len(data) == 0 {
		logger.Fatalf("\n  t.KBData is empty \n")

	}
	for key, value := range data {

		t.Data = strings.ReplaceAll(t.Data, key, value)
	}
	return t
}
func (t *UCProcess) Export() string {
	var (
		logger = t.Logger
	)

	if len(t.Data) == 0 {
		logger.Fatalf("\nt.Data or t.KBData is empty\n")
	}
	return t.Data
}
