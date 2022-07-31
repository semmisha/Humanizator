package httpRouter

import (
	"Humanizator/processor"
	"html/template"
	"io"
	"net/http"
	"strings"
)

func (t *Router) HandlerIndex(w http.ResponseWriter, r *http.Request) {

	type tmpltStruct struct {
		Title string
	}

	var (
		tmplData = tmpltStruct{Title: "hello bitches"}
		logger   = t.logger
	)

	tmplt := template.Must(template.ParseFiles("./templates/web/index.html"))

	err := tmplt.Execute(w, tmplData)
	if err != nil {
		logger.Fatalf("\nUnable to execute html template, error: %v\n")

	}

}

func (t *Router) HandlerUploadFile(w http.ResponseWriter, r *http.Request) {

	var (
		logger  = t.logger
		envData = t.Env
	)

	file, fileHeader, err := r.FormFile("myFile")

	if err != nil || fileHeader.Size == 0 {
		logger.Errorf("\nUnable to recieve file, error: %v\n", err)
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}

	fileBody, err := io.ReadAll(file)
	//fmt.Printf("\nfileBody size: %s\n", string(fileBody))
	if err != nil {

		logger.Fatalf("\nUnable to readAll file or file is empty, error: %v\n", err)

	}

	result := processor.ProcessWeb(string(fileBody), envData, logger)

	tmplt := template.Must(template.ParseFiles("./templates/web/postForm.html"))
	if err != nil {
		logger.Fatalf("\nUnable to parse templates, error:%v\n", err)
	}

	err = tmplt.Execute(w, string(result))
	if err != nil {

		logger.Fatalf("\nUnable to past data into template, error:%v\n", err)

	}

	defer file.Close()

}
func (t *Router) HandlerUploadText(w http.ResponseWriter, r *http.Request) {

	var (
		logger = t.logger
	)
	text := r.FormValue("myText")
	if len(strings.TrimSpace(text)) <= 0 {
		logger.Warnf("\nTextForm is empty\n")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	resp := processor.ProcessWeb(text, nil, logger)
	template.Must(template.ParseFiles("./templates/web/postForm.html")).Execute(w, string(resp))
	//w.Write(resp)
}
