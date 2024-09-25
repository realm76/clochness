package app

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"html/template"
	"io"
	"net"
	"net/http"
)

type server struct {
	logger    *zap.SugaredLogger
	templates *templates
	db        *sql.DB
	router    *chi.Mux
}

func NewServer(logger *zap.SugaredLogger, db *sql.DB) *server {
	tmpls := &templates{}
	tmpls.parseTemplates()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s := &server{
		logger:    logger,
		db:        db,
		templates: tmpls,
		router:    r,
	}

	r.Route("/tracker", s.tracker)

	return s
}

func (s *server) ListenAndServe() {
	http.ListenAndServe(":3000", s.router)
}

func (s *server) HttpServer() *http.Server {
	return &http.Server{
		Addr:    net.JoinHostPort("", "3000"),
		Handler: s.router,
	}
}

func (s *server) renderTemplate(w http.ResponseWriter, status int, tmpl *template.Template, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	if err := tmpl.Execute(w, data); err != nil {
		s.renderError(w, err, 500)
		return
	}
}

func (s *server) renderError(w io.Writer, err error, status int) {
	if err := s.templates.ErrorTmpl.Execute(w, map[string]string{
		"Message": err.Error(),
	}); err != nil {
		s.logger.Errorln(err)
		_, _ = w.Write([]byte(err.Error()))
	}
}
