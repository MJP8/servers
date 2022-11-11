package http
import (
	"text/template"
	"net/http"
)
type Template struct {
	Filename string
	Value interface{}
}
func RenderTemplate(w ResponseWriter, tmpl *Template) {
	t, err := template.ParseFiles(tmpl.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, tmpl.Value)
}
func NewTemplate(filename string, val interface{}) *Template {
	return &Template{Filename: filename, Value: val}
}