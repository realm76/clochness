package nodequickcreator

import (
	"context"
	_ "embed"
	"github.com/realm76/ranger/ent"
	"go.uber.org/zap"
	"html/template"
	"io"
)

//go:embed nodequickcreator.gohtml
var templateContent string

var componentTemplate = template.Must(template.New("component").Parse(templateContent))

type NodeQuickCreatorFactory struct {
	logger *zap.SugaredLogger
	db     *ent.Client
}

type NodeQuickCreatorComponent struct {
}

func NewNodeQuickCreatorFactory(logger *zap.SugaredLogger, db *ent.Client) *NodeQuickCreatorFactory {
	return &NodeQuickCreatorFactory{
		logger: logger,
		db:     db,
	}
}

func (f *NodeQuickCreatorFactory) Make(ctx context.Context) (*NodeQuickCreatorComponent, error) {
	return &NodeQuickCreatorComponent{}, nil
}

func (c *NodeQuickCreatorComponent) Render(w io.Writer) error {
	if err := componentTemplate.Execute(w, c); err != nil {
		return err
	}

	return nil
}
