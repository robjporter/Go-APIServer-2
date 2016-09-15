package render

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/unrolled/render"
)

var options = render.Options{IndentJSON: true, IndentXML: true, IsDevelopment: true}

func RenderTemplate(w http.ResponseWriter, templates []string, name string, data interface{}) error {
	t, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, name, data)
	if err != nil {
		return err
	}

	return nil
}

func RenderTemplate2(w http.ResponseWriter, templates []string, name string, data interface{}) error {
	t, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	var tmp io.Writer
	err = t.ExecuteTemplate(tmp, name, data)
	if err != nil {
		return err
	}

	fmt.Println(tmp)

	//r := render.New(options)
	//r.HTML(w, http.StatusOK, tmp, "")
	return nil
}

func RenderData(w http.ResponseWriter, status int, v interface{}) {
	r := render.New(options)
	r.Data(w, status, []byte(v.(string)))
}

func RenderText(w http.ResponseWriter, status int, v interface{}) {
	r := render.New(options)
	r.Text(w, status, v.(string))
}

func RenderJSONP(w http.ResponseWriter, status int, v interface{}) {
	r := render.New(options)
	r.JSONP(w, status, "callbackName", v)
}

func RenderJSON(w http.ResponseWriter, status int, v interface{}) {
	r := render.New(options)
	r.JSON(w, status, v)
}

func RenderXML(w http.ResponseWriter, status int, v interface{}) {
	r := render.New(options)
	r.XML(w, status, v)
}

func GetBaseTemplates() []string {
	return []string{}
}
