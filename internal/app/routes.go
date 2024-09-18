package app

import (
	"bytes"
	"embed"
	"github.com/realm76/ranger/ent"
	encoder "github.com/realm76/ranger/internal"
	"github.com/realm76/ranger/internal/app/handlers"
	"github.com/realm76/ranger/internal/components/nodequickcreator"
	"github.com/realm76/ranger/internal/components/nodeslist"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

//go:embed templates/*.gohtml
var templateFiles embed.FS

var templates = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))

func addRoutes(logger *zap.SugaredLogger, db *ent.Client, mux *http.ServeMux) {
	enc := encoder.NewEncoder(logger, templates)
	nodeQuickCreatorFactory := nodequickcreator.NewNodeQuickCreatorFactory(logger, db)
	nodesListFactory := nodeslist.NewNodesListFactory(logger, db)

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

		nodeQuickCreatorComponent, err := nodeQuickCreatorFactory.Make(r.Context())
		if err != nil {
			enc.EncodeError(w, err)
			return
		}

		nodesListComponent, err := nodesListFactory.Make(r.Context(), nodeslist.FromRequest(r))
		if err != nil {
			enc.EncodeError(w, err)
			return
		}

		var buf bytes.Buffer
		var buf2 bytes.Buffer

		err = nodeQuickCreatorComponent.Render(&buf)
		if err != nil {
			enc.EncodeError(w, err)
			return
		}

		err = nodesListComponent.Render(&buf2)
		if err != nil {
			enc.EncodeError(w, err)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		err = templates.ExecuteTemplate(w, "index.gohtml", map[string]interface{}{
			"QuickCreator": template.HTML(buf.String()),
			"NodesList":    template.HTML(buf2.String()),
		})
		if err != nil {
			enc.EncodeError(w, err)
			return
		}
	})

	mux.HandleFunc("/components/nodeslist", handlers.HandleGetNodesList(logger, db, enc))
	mux.HandleFunc("/components/nodequickcreator", handlers.HandleNodeQuickCreator(logger, db, enc))
}
