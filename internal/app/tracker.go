package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/realm76/clochness/internal/clochness"
	"net/http"
)

func (s *server) tracker(cr chi.Router) {
	cr.Get("/", s.getTracker)
	cr.Post("/", s.postTracker)
}

func (s *server) getTracker(w http.ResponseWriter, req *http.Request) {
	entries, err := s.queries.ListEntries(req.Context())
	if err != nil {
		s.renderError(w, err, 500)
		return
	}

	s.renderTemplate(w, http.StatusOK, s.templates.TrackerTmpl, map[string]interface{}{
		"Entries": entries,
	})
}

func (s *server) postTracker(w http.ResponseWriter, req *http.Request) {
	desc := req.FormValue("description")

	entry := clochness.CreateEntryParams{
		UserID:      1,
		Description: desc,
	}

	_, err := s.queries.CreateEntry(req.Context(), entry)
	if err != nil {
		s.renderError(w, err, 500)
		return
	}

	s.getTracker(w, req)
}
