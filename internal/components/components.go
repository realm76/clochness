package components

import (
	"database/sql"
	"github.com/go-jet/jet/v2/postgres"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/realm76/clochness/db/clochness/public/model"
	"github.com/realm76/clochness/db/clochness/public/table"
	"github.com/realm76/clochness/internal/entity"
	infrastructureA "github.com/realm76/clochness/internal/infrastructure"
	"go.uber.org/zap"
)

type Components struct {
	logger *zap.SugaredLogger
	db     *sql.DB
}

func NewComponents(logger *zap.SugaredLogger, db *sql.DB) *Components {
	return &Components{logger: logger, db: db}
}

func (c *Components) Index() (Node, error) {
	entriesStmt := postgres.SELECT(table.Entries.AllColumns).FROM(table.Entries.Table)

	var entries []entity.Entry
	var dbEntries []model.Entries

	err := entriesStmt.Query(c.db, &dbEntries)
	if err != nil {
		return nil, err
	}

	for _, entry := range dbEntries {
		entries = append(entries, infrastructureA.EntriesToEntry(entry))
	}

	return entryList(entries), nil
}

func entryList(entries []entity.Entry) Node {
	return Div(
		Ul(
			Group(Map(entries, func(entry entity.Entry) Node {
				return Li(Text(entry.Description))
			})),
		),
	)
}

func nodeQuickCreator() Node {
	return Div(
		Input(),
	)
}
