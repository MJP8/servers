# servers
**A Go module to help with creating and managing servers.**
## Documentation
### Get started
#### Install
To install, on Terminal or Command Prompt, run `go get github.com/MJP8/servers`. Include it in your program like this:
```go
package main
import (
	"github.com/MJP8/servers/http"
)
func main() {
	err := http.Init(8080)
	if err != nil {
		return
	}
	http.HandleStaticFile("/", "index.html", "text/html")
	http.Serve()
}
```
### Constants
```go
const ( // error codes
	InternalServerError = 500 // when you run the Error() function only InternalServerError will send the error to the browser
	NotFoundError = 404
	PathError = 402
	FileError = 336
)
```
### Functions
#### `func Error(errCode int, msg string, w ResponseWriter) *ServerError`
`Error()` return a [`*ServerError`](#type-servererror-struct).
#### `func HandleCustom(url string, cb func(ResponseWriter, *Request))`
`HandleCustom()` adds a handler for the URL specified. It runs the callback in the `cb` parameter.
#### `func HandleStaticFile(url string, filename string)`
`HandleStaticFile()` adds a handler for the URL, sending the data in the file.
#### `func HandleTemplate(url string, tmpl *Template)`
`HandleTemplate()` adds a handler for the URL specified, sending the template to the browser. See [`type Template`](#type-Template-struct) for more information
#### `func Init(port int)`
`Init()` initiates the server. The parameter `port` is the port for your server.
#### `func RenderTemplate(w ResponseWriter, tmpl *Template)`
`RenderTemplate()` renders the specified template. Used for inside the callback in the [`HandleCustom()`](#func-handlecustomurl-string-cb-funcresponsewriter-request) function.
#### `func Serve()`
`Serve()` starts the web server using the port set in the [`Init()`](#func-init-port-int) function.
### Types
#### `type Request`
```go
type Request = http.Request
```
`Request` defines a HTTP request from the `net/http` library.
#### `type ResponseWriter`
```go
type ResponseWriter = http.ResponseWriter
```
`ResponseWriter` creates a HTTP response from the `net/http` library.
#### `type ServerError struct`
```go
type ServerError struct {
	Code int
	Msg string
}
```
`ServerError` defines an error.
##### `func (s *ServerError) ToErr() error`
Converts a `ServerError` into an variable of type `error`.
#### `type Template struct`
```go
type Template struct {
	Filename string
	Value interface{}
}
```
`Template` defines a HTML template.
##### `func NewTemplate(filename string, val interface{}) *Template`
`NewTemplate()` creates a new `*Template`.<br/>
Example:
```go
package main
import (
	"time"
	"github.com/MJP8/servers/http"
)
func main() {
	err := http.Init(8080)
	if err != nil {
		return
	}
	tmpl := NewTemplate("index.html", time.Now().Format("15:04"))
	http.HandleTemplate("/", tmpl)
	http.Serve()
}
```
