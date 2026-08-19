package httpapi

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTaskBehavior(t *testing.T) {
	rr := httptest.NewRecorder()
	TaskHTTPHandler(rr, httptest.NewRequest("POST", "/task", nil))
	body := rr.Body.String()
	if !strings.Contains(body, "closed") || strings.Contains(body, "\"Reminder\":true") {
		t.Fatalf("body=%s", body)
	}
}
