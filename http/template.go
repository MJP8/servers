package http
import (
	"text/template"
)
type Template struct {
	Filename string
	Value interface{}
}
func RenderTemplate(w http.ResponseWriter, tmpl *Template) {
	t, err := template.ParseFiles(tmpl.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, tmpl.Value)
}
func NewTemplate(filename string, val interface{}) *Template {
	return &Template{Filename: filename, Value: val}
}