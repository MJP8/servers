package http
import (
	"net/http"
	"fmt"
	"os"
	"log"
)
type ResponseWriter = http.ResponseWriter
type Request = http.Request
var handleFuncs map[string]http.HandlerFunc
var inited bool
var port string
func Init(portint int) error {
	if len(fmt.Sprint(portint)) > 4 {
		return fmt.Errorf("%d: incorrect port value", portint)
	}
	handleFuncs = make(map[string]http.HandlerFunc)
	inited = true
	port = fmt.Sprintf(":%d", portint)
	return nil
}
func HandleStaticFile(url string, filename string, contentType string) {
	if !inited {
		return
	}
	handleFuncs[url] = func(w http.ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", contentType)
		data, err := os.ReadFile(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "%s", data)
	}
}
func HandleTemplate(url string, tmpl *Template) {
	if !inited {
		return
	}
	handleFuncs[url] = func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, tmpl)
	}
}
func HandleCustom(url string, cb func(w ResponseWriter, r *Request)) {
	if !inited {
		return
	}
	handleFuncs[url] = func(w http.ResponseWriter, r *http.Request) {
		cb(w, r)
	}
}
func Serve() {
	for url, handler := range handleFuncs {
		http.HandleFunc(url, handler)
	}
	log.Fatal(http.ListenAndServe(port, nil))
}