package app

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s *server) tracker(cr chi.Router) {
	cr.Get("/", s.getTracker)
	cr.Post("/", s.postTracker)
}

func (s *server) getTracker(w http.ResponseWriter, req *http.Request) {
	s.renderTemplate(w, http.StatusOK, s.templates.TrackerTmpl, nil)
}

func (s *server) postTracker(writer http.ResponseWriter, request *http.Request) {

}
