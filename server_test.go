package try

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJsonHandler(t *testing.T) {
	w := httptest.NewRecorder()
	err := writeJsonHeader(w, http.StatusOK, "")
	if err != nil {
		t.Fatal("This is a big error")
	}
}
