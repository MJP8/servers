package http
import (
	"fmt"
	"net/http"
)
const (
	InternalServerError = 500
	PathError = 402
	FileError = 336
	NotFoundError = 404
)
func GetString(code int) string {
	switch code {
	case InternalServerError:
		return "Internal server error"
	case PathError:
		return "Path to page not found"
	case FileError:
		return "File not found"
	case NotFoundError:
		return "Path or file not found"
	}
	return "Unrecognized status code"
}
type ServerError struct {
	Code int
	Msg string
}
func Error(errNum int, msg string, w ResponseWriter) *ServerError {
	writer := w
	statuscode := errNum
	message := msg
	if message == "" {
		message = GetString(statuscode)
	} else if writer != nil {
		return &ServerError{Code: statuscode, Msg: message}
	}
	serverError := &ServerError{Code: statuscode, Msg: message}
	switch statuscode {
	case InternalServerError:
		http.Error(w, serverError.ToErr().Error(), http.StatusInternalServerError)
	}
	return serverError
}
func (s *ServerError) ToErr() (err error) {
	err = fmt.Errorf("%s Error code %d", s.Msg, s.Code)
	return
}