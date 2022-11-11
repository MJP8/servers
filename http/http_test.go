package http
import (
	"testing"
)
func TestInit(t *testing.T) {
	err := Init(8000)
	if err != nil {
		t.Fatalf("Error: Init() returns error for valid port value. %v", err)
	}
	err = Init(555552)
	if err == nil {
		t.Fatalf("Error: Init returns no error for invalid port value %d", 55552)
	}
}