package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	input := "This is echo"
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(input))
	echo(rec, req)
	out := rec.Body.String()
	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK, but got '%s'", rec.Result().Status)
	}
	if input != out {
		t.Errorf("expected echo output '%s', but got '%s'", input, out)
	}
}

func TestEchoHelp(t *testing.T) {
	cases := []struct {
		name   string
		method string
		body   io.Reader
	}{
		{"Empty Post", http.MethodPost, strings.NewReader("")},
		{"Empty Post with nil body", http.MethodPost, nil},
		{"Empty Get", http.MethodGet, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, "/", tc.body)
			echo(rec, req)
			out := rec.Body.String()
			if !strings.HasPrefix(out, "This is echo service, try posting something") {
				t.Error("expected service to reply with help message on empty input")
			}
		})
	}
}
