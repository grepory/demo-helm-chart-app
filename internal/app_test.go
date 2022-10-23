package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/message", nil)
	if err != nil {
		t.Fatal(err)
	}

	app := App{"testmessage"}

	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned bad status. want %v, got %v", http.StatusOK, status)
	}

	if rr.Body.String() != "testmessage" {
		t.Errorf("expected 'testmessage' got '%v'", rr.Body.String())
	}
}
