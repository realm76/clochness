package nodeslist

import (
	"context"
	_ "embed"
	"github.com/realm76/ranger/ent"
	"github.com/realm76/ranger/ent/node"
	"go.uber.org/zap"
	"html/template"
	"io"
	"net/http"
	"strings"
)

//go:embed nodeslist.gohtml
var templateContent string

var componentTemplate = template.Must(template.New("component").Parse(templateContent))

type NodesListRequest struct {
	ParentNodeHandle string
}

type NodesListComponentFactory struct {
	logger *zap.SugaredLogger
	db     *ent.Client
}

type NodesListComponent struct {
	Nodes []ent.Node
}

func FromRequest(r *http.Request) NodesListRequest {
	if r == nil {
		return NodesListRequest{}
	}

	return NodesListRequest{
		ParentNodeHandle: strings.TrimSpace(r.URL.Query().Get("parent")),
	}
}

func NewNodesListFactory(logger *zap.SugaredLogger, db *ent.Client) *NodesListComponentFactory {
	return &NodesListComponentFactory{
		logger: logger,
		db:     db,
	}
}

func (f *NodesListComponentFactory) Make(ctx context.Context, req NodesListRequest) (*NodesListComponent, error) {
	query := f.db.Node.Query()

	if req.ParentNodeHandle != "" {
		query.Where(node.ParentHandle(req.ParentNodeHandle))
	}

	rawNodes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	nodes := make([]ent.Node, len(rawNodes))

	for i, n := range rawNodes {
		if n != nil {
			nodes[i] = *n
		}
	}

	return &NodesListComponent{
		Nodes: nodes,
	}, nil
}

func (n *NodesListComponent) Render(w io.Writer) error {
	if err := componentTemplate.Execute(w, n); err != nil {
		return err
	}

	return nil
}
