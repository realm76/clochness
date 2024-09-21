package app

import (
	"database/sql"
	"embed"
	encoder "github.com/realm76/clochness/internal"
	"github.com/realm76/clochness/internal/components"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

//go:embed templates/*.gohtml
var templateFiles embed.FS

var templates = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))

func addRoutes(logger *zap.SugaredLogger, db *sql.DB, mux *http.ServeMux) {
	enc := encoder.NewEncoder(logger, templates)
	cmps := components.NewComponents(logger, db)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok 1223"))
		if err != nil {
			return
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		c, err := cmps.Index()
		if err != nil {
			enc.EncodeError(w, err)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		err = c.Render(w)
		if err != nil {
			enc.EncodeError(w, err)
			return
		}
	})
}
