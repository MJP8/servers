package main
import (
	s "github.com/MJP8/servers/http"
	"time"
	"os"
	"fmt"
)
type Day struct {
	Hour int
	Minute int
}
func main() {
	day := &Day{} // set a Day struct
	day.Hour, day.Minute, _ = time.Now().Clock() // populate the struct with the current time
	template := s.NewTemplate("examples/templates/index.html", day)
	s.Init(4040) // initiate the server
	s.HandleCustom("/save/", func(w s.ResponseWriter, r *s.Request) {
		r.ParseForm()
		saveTime := r.Form.Get("time")
		err := os.WriteFile("examples/output/time.txt", []byte(saveTime), 0600)
		if err != nil {
			serr := s.Error(s.InternalServerError, err.Error(), w)
			err = serr.ToErr()
			fmt.Printf("%v", err)
		}
		fmt.Fprintf(w, "<script>location.assign(\"/\");</script>") // redirect to "/"
	})
	s.HandleTemplate("/", template) // add a template handler for the path "/"
	s.HandleStaticFile("/styles.css", "examples/static/styles.css", "text/css") // add a static handler for "/styles.css"
	s.Serve() // start the server
}