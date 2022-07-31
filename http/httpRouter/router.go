package httpRouter

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewRouterVar(data string, env map[string]string, logger *logrus.Logger) *Router {
	return &Router{
		Data:   data,
		Env:    env,
		logger: logger,
	}
}

func (t *Router) Setup() {
	var (
		logger = t.logger
	)
	http.HandleFunc("/", t.HandlerIndex)
	http.HandleFunc("/uploadFile", t.HandlerUploadFile)
	http.HandleFunc("/uploadText", t.HandlerUploadText)
	logger.Fatal(http.ListenAndServe(":8080", nil))

}
