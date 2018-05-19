package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStore(t *testing.T) {
	app := App{}
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/api/v1/admins", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rec := httptest.NewRecorder()
	app.Store(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", res.Status)
	}
}
