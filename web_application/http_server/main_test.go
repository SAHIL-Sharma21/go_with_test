package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Peeper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Peeper")
		respose := httptest.NewRecorder()

		server.ServeHTTP(respose, request)

		assertResponseBody(t, respose.Body.String(), "20")
	})

	t.Run("return the Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		respone := httptest.NewRecorder()

		server.ServeHTTP(respone, request)

		assertResponseBody(t, respone.Body.String(), "10")
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
