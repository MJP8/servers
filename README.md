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
	http.Init(":8000")
	http.HandleStaticFile("/", "index.html")
	http.Serve()
}
```
### Functions
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
	http.Init(8080)
	tmpl = NewTemplate("index.html", time.Now())
	http.HandleTemplate("/", tmpl)
	http.Serve()
}
```
