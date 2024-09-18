package handlers

import (
	_ "embed"
	"encoding/json"
	"github.com/realm76/ranger/ent"
	encoder "github.com/realm76/ranger/internal"
	"github.com/realm76/ranger/internal/components/nodequickcreator"
	"github.com/realm76/ranger/internal/components/nodeslist"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

func HandleGetNodesList(logger *zap.SugaredLogger, db *ent.Client, encoder encoder.HtmlEncoder) http.HandlerFunc {
	if logger == nil {
		panic("nil logger")
	}

	if encoder == nil {
		panic("nil encoder")
	}

	componentFactory := nodeslist.NewNodesListFactory(logger, db)

	return func(w http.ResponseWriter, r *http.Request) {
		req := nodeslist.FromRequest(r)
		component, err := componentFactory.Make(r.Context(), req)
		if err != nil {
			encoder.EncodeError(w, err)
			return
		}

		encoder.EncodeComponent(w, r, http.StatusOK, component)
	}
}

func HandleNodeQuickCreator(logger *zap.SugaredLogger, db *ent.Client, encoder encoder.HtmlEncoder) http.HandlerFunc {
	if logger == nil {
		panic("nil logger")
	}

	if encoder == nil {
		panic("nil encoder")
	}

	componentFactory := nodequickcreator.NewNodeQuickCreatorFactory(logger, db)

	type NodeCreateRequest struct {
		Title string `json:"title"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK

		if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				encoder.EncodeError(w, err)
				return
			}

			var req NodeCreateRequest
			if err := json.Unmarshal(body, &req); err != nil {
				encoder.EncodeError(w, err)
				return
			}

			_, err = db.Node.Create().
				SetTitle(req.Title).
				SetHandle("/" + strings.ToLower(req.Title)).
				Save(r.Context())
			if err != nil {
				encoder.EncodeError(w, err)
				return
			}

			status = http.StatusCreated
		}

		component, err := componentFactory.Make(r.Context())
		if err != nil {
			encoder.EncodeError(w, err)
			return
		}

		encoder.EncodeComponent(w, r, status, component)
	}
}
