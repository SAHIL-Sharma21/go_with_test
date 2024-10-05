package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Peeper")
		respose := httptest.NewRecorder()

		PlayerServer(respose, request)

		got := respose.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})

	t.Run("return the Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		respone := httptest.NewRecorder()

		PlayerServer(respone, request)

		got := respone.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
	})
}

// code refactoring
func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
