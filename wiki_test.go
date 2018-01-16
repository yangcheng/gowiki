package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPageNameValidatd(t *testing.T) {

	names := []string{"/edit/normal123", "/edit/normal123&*"}
	statuses := []int{http.StatusOK, http.StatusNotFound}
	if len(names) != len(statuses) {
		t.Fatal("number of test names is different then name of status")
		return
	}

	for i, name := range names {
		req, err := http.NewRequest("GET", name, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(makeHandler(editHandler))

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != statuses[i] {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, statuses[i])
		}
	}

}
