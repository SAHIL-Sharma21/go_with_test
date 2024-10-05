package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	server := &PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}

// type HandlerFunc func(http.ResponseWriter, *http.Request)

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "players/")

// 	if player == "Peeper" {
// 		fmt.Fprintf(w, "20")
// 		return
// 	}

// 	if player == "Floyd" {
// 		fmt.Fprintf(w, "10")
// 		return
// 	}
// }

// refactoring the code
type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "players")
	score := p.store.GetPlayerScore(player)

	fmt.Fprintf(w, "%d", score)
}

type StubPlayerStore struct {
	score map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.score[name]
	return score
}
