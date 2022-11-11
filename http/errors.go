package http
import (
	"fmt"
)
type StatusCode int
const (
	InternalServerError StatusCode = 500
	PathError StatusCode = 402
	FileError StatusCode = 336
	NotFoundError StatusCode = 404
)
func (s StatusCode) GetString() string {
	switch s {
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
	Code StatusCode
	Msg string
}
func Error(errNum StatusCode, msg string, w ResponseWriter) *ServerError {
	writer := w
	statuscode := errNum
	message := msg
	if message == "" {
		message = statuscode.GetString()
	} else if writer != nil {
		return &ServerError{Code: statuscode, Msg: message}
	}
	return &ServerError{Code: statuscode, Msg: message}
}
func (s *ServerError) ToErr() (err error) {
	err = fmt.Errorf("%s Error code %d", s.Msg, s.Code)
	return
}