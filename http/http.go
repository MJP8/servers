package http
import (
	"net/http"
	"fmt"
	"os"
	"log"
)
type ResponseWriter = http.ResponseWriter
type Request = http.Request
var handleFuncs = [string]func(http.ResponseWriter, *http.Request)
var init = false
var port string
func Init(portint int) {
	handleFuncs = make(map[string]func(ResponseWriter, *Request))
	init = true
	port = fmt.Sprintf(":%d", portint)
}
func HandleStaticFile(url string, filename string) {
	if !init {
		return
	}
	handleFuncs[url] = func(w http.ResponseWriter, r *Request) {
		data, err := os.ReadFile(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "%s", data)
	});
}
func HandleTemplate(url string, tmpl *Template) {
	if !init {
		return
	}
	handleFuncs[url] = func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, tmpl)
	}
}
func HandleCustom(url string, cb func(w ResponseWriter, r *Request)) {
	if !init {
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