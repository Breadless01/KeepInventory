package web

import (
	"html/template"
	"log"
	"net/http"

	"KeepInventory/internal/application"
)

// Server kapselt HTTP-Server und Templates.
type Server struct {
	bauteilService *application.BauteilService
	tmpl           *template.Template
}

func NewServer(bauteilService *application.BauteilService) *Server {
	tmpl, err := template.ParseFiles("web/templates/bauteile.html")
	if err != nil {
		log.Fatalf("konnte Templates nicht laden: %v", err)
	}

	return &Server{
		bauteilService: bauteilService,
		tmpl:           tmpl,
	}
}

func (s *Server) routes() {
	http.HandleFunc("/bauteile", s.handleBauteile)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/bauteile", http.StatusFound)
	})
}

// Start startet den HTTP-Server auf der angegebenen Adresse (z.B. ":8080").
func (s *Server) Start(addr string) error {
	s.routes()
	return http.ListenAndServe(addr, nil)
}
