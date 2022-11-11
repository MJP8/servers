package http
import (
	"testing"
)
func TestGetString(t *testing.T) {
	wrongNum := 2
	wrong := GetString(wrongNum)
	right := GetString(InternalServerError)
	if wrong != "Unrecognized status code" {
		t.Fatalf("Unrecognized status code doesn't return error")
	} 
	if right == "Unrecognized status code" {
		t.Fatalf("GetString() doesn't recognize correct status code")
	}
}
func TestError(t *testing.T) {
	serverErr1 := Error(NotFoundError, "", nil)
	serverErr2 := Error(NotFoundError, "test error", nil)
	if serverErr1.Msg != "Path or file not found" || serverErr1.Code != NotFoundError {
		t.Fatalf("Error() returns incorrect data")
	}
	if serverErr2.Msg != "test error" {
		t.Fatalf("Error() does not return custom error message")
	}
}