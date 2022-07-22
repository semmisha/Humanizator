package router

import (
	"Humanizator/processor"
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func NewRouter(logger *logrus.Logger) *Router {
	return &Router{

		logger: logger,
	}
}

func (t *Router) Setup() {
	var (
		logger = t.logger
	)
	http.HandleFunc("/", t.HandlerIndex)
	http.HandleFunc("/upload", t.HandlerFormReader)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatalf("Unable to Listen and Serve, error: %v", err)

	}

}

func (t *Router) HandlerIndex(w http.ResponseWriter, r *http.Request) {
	var logger = t.logger
	tmplt, err := template.ParseFiles("./templates/http/index.html")
	fmt.Print(tmplt)
	if err != nil {
		logger.Fatalf("\nUnable to parse html template, error: %v\n", err)
	}
	err = tmplt.Execute(w, "")
	if err != nil {
		logger.Fatalf("\nUnable to execute html template, error: %v\n")

	}

}

////func (t *Router) HandlerUpload(w http.ResponseWriter, r *http.Request) {
//
//	file, _, err := r.FormFile("myFile")
//	if err != nil {
//		fmt.Printf("\nUnable to recieve file, error: %v\n", err)
//	}
//	fileByte, err := io.ReadAll(file)
//	if err != nil {
//
//		fmt.Printf("\nUnable to readAll file, error:%v\n")
//
//	}
//	t.Data = string(fileByte)
//	fmt.Println(t.Data)
//}

func (t *Router) HandlerFormReader(w http.ResponseWriter, r *http.Request) {

	var (
		logger = t.logger
	)

	request := r.FormValue("myFile")

	if len(request) == 0 {

		logger.Errorf("\nForm is empty\n")

	}
	result := processor.ProcessWeb(request, nil, logger)

	w.Write(result)
}

func (t *Router) Export(rChan chan string) {

	prock := <-t.Data
	fmt.Print(prock)
	rChan <- prock

}
