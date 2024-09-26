package app

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/realm76/clochness/internal/clochness"
	"go.uber.org/zap"
	"html/template"
	"io"
	"net"
	"net/http"
	"strings"
)

//go:embed assets/*
var staticFiles embed.FS

type server struct {
	logger    *zap.SugaredLogger
	templates *templates
	db        *pgxpool.Pool
	router    *chi.Mux
	queries   *clochness.Queries
}

func NewServer(logger *zap.SugaredLogger, db *pgxpool.Pool) *server {
	tmpls := &templates{}
	tmpls.parseTemplates()

	queries := clochness.New(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.FS(staticFiles))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	s := &server{
		logger:    logger,
		db:        db,
		templates: tmpls,
		router:    r,
		queries:   queries,
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

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
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
