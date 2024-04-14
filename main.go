package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

type GameServer struct {
	t *template.Template
	State
}

func (s *GameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.t.Execute(w, s.State)
}

func NewGameServer(templateDir string) (s *GameServer, err error) {
	var t *template.Template

	t, err = template.ParseGlob(templateDir + "/**")
	if err != nil {
		return
	}

	s = &GameServer{t, State{}}
	s.State.Message = "Game Created"
	return
}

var (
	address     string = "127.0.0.1:8080"
	templateDir string = "templates"
)

func init() {
	flag.StringVar(&address, "address", address, "Server address")
	flag.StringVar(&templateDir, "board-template", templateDir, "Board template")
}

func main() {
	flag.Parse()

	s, err := NewGameServer(templateDir)
	if err != nil {
		log.Fatalf("error loading templates from '%s': %v", templateDir, err)
	}

	log.Printf("Starting server on %s", address)
	http.ListenAndServe(address, s)

	log.Printf("Server stopped")

}
