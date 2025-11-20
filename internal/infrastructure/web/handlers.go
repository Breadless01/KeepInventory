package web

import (
	"net/http"
	"strconv"

	"KeepInventory/internal/domain"
)

type bauteileViewModel struct {
	Bauteile []*domain.Bauteil
	Error    string
}

func (s *Server) handleBauteile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.listBauteile(w, r, "")
	case http.MethodPost:
		s.createBauteil(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) listBauteile(w http.ResponseWriter, r *http.Request, errMsg string) {
	bauteile, err := s.bauteilService.ListBauteile()
	if err != nil {
		http.Error(w, "Fehler beim Laden der Bauteile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	vm := bauteileViewModel{
		Bauteile: bauteile,
		Error:    errMsg,
	}

	if err := s.tmpl.Execute(w, vm); err != nil {
		http.Error(w, "Template-Fehler: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) createBauteil(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Ungültiges Formular", http.StatusBadRequest)
		return
	}

	name := r.Form.Get("name")
	beschreibung := r.Form.Get("beschreibung")
	lagerort := r.Form.Get("lagerort")
	lagerbestandStr := r.Form.Get("lagerbestand")

	lagerbestand := 0
	if lagerbestandStr != "" {
		if v, err := strconv.Atoi(lagerbestandStr); err == nil {
			lagerbestand = v
		} else {
			s.listBauteile(w, r, "Lagerbestand muss eine Zahl sein")
			return
		}
	}

	if _, err := s.bauteilService.CreateBauteil(name, beschreibung, lagerort, lagerbestand); err != nil {
		s.listBauteile(w, r, err.Error())
		return
	}

	http.Redirect(w, r, "/bauteile", http.StatusSeeOther)
}
