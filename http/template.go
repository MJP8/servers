package http
import (
	"text/template"
	"net/http"
	"fmt"
)
type Template struct {
	Filename string
	Value interface{}
}
func RenderTemplate(w ResponseWriter, tmpl *Template) error {
	t, err := template.ParseFiles(tmpl.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	if w == nil || tmpl == nil {
		return fmt.Errorf("nil arguments")
	}
	t.Execute(w, tmpl.Value)
	return nil
}
func NewTemplate(filename string, val interface{}) *Template {
	return &Template{Filename: filename, Value: val}
}